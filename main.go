package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type House struct {
	Value    float64 `json:"value"`
	Income   float64 `json:"income"`
	Age      int     `json:"age"`
	Rooms    int     `json:"rooms"`
	Bedrooms int     `json:"bedrooms"`
	Pop      int     `json:"pop"`
	HH       int     `json:"hh"`
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./main <input_csv_file> <output_json_file>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Convert CSV to JSON Lines
	err := convertCSVtoJSON(inputFile, outputFile)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Conversion completed successfully!")
}

func convertCSVtoJSON(inputPath, outputPath string) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open input file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV file: %v", err)
	}

	// Validate CSV structure
	if len(records) < 1 {
		return errors.New("CSV file is empty or improperly formatted")
	}
	headers := records[0]

	// Ensure headers match expected format
	expectedHeaders := []string{"value", "income", "age", "rooms", "bedrooms", "pop", "hh"}
	if len(headers) != len(expectedHeaders) {
		return errors.New("CSV headers do not match expected structure")
	}

	// Open output JSON file
	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer outFile.Close()

	// Parse CSV data and write to JSON Lines
	for _, record := range records[1:] {
		house, err := parseRecord(record)
		if err != nil {
			return fmt.Errorf("error parsing record: %v", err)
		}

		jsonData, err := json.Marshal(house)
		if err != nil {
			return fmt.Errorf("error encoding JSON: %v", err)
		}

		_, err = outFile.Write(append(jsonData, '\n'))
		if err != nil {
			return fmt.Errorf("error writing to output file: %v", err)
		}
	}

	return nil
}

func parseRecord(record []string) (House, error) {
	if len(record) != 7 {
		return House{}, errors.New("invalid record length")
	}

	value, err := strconv.ParseFloat(record[0], 64)
	if err != nil {
		return House{}, fmt.Errorf("invalid value: %v", err)
	}

	income, err := strconv.ParseFloat(record[1], 64)
	if err != nil {
		return House{}, fmt.Errorf("invalid income: %v", err)
	}

	age, err := strconv.Atoi(record[2])
	if err != nil {
		return House{}, fmt.Errorf("invalid age: %v", err)
	}

	rooms, err := strconv.Atoi(record[3])
	if err != nil {
		return House{}, fmt.Errorf("invalid rooms: %v", err)
	}

	bedrooms, err := strconv.Atoi(record[4])
	if err != nil {
		return House{}, fmt.Errorf("invalid bedrooms: %v", err)
	}

	pop, err := strconv.Atoi(record[5])
	if err != nil {
		return House{}, fmt.Errorf("invalid pop: %v", err)
	}

	hh, err := strconv.Atoi(record[6])
	if err != nil {
		return House{}, fmt.Errorf("invalid hh: %v", err)
	}

	return House{value, income, age, rooms, bedrooms, pop, hh}, nil
}
