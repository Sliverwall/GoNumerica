package arei

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
