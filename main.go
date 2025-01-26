package main

import (
	"errors"
	"strconv"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
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
	// Check if number of arguments provided is correct
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./main <input_csv_file> <output_json_file>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Try converting CSV to JSON Line, if there is something wrong show an error
	err := convertCSVtoJSON(inputFile, outputFile)
	if err != nil {
		fmt.Println("Error":, err)
		os.Exit(1)
	}

	fmt.Println("Conversion was successful!")
}

func convertCSVtoJSON(inputPath, outputPath string) error {
	//Open the input CSV file
	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("could not open the CSV file: %v", err)
	}
	defer file.Close()

	//Read all of the CSV file records
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("could not read the CSV file: %v", err)
	}

	// If the file is empty, return an error
	if len(records) < 1 {
		return errors.New("CSV file is empty or improperly formatted")
	}
	//Check the headers
	headers := records[0]

	if len(headers) != 7 {
		return errors.New("CSV headers do not match expected structure")
	}

	// Create ouput JSON file
	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer outFile.Close()

	// Loop each record and converT it to JSON
	for _, record := range records[1:] {
		house, err := parseRecord(record)
		if err != nil {
			return err
		}
// Convert house struct to JSON
		jsonData, err := json.Marshal(house)
		if err != nil {
			return err
//Write JSON to output file
		_, err = outFile.Write(append(jsonData, '\n'))
		if err != nil {
			return fmt.Errorf("error writing to output file: %v", err)
		}
	}

	return nil
}

func parseRecord(record []string) (House, error) {
	if len(record) != 7 {
		return House{}, errors.New("record has wrong number of fields")
	}
//Handle errors in each field
	value, err := strconv.ParseFloat(record[0], 64)
	if err != nil {
		return House{}, err
	}

	income, err := strconv.ParseFloat(record[1], 64)
	if err != nil {
		return House{}, err
	}

	age, err := strconv.Atoi(record[2])
	if err != nil {
		return House{}, err
	}

	rooms, err := strconv.Atoi(record[3])
	if err != nil {
		return House{}, err
	}

	bedrooms, err := strconv.Atoi(record[4])
	if err != nil {
		return House{}, err
	}

	pop, err := strconv.Atoi(record[5])
	if err != nil {
		return House{}, err
	}

	hh, err := strconv.Atoi(record[6])
	if err != nil {
		return House{}, err
	}
//Return populated house struct
	return House{value, income, age, rooms, bedrooms, pop, hh}, nil
}
