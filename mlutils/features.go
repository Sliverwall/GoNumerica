package mlutils

import (
	"GoNumerica/arei"
	"errors"
	"log"
)

// Package to handle feature transforms, primarlly using areis

// LabelEncode takes mx1 Arei then returns an mxn Arei with each entry converted into a one-hot label
func LabelEncode(y *arei.Arei) (*arei.Arei, error) {
	// Set up to only process one feature column
	if y.Shape[1] != 1 {
		return nil, errors.New("label encoder only proccess 1d feature columns")
	}

	// Part 1: Initialize needed one_hot encoded labels
	// mx1 arei with unique labels
	uniqueLabels := arei.Unique(y)

	// Init slice to hold all needed one_hot encoded labels
	var encodedLabels [][]float64
	// Loop through each element in unique labels slice
	for i := range uniqueLabels.Shape[0] {
		// Init slice to hold encoded label
		encodedLabel := make([]float64, uniqueLabels.Shape[0])
		// Loop through each element in unique label slice again
		for j := range uniqueLabels.Shape[0] {
			// Init a 1 at the jth element to identify label. i.e. the 3rd unique label will []float64{0,0,1}
			if i == j {
				encodedLabel[j] = 1
			} else {
				encodedLabel[j] = 0
			}
		}
		// After doing 2nd loop, added encoded label to encoded labels slice
		// Can match position in slice to position in unique labels to translate back
		log.Println(encodedLabel)
		encodedLabels = append(encodedLabels, encodedLabel)

	}

	// Part 2: Translate inputted arei into one_hot encoded
	rows, cols := y.Shape[0], uniqueLabels.Shape[0] // Number of rows remain same as input, but columns expand for each new unique label
	// Init proper shape slice for oneHotLabel
	oneHotLabel := make([][]float64, rows)
	for i := range oneHotLabel {
		oneHotLabel[i] = make([]float64, cols)
	}

	// Loop through each row
	for i := range rows {
		// Check the element in the original arei
		element, _ := y.Index(i, 0)
		var elementIndex int
		// Check where it appears in unique label list
		for index, value := range uniqueLabels.Data {
			if value == element {
				// value found set index
				elementIndex = index
			}
		}
		// Set oneHotLabel to encodedLabel found for the particular label
		oneHotLabel[i] = encodedLabels[elementIndex]
	}

	// Create and return an arrei using oneHotLabel data
	return arei.NewArei(oneHotLabel)

}
