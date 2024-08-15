package arei

import (
	"errors"
	"fmt"
	"math"
)

// Search for collection of functions that search Areis

// Where returns the indices of elements that satisfy the given condition.
func Where(a *Arei, condition func(float64) bool) *Arei {
	var numRows, numCols int

	if len(a.Shape) == 2 {
		numRows = a.Shape[0]
		numCols = a.Shape[1]
	} else if len(a.Shape) == 1 {
		numRows = 1
		numCols = a.Shape[0]
	}

	// Initialize an empty slice to store the indices of max elements for each row
	indices := make([][]float64, 0)
	// Iterate over each row
	for i := 0; i < numRows; i++ {
		// Iterate over each column
		for j := 0; j < numCols; j++ {
			element := a.Data[i*numCols+j]
			if condition(element) {
				hold := make([]float64, 2)
				hold[0] = float64(i)
				hold[1] = float64(j)
				indices = append(indices, hold)
			}
		}
	}

	// Return the indices as a new Arei
	resultArei, _ := NewArei(indices)
	return resultArei
}

// WhereMax finds the indices of the maximum value along each row of an arei.
func WhereMax(a *Arei) *Arei {

	var numRows, numCols int

	if len(a.Shape) == 2 {
		numRows = a.Shape[0]
		numCols = a.Shape[1]
	} else if len(a.Shape) == 1 {
		numRows = 1
		numCols = a.Shape[0]
	}

	// Initialize an empty slice to store the indices of max elements for each row
	indices := make([]float64, numRows)

	// Iterate over each row
	for i := 0; i < numRows; i++ {
		maxValue := math.Inf(-1)
		maxIndex := 0
		// Iterate over each column in the current row to find the max value
		for j := 0; j < numCols; j++ {
			element := a.Data[i*numCols+j]
			if element > maxValue {
				maxValue = element
				maxIndex = j
			}
		}
		// Store the index of the max element for the current row
		indices[i] = float64(maxIndex)
	}

	// Return the indices as a new Arei
	resultIndices, _ := NewArei(indices)
	return resultIndices
}

// WhereMax finds the indices of the maximum value along each row of an arei.
func WhereMin(a *Arei) *Arei {

	var numRows, numCols int

	if len(a.Shape) == 2 {
		numRows = a.Shape[0]
		numCols = a.Shape[1]
	} else if len(a.Shape) == 1 {
		numRows = 1
		numCols = a.Shape[0]
	}

	// Initialize an empty slice to store the indices of max elements for each row
	indices := make([]float64, numRows)

	// Iterate over each row
	for i := 0; i < numRows; i++ {
		minValue := math.Inf(1)
		minIndex := 0
		// Iterate over each column in the current row to find the min value
		for j := 0; j < numCols; j++ {
			element := a.Data[i*numCols+j]
			if element > minValue {
				minValue = element
				minIndex = j
			}
		}
		// Store the index of the min element for the current row
		indices[i] = float64(minIndex)
	}

	// Return the indices as a new Arei
	resultIndices, _ := NewArei(indices)
	return resultIndices
}

// Row returns a specified row, by index, of an aeri as a 1D aeri
func Row(a *Arei, rowIndex int) (*Arei, error) {
	// 1D areis cannot be searched by row
	if len(a.Shape) == 1 {
		return nil, errors.New("1d aeri only have 1 row")
	}

	// Check for valid length
	if math.Abs(float64(rowIndex)) > float64(a.Shape[0]) {
		return nil, errors.New("index out of bounds")
	}
	// if negative index, count backwards
	if rowIndex < 0 {
		// Shape is not 0-indexed, thus negative rowIndex + shape will give backward result
		rowIndex += a.Shape[0]
	}

	resultData := make([]float64, a.Shape[1])

	for i := range a.Shape[1] {
		value, _ := a.Index(rowIndex, i)
		resultData[i] = value
	}
	return NewArei(resultData)
}

// SwapRows swaps two rows in the Arei matrix
func (a *Arei) SwapRows(row1, row2 int) error {
	if len(a.Shape) != 2 {
		return errors.New("swapRows can only be used on 2D matrices")
	}

	cols := a.Shape[1]
	for col := 0; col < cols; col++ {
		// Swap elements in row1 and row2 for all columns
		temp := a.Data[row1*cols+col]
		a.Data[row1*cols+col] = a.Data[row2*cols+col]
		a.Data[row2*cols+col] = temp
	}
	return nil
}

// RemoveRow removes the specified row from the Arei and returns a new Arei without that row.
func RemoveRow(a *Arei, rowIndex int) (*Arei, error) {
	// Ensure that a is a 2D matrix
	if len(a.Shape) != 2 {
		return nil, errors.New("arei must be a 2D matrix")
	}

	rows, cols := a.Shape[0], a.Shape[1]

	// Check if rowIndex is out of bounds
	if rowIndex < 0 || rowIndex >= rows {
		return nil, fmt.Errorf("row index %d is out of bounds", rowIndex)
	}

	// Create a new data slice for the result
	newData := make([]float64, 0, (rows-1)*cols)

	// Copy all rows except the row at rowIndex
	for i := 0; i < rows; i++ {
		if i != rowIndex {
			newData = append(newData, a.Data[i*cols:(i+1)*cols]...)
		}
	}

	// Return the new Arei with one fewer row
	return &Arei{
		Shape: []int{rows - 1, cols},
		Data:  newData,
	}, nil
}

// Column returns a specified column, by index, of an aeri as a 1D aeri
func Column(a *Arei, colIndex int) (*Arei, error) {
	// 1D areis cannot be searched by row
	if len(a.Shape) == 1 {
		return nil, errors.New("1d aeri only have 1 row")
	}

	// Check for valid length
	if math.Abs(float64(colIndex)) > float64(a.Shape[1]) {
		return nil, errors.New("index out of bounds")
	}
	// if negative index, count backwards
	if colIndex < 0 {
		// Shape is not 0-indexed, thus negative colIndex + shape will give backward result
		colIndex += a.Shape[0]
	}

	resultData := make([]float64, a.Shape[1])

	for i := range a.Shape[0] {
		value, _ := a.Index(i, colIndex)
		resultData[i] = value
	}
	return NewArei(resultData)
}

// RemoveColumn removes the specified column from the Arei and returns a new Arei without that column.
func RemoveColumn(a *Arei, colIndex int) (*Arei, error) {
	// Ensure that a is a 2D matrix
	if len(a.Shape) != 2 {
		return nil, errors.New("arei must be a 2D matrix")
	}

	rows, cols := a.Shape[0], a.Shape[1]

	// Check if colIndex is out of bounds
	if colIndex < 0 || colIndex >= cols {
		return nil, fmt.Errorf("column index %d is out of bounds", colIndex)
	}

	// Create a new data slice for the result
	newData := make([]float64, 0, rows*(cols-1))

	// Copy all columns except the column at colIndex
	for i := 0; i < rows; i++ {
		start := i * cols
		newData = append(newData, a.Data[start:start+colIndex]...)
		newData = append(newData, a.Data[start+colIndex+1:start+cols]...)
	}

	// Return the new Arei with one fewer column
	return &Arei{
		Shape: []int{rows, cols - 1},
		Data:  newData,
	}, nil
}
