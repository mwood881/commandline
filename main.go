package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"

	// import the newest version of cobra after installing the latest version
	"github.com/spf13/cobra"
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

// cobra uses var for input and output files

var inputFile string
var outputFile string

//Include root command with cmd from cobra

var rootCmd = &cobra.Command{
	Use:   "csv2json",
	Short: "Convert CSV to JSON",
	Long:  "A program that converts CSV files to JSON lines partficularly for the example house data",
	Run: func(cmd *cobra.Command, args []string) {
		// CSV to JSON
		err := convertCSVtoJSON(inputFile, outputFile)
		if err != nil {
			fmt.Printf("Error: reading file failed", err)
		}
		fmt.Println("Converted successfully")
	},
}

// Conversion function
func convertCSVtoJSON(inputPath, outputPath string) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("input file failed", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("reading file failed", err)
	}

	// CSV structure correct
	if len(records) < 1 {
		return errors.New("CSV file not formatted right")
	}
	headers := records[0]

	if len(headers) != 7 {
		return errors.New("CSV headers do not match structure")
	}

	//Open output JSON file
	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("creating output file failure", err)
	}
	defer outFile.Close()

	// Write opening bracket for the JSON array
	outFile.Write([]byte("[\n"))

	//  write to JSON Lines with CSV data
	for i, record := range records[1:] {
		house, err := parseRecord(record)
		if err != nil {
			return fmt.Errorf("error", err)
		}

		jsonData, err := json.Marshal(house)
		if err != nil {
			return fmt.Errorf("error with JSON", err)
		}

		// Write JSON data to file
		_, err = outFile.Write(jsonData)
		if err != nil {
			return fmt.Errorf("error writing to output file", err)
		}

		// Write a newline after each JSON object to separate entries
		_, err = outFile.Write([]byte("\n"))
		if err != nil {
			return fmt.Errorf("error writing newline to output file", err)
		}
		// Check if this is the last record
		if i < len(records)-2 {
			// Not the last record, write a comma after it
			outFile.Write([]byte(",\n"))
		} else {
			// For the last record, write only a newline (no comma)
			outFile.Write([]byte("\n"))
		}
	}
	// Write closing bracket for the JSON array
	outFile.Write([]byte("]\n"))
	// Remember the JSONlist needs brackets and then commas after each object in the file
	//Therefore, it looks like this [{object1}, {object2}]

	return nil

}

// Function for CSV reading in file
func parseRecord(record []string) (House, error) {
	if len(record) != 7 {
		return House{}, errors.New("invalid length")
	}

	value, err := strconv.ParseFloat(record[0], 64)
	if err != nil {
		return House{}, fmt.Errorf("invalid value", err)
	}

	income, err := strconv.ParseFloat(record[1], 64)
	if err != nil {
		return House{}, fmt.Errorf("invalid income", err)
	}

	age, err := strconv.Atoi(record[2])
	if err != nil {
		return House{}, fmt.Errorf("invalid age", err)
	}

	rooms, err := strconv.Atoi(record[3])
	if err != nil {
		return House{}, fmt.Errorf("invalid rooms", err)
	}

	bedrooms, err := strconv.Atoi(record[4])
	if err != nil {
		return House{}, fmt.Errorf("invalid bedrooms", err)
	}

	pop, err := strconv.Atoi(record[5])
	if err != nil {
		return House{}, fmt.Errorf("invalid pop", err)
	}

	hh, err := strconv.Atoi(record[6])
	if err != nil {
		return House{}, fmt.Errorf("invalid hh", err)
	}

	return House{value, income, age, rooms, bedrooms, pop, hh}, nil
}

func init() {
	// Define flags
	rootCmd.PersistentFlags().StringVarP(&inputFile, "input", "i", "", "Path to the input CSV file")
	rootCmd.PersistentFlags().StringVarP(&outputFile, "output", "o", "", "Path to the output JSON file")

	// Mark flags as required
	rootCmd.MarkPersistentFlagRequired("input")
	rootCmd.MarkPersistentFlagRequired("output")
}

func main() {
	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
