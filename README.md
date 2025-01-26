# CSV to JSON Converter using Go

Github Repository: https://github.com/mwood8881/commandline.git

This program takes data from a CSV file and converts it into a JSON file. Specifically for this program, we used data from a 
house related dataset. To perform this program, Go language was used along with the cobra package to create a commandline application.

## Project Description

We are able to convert the CSV data like value, income, and rooms of house data into a JSON file. 

### Features

- converts CSV to JSON file.
- Each house data row in the CSV file will be one line in the JSON file.

## Prerequisites
Install the following: 
1. [Go](https://go.dev/doc/install) 
2. `cobra` package (install it via `go get github.com/spf13/cobra`).


## Get Started

1) Clone the repository:
```bash
git clone https://github.com/mwood8881/commandline.git
cd commandline

```
2) Install Cobra:
```bash
go get github.com/spf13/cobra
```
3) build the program which is in main.go

4) Run the program:
```
./csv2json --input housesInput.csv --output housesOutput.json

```
4) Test the output file:
```
./csv2json --input housesInput.csv --output housesOutput.json

```


***Code Example***
```
100000,50000,10,4,3,2000,1200
150000,60000,8,5,4,2500,1500
```

```
[
  {"value":100000,"income":50000,"age":10,"rooms":4,"bedrooms":3,"pop":2000,"hh":1200},
  {"value":150000,"income":60000,"age":8,"rooms":5,"bedrooms":4,"pop":2500,"hh":1500}
]

```
It is important to note that the JSON file must have brackets at the start and end of the file and commas separating each object or line in the JSON file. 

You can test if the JSON file is formatted correctly with: https://jsonlint.com/ 


AI application was used to understand how to include the brackets at the beginning and end of the JSON file and include commas after each object or line. This continued to produce errors until I used AI to help understand this was the issue. 
```bash
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
```
