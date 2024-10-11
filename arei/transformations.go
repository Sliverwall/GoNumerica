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

// DivT takes a given x then divides each element of the matrix by x.
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

// AddT takes a given x then adds each element of the matrix by x.
func AddT(a *Arei, x float64) *Arei {

	resultData := make([]float64, len(a.Data))

	for i := range a.Data {
		resultData[i] = a.Data[i] + x
	}

	return &Arei{
		Shape: a.Shape,
		Data:  resultData,
	}
}

// SubT takes a given x then subtracts each element of the matrix by x.
func SubT(a *Arei, x float64) *Arei {

	resultData := make([]float64, len(a.Data))

	for i := range a.Data {
		resultData[i] = a.Data[i] - x
	}

	return &Arei{
		Shape: a.Shape,
		Data:  resultData,
	}
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

// ArgMax takes an arei, then returns an arei with 1 in the row-column position with the highest value for that column. direction > 0 to set along row, and <= 0 to set along column
func ArgMax(a *Arei, direction int) *Arei {
	// Init A with same shape as  a
	argMaxArei, _ := Zeros(a.Shape)

	// Init slice to store data on which row-column should be 1
	argmaxMap := make([]int, a.Shape[1])

	// If direction > 0,  get argmax along row, otherwise along column

	if direction > 0 {
		// Loop through each column
		for i := range a.Shape[0] {
			// Init values to keep track of max value and the row
			maxJ := 0
			maxValue := math.Inf(-1)
			// Loop through each row
			for j := range a.Shape[1] {

				// Check is record is greater than current max value
				element, _ := a.Index(i, j)
				if element > maxValue {
					// Set new maxI and new maxValue if so
					maxValue = element
					maxJ = j
				}
			}
			// After looping thorugh all rows in column, store row in map. Index of map denotes j
			argmaxMap[i] = maxJ
		}
	} else {
		// Loop through each column
		for j := range a.Shape[1] {
			// Init values to keep track of max value and the row
			maxI := 0
			maxValue := math.Inf(-1)
			// Loop through each row
			for i := range a.Shape[0] {

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
	}

	// Loop through each row
	for i := range a.Shape[0] {
		// Loop through each column
		for j := range a.Shape[1] {

			if direction > 0 {
				// Set result data using softmax
				jValue := argmaxMap[i]

				// If current j matches value stored in argmaxMap, then current element is max value for the column, so set value to 1, otherwise 0.
				if jValue == j {
					argMaxArei.SetIndex(1, i, j)
				}
			} else {

				// Set result data using softmax
				iValue := argmaxMap[j]

				// If current i matches value stored in argmaxMap, then current element is max value for the column, so set value to 1, otherwise 0.
				if iValue == i {
					argMaxArei.SetIndex(1, i, j)
				}
			}

		}
	}

	return argMaxArei
}

// SoftMax takes an arei, then returns an arei with with the softmax for each element in comparsion to the sum of the elemnet's column softmax
func SoftMax(a *Arei, direction int) *Arei {

	// Init A with same shape as  a
	softMaxArei, _ := Zeros(a.Shape)
	// Init slice to store data on which row-column should be 1
	softmaxMap := make([]float64, a.Shape[1])

	// Direction <= 0 to softmax across column, > 0 for across row
	if direction > 0 {
		// Loop through each column
		for i := range a.Shape[0] {
			// Init values to keep the sum of the softmax for the column
			sum := 0.0

			// Loop through each row
			for j := range a.Shape[1] {

				// Get the element
				element, _ := a.Index(i, j)
				// Add e^element to sum variable
				sum += math.Exp(element)
			}
			// After looping thorugh all columns in row, store row in map. Index of map denotes i
			softmaxMap[i] = sum
		}

		// Loop through each col
		for j := range a.Shape[1] {
			// Loop through each row
			for i := range a.Shape[0] {
				// Set result data using softmax
				element, _ := a.Index(i, j)
				softMaxSum := softmaxMap[i]
				eBaseElement := math.Exp(element)
				softMaxElement := eBaseElement / softMaxSum

				// Set softMaxElement to resultData
				softMaxArei.SetIndex(softMaxElement, i, j)

			}
		}
	} else {

		// Loop through each column
		for j := range a.Shape[1] {
			// Init values to keep the sum of the softmax for the column
			sum := 0.0

			// Loop through each row
			for i := range a.Shape[0] {

				// Get the element
				element, _ := a.Index(i, j)
				// Add e^element to sum variable
				sum += math.Exp(element)
			}
			// After looping thorugh all rows in column, store row in map. Index of map denotes j
			softmaxMap[j] = sum
		}

		// Loop through each row
		for i := range a.Shape[0] {
			// Loop through each column
			for j := range a.Shape[1] {
				// Set result data using softmax
				element, _ := a.Index(i, j)
				softMaxSum := softmaxMap[j]
				eBaseElement := math.Exp(element)
				softMaxElement := eBaseElement / softMaxSum

				// Set softMaxElement to resultData
				softMaxArei.SetIndex(softMaxElement, i, j)

			}
		}
	}

	return softMaxArei
}

// RowWiseSum takes an arei then returns a column vector of the sums of each row

func RowWiseSum(a *Arei) *Arei {
	// Init map to hold data
	data := make([]float64, a.Shape[0])
	// Loop through each row
	for i := range a.Shape[0] {
		// Init sum for each row
		rowSum := 0.0
		// Loop through each column
		for j := range a.Shape[1] {
			element, _ := a.Index(i, j)
			rowSum += element
		}
		// After looping through all columns in row, add sum to data map
		data[i] = rowSum
	}

	// Return a Arei with all data with mx1 shape
	return &Arei{
		Shape: []int{a.Shape[0], 1},
		Data:  data,
	}
}
