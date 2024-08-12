package arei

import (
	"errors"
	"fmt"
)

// Arei represents a multi-dimensional array, including vectors and matrices, for float64 data.
type Arei struct {
	shape []int
	data  []float64
}

// NewArei creates a new Arei based on the provided data.
// It automatically determines the shape based on the type of data.
func NewArei(data interface{}) (*Arei, error) {
	var shape []int
	var flatData []float64

	switch v := data.(type) {
	case []float64:
		shape = []int{len(v)}
		flatData = v
	case [][]float64:
		if len(v) == 0 {
			return nil, errors.New("data cannot be empty")
		}
		shape = []int{len(v), len(v[0])}
		flatData = make([]float64, 0, shape[0]*shape[1])
		for _, row := range v {
			if len(row) != shape[1] {
				return nil, errors.New("all rows must have the same number of columns")
			}
			flatData = append(flatData, row...)
		}
	default:
		return nil, errors.New("unsupported data type")
	}

	return &Arei{
		shape: shape,
		data:  flatData,
	}, nil
}

// String returns a string representation of the Arei.
func (a *Arei) String() string {
	switch len(a.shape) {
	// String for vector form
	case 1:
		return fmt.Sprintf("Vector: %v", a.data)
	// String for Matrix form.
	case 2:
		rows, cols := a.shape[0], a.shape[1]
		result := "\n"
		for i := 0; i < rows; i++ {
			// Avoid extra space at final row
			if i+1 == rows {
				result += fmt.Sprintf("%v", a.data[i*cols:(i+1)*cols])
			} else {
				result += fmt.Sprintf("%v\n", a.data[i*cols:(i+1)*cols])
			}
		}
		return result
	default:
		return "Arei of unsupported dimension"
	}
}

// SameShape checks if two Areis have the same shape.
func (a *Arei) SameShape(other *Arei) bool {
	if len(a.shape) != len(other.shape) {
		return false
	}

	for i, dim := range a.shape {
		if dim != other.shape[i] {
			return false
		}
	}

	return true
}

// Reshape changes the shape of the Arei, keeping the data intact.
func (a *Arei) Reshape(newShape []int) error {
	size := 1
	for _, dim := range newShape {
		size *= dim
	}

	if size != len(a.data) {
		return errors.New("new shape must have the same number of elements as the original")
	}

	a.shape = newShape
	return nil
}

// Transpose transposes a Areis. Works for both 1D and 2D
func (a *Arei) Transpose() {

	transposedData := make([]float64, len(a.data))
	var rows, cols int
	if len(a.shape) == 2 {
		rows, cols = a.shape[0], a.shape[1]
	} else if len(a.shape) == 1 {
		rows, cols = 1, a.shape[0]
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposedData[j*rows+i] = a.data[i*cols+j]
		}
	}
	a.shape = []int{cols, rows}
	a.data = transposedData

}

// Flatten takes an Arei and forces it to be 1xN
func (a *Arei) Flatten() error {
	// Check if vector
	if len(a.shape) == 1 {
		return errors.New("cannot flatten a 1D arei")
	}
	transposedData := make([]float64, len(a.data))
	// Rows must be 1 and cols will equal the current rows * cols
	rows, cols := 1, (a.shape[0] * a.shape[1])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposedData[j*rows+i] = a.data[i*cols+j]
		}
	}
	a.shape = []int{rows, cols}
	a.data = transposedData
	return nil
}

// Count returns the number of elements in the arei
func (a *Arei) Count() int {
	var result int
	if len(a.shape) == 1 {
		result = a.shape[0]
	} else {
		result = a.shape[0] * a.shape[1]
	}
	return result
}
