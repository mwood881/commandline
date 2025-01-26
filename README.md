**CSV to JSON Converter using Go**

This program takes data from a CSV file and converts it into a JSON file. Specifically for this program, we used data from a 
house related dataset. To perform this program, Go language was used along with the cobra package to create a commandline application.

***Project Description***

We are able to convert the CSV data like value, income, and rooms of house data into a JSON file. 

***Features***

- converts CSV to JSON file.
- Each house data row in the CSV file will be one line in the JSON file.

***PrePrograming***

Be sure to have Go installed on your computer using the instructions here: https://go.dev/doc/install. 

Next, install the cobra package: 
```
go get github.com/spf13/cobra 
```


***The Program***
1) Clone the repository:
```
git clone https://github.com/mwood8881/commandline.git
cd commandline

```
2) Build the program:

3) Run the program:
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
