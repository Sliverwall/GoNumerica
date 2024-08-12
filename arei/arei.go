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

// sameShape checks if two Areis have the same shape.
func (a *Arei) sameShape(other *Arei) bool {
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

// Transpose transposes a Areis. Works for both 1D and 2D
func (fa *Arei) Transpose() {

	transposedData := make([]float64, len(fa.data))
	rows, cols := fa.shape[0], fa.shape[1]

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposedData[j*rows+i] = fa.data[i*cols+j]
		}
	}
	fa.shape = []int{cols, rows}
	fa.data = transposedData
}
