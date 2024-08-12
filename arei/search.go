package arei

import "math"

// Search for collection of functions that search Areis

// Where returns the indices of elements that satisfy the given condition.
func Where(a *Arei, condition func(float64) bool) (*Arei, error) {
	var numRows, numCols int

	if len(a.shape) == 2 {
		numRows = a.shape[0]
		numCols = a.shape[1]
	} else if len(a.shape) == 1 {
		numRows = 1
		numCols = a.shape[0]
	}

	// Initialize an empty slice to store the indices of max elements for each row
	indices := make([][]float64, 0)
	// Iterate over each row
	for i := 0; i < numRows; i++ {
		// Iterate over each column
		for j := 0; j < numCols; j++ {
			element := a.data[i*numCols+j]
			if condition(element) {
				hold := make([]float64, 2)
				hold[0] = float64(i)
				hold[1] = float64(j)
				indices = append(indices, hold)
			}
		}
	}

	// Return the indices as a new Arei
	resultArei, err := NewArei(indices)
	return resultArei, err
}

// WhereMax finds the indices of the maximum value along each row of an arei.
func WhereMax(a *Arei) (*Arei, error) {

	var numRows, numCols int

	if len(a.shape) == 2 {
		numRows = a.shape[0]
		numCols = a.shape[1]
	} else if len(a.shape) == 1 {
		numRows = 1
		numCols = a.shape[0]
	}

	// Initialize an empty slice to store the indices of max elements for each row
	indices := make([]float64, numRows)

	// Iterate over each row
	for i := 0; i < numRows; i++ {
		maxValue := math.Inf(-1)
		maxIndex := 0
		// Iterate over each column in the current row to find the max value
		for j := 0; j < numCols; j++ {
			element := a.data[i*numCols+j]
			if element > maxValue {
				maxValue = element
				maxIndex = j
			}
		}
		// Store the index of the max element for the current row
		indices[i] = float64(maxIndex)
	}

	// Return the indices as a new Arei
	return NewArei(indices)
}

// WhereMax finds the indices of the maximum value along each row of an arei.
func WhereMin(a *Arei) (*Arei, error) {

	var numRows, numCols int

	if len(a.shape) == 2 {
		numRows = a.shape[0]
		numCols = a.shape[1]
	} else if len(a.shape) == 1 {
		numRows = 1
		numCols = a.shape[0]
	}

	// Initialize an empty slice to store the indices of max elements for each row
	indices := make([]float64, numRows)

	// Iterate over each row
	for i := 0; i < numRows; i++ {
		minValue := math.Inf(1)
		minIndex := 0
		// Iterate over each column in the current row to find the min value
		for j := 0; j < numCols; j++ {
			element := a.data[i*numCols+j]
			if element > minValue {
				minValue = element
				minIndex = j
			}
		}
		// Store the index of the min element for the current row
		indices[i] = float64(minIndex)
	}

	// Return the indices as a new Arei
	return NewArei(indices)
}
