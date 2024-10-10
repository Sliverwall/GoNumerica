package arei

import "errors"

// Functions that produce scalars from arei inputs

// SumElements takes an arei then sums all its elements then returns the scalar sum
func SumElements(a *Arei) float64 {
	// Init variable to hold sum term
	sum := 0.0

	// Loop through each row
	for i := range a.Shape[0] {
		// Loop through each column
		for j := range a.Shape[1] {
			// Get element from arei
			element, _ := a.Index(i, j)

			// Add element to sum total
			sum += element
		}
	}
	// Return the total sum
	return sum
}

// ProdElements takes an arei then multiplies all its elements then returns the scalar product
func ProdElements(a *Arei) float64 {
	// Init variable to hold prod term
	prod := 1.0

	// Loop through each row
	for i := range a.Shape[0] {
		// Loop through each column
		for j := range a.Shape[1] {
			// Get element from arei
			element, _ := a.Index(i, j)

			// multiply element to prod total
			prod *= element
		}
	}
	// Return the total prod
	return prod
}

// Trace takes a sqaure matrix, then returns the sum of the diagonal elements

func Trace(a *Arei) (t float64, err error) {
	// Check if matrix is not sqaure
	if a.Shape[0] != a.Shape[1] {
		// matrix is not sqaure, return error
		return 0.0, errors.New("only sqaure matricies have a trace")
	}

	// Init variable to hold trace
	trace := 0.0

	// Loop through each column
	for j := range a.Shape[1] {
		// Get element from column's diagonal element
		element, _ := a.Index(j, j)

		// Sum elements to get trace
		trace += element
	}

	return trace, nil
}
