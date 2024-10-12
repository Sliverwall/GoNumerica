package arei

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// Functions that read in data then convert to Areis

// Function to read a .data file and return its content as an arei.
func ReadDataFile(filePath string) (*Arei, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)
	reader.Comma = ',' // Set delimiter (comma)

	// Read all records from the file
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Prepare a 2D slice to hold float64 values
	floatRecords := make([][]float64, len(records))

	// Convert each record to float64
	for i, record := range records {
		floatRow := make([]float64, len(record))
		for j, value := range record {
			// Convert each string value to float64
			floatValue, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return nil, fmt.Errorf("error converting value '%s' in row %d, column %d: %v", value, i, j, err)
			}
			floatRow[j] = floatValue
		}
		floatRecords[i] = floatRow
	}

	// return floatRecords as an Arei
	a, err := NewArei(floatRecords)
	if err != nil {
		return nil, err
	}
	return a, nil
}
