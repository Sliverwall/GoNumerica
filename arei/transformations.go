package arei

import (
	"errors"
	"math"
)

// Transformations take an existing Arei and apply a transformation to create a new Arei

// Transform takes an Arei and a given transformation function and applies to each element
func Transform(a *Arei, transformation func(float64) float64) *Arei {
	resultData := make([]float64, len(a.Data))

	for i := range a.Data {
		resultData[i] = transformation(a.Data[i])
	}

	return &Arei{
		Shape: a.Shape,
		Data:  resultData,
	}
}

// Exp takes each element in an Arei with base euler's number
func Exp(a *Arei) *Arei {

	resultData := make([]float64, len(a.Data))

	for i := range a.Data {
		resultData[i] = math.Exp(a.Data[i])
	}

	return &Arei{
		Shape: a.Shape,
		Data:  resultData,
	}
}

// Sign creates a new arei with flipped signs on each element
func Sign(a *Arei) *Arei {

	resultData := make([]float64, len(a.Data))

	for i := range a.Data {
		resultData[i] = a.Data[i] * -1
	}

	return &Arei{
		Shape: a.Shape,
		Data:  resultData,
	}
}

// MultiT takes a given x then multiplies each element of the matrix by x.
func MultiT(a *Arei, x float64) *Arei {

	resultData := make([]float64, len(a.Data))

	for i := range a.Data {
		resultData[i] = a.Data[i] * x
	}

	return &Arei{
		Shape: a.Shape,
		Data:  resultData,
	}
}

// DivT takes a given x then multiplies each element of the matrix by x.
func DivT(a *Arei, x float64) (*Arei, error) {

	// Handle x = 0
	if x == 0 {
		return nil, errors.New("x cannot be 0")
	}
	resultData := make([]float64, len(a.Data))

	for i := range a.Data {
		resultData[i] = a.Data[i] / x
	}

	return &Arei{
		Shape: a.Shape,
		Data:  resultData,
	}, nil
}

// Maximum takes a given x then compares it to each element in matrix arei a. If x is > ai, x, otherwise ai.
func Maximum(a *Arei, x float64) *Arei {

	resultData := make([]float64, len(a.Data))

	for i := range a.Data {
		if x > a.Data[i] {
			resultData[i] = x
		} else {
			resultData[i] = a.Data[i]
		}
	}

	return &Arei{
		Shape: a.Shape,
		Data:  resultData,
	}
}

// Compare takes a given x then compares it to each element in arei a. if x > ai, x, otherwise 1
func Compare(a *Arei, x float64) *Arei {

	resultData := make([]float64, len(a.Data))

	for i := range a.Data {
		if x > a.Data[i] {
			resultData[i] = x
		} else {
			resultData[i] = 1
		}
	}

	return &Arei{
		Shape: a.Shape,
		Data:  resultData,
	}
}

// ArgMax takes an arei, then returns an arei with 1 in the row-column position with the highest value for that column.
func ArgMax(a *Arei) *Arei {
	resultData := make([]float64, len(a.Data))

	// Check if vector or not
	cols := 1
	if len(a.Shape) != 1 {
		cols = a.Shape[1]
	}
	// Init slice to store data on which row-column should be 1
	argmaxMap := make([]int, cols)
	// Loop through each column
	for j := range cols {
		// Init values to keep track of max value and the row
		maxI := 0
		maxValue := math.Inf(-1)
		// Loop through each row
		for i := range a.Data {

			// Check is record is greater than current max value
			element, _ := a.Index(i, j)
			if element > maxValue {
				// Set new maxI and new maxValue if so
				maxValue = element
				maxI = i
			}
		}
		// After looping thorugh all rows in column, store row in map. Index of map denotes j
		argmaxMap[j] = maxI
	}

	// Loop through each row
	for i := range a.Data {
		// Loop through each column
		for j := range cols {
			// Set result data using softmax
			iValue := argmaxMap[j]

			// If current i matches value stored in argmaxMap, then current element is max value for the column, so set value to 1, otherwise 0.
			if iValue == i {
				resultData[i*cols+j] = 1
			} else {
				resultData[i*cols+j] = 0
			}

		}
	}

	return &Arei{
		Shape: a.Shape,
		Data:  resultData,
	}
}

// SoftMax takes an arei, then returns an arei with with the softmax for each element in comparsion to the sum of the elemnet's column softmax
func SoftMax(a *Arei) *Arei {
	resultData := make([]float64, len(a.Data))

	// Check if vector or not
	cols := 1
	if len(a.Shape) != 1 {
		cols = a.Shape[1]
	}
	// Init slice to store data on which row-column should be 1
	softmaxMap := make([]float64, cols)
	// Loop through each column
	for j := range cols {
		// Init values to keep the sum of the softmax for the column
		sum := 0.0

		// Loop through each row
		for i := range a.Data {

			// Get the element
			element, _ := a.Index(i, j)
			// Add e^element to sum variable
			sum += math.Exp(element)
		}
		// After looping thorugh all rows in column, store row in map. Index of map denotes j
		softmaxMap[j] = sum
	}

	// Loop through each row
	for i := range a.Data {
		// Loop through each column
		for j := range cols {
			// Set result data using softmax
			element, _ := a.Index(i, j)
			softMaxSum := softmaxMap[j]
			eBaseElement := math.Exp(element)
			softMaxElement := eBaseElement / softMaxSum

			// Set softMaxElement to resultData
			resultData[i*cols+j] = softMaxElement

		}
	}

	return &Arei{
		Shape: a.Shape,
		Data:  resultData,
	}
}
