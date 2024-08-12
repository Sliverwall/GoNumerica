package arei

import (
	"errors"
	"fmt"
	"math"
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
	// String representation for a vector (1D array).
	case 1:
		return fmt.Sprintf("%v", a.data)
	// String representation for a matrix (2D array).
	case 2:
		rows, cols := a.shape[0], a.shape[1]
		result := "["
		for i := 0; i < rows; i++ {
			// Slice the data to get the current row, and format it.
			row := a.data[i*cols : (i+1)*cols]
			result += fmt.Sprintf("%v", row)
			if i < rows-1 {
				result += "," // Add a comma and newline between rows.
			}
		}
		result += "]"
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

// Transpose takes each row in the Arei and repositions it as a column
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
	return len(a.data)
}

// Index returns the element at the specified index for 1D Areis or at the specified row and column for 2D Areis.
func (a *Arei) Index(indices ...int) (float64, error) {
	switch len(a.shape) {
	case 1:
		if len(indices) != 1 {
			return 0, errors.New("1D Arei requires exactly 1 index")
		}
		if indices[0] < 0 || indices[0] >= a.shape[0] {
			return 0, errors.New("index out of bounds")
		}
		return a.data[indices[0]], nil
	case 2:
		if len(indices) != 2 {
			return 0, errors.New("2D Arei requires exactly 2 indices")
		}
		row, col := indices[0], indices[1]
		if row < 0 || row >= a.shape[0] || col < 0 || col >= a.shape[1] {
			return 0, errors.New("index out of bounds")
		}
		return a.data[row*a.shape[1]+col], nil
	default:
		return 0, errors.New("Arei of unsupported dimension")
	}
}

// SetIndex sets the element at the specified index for 1D Areis or at the specified row and column for 2D Areis.
func (a *Arei) SetIndex(value float64, indices ...int) error {
	switch len(a.shape) {
	case 1:
		if len(indices) != 1 {
			return errors.New("1D Arei requires exactly 1 index")
		}
		if indices[0] < 0 || indices[0] >= a.shape[0] {
			return errors.New("index out of bounds")
		}
		a.data[indices[0]] = value
		return nil
	case 2:
		if len(indices) != 2 {
			return errors.New("2D Arei requires exactly 2 indices")
		}
		row, col := indices[0], indices[1]
		if row < 0 || row >= a.shape[0] || col < 0 || col >= a.shape[1] {
			return errors.New("index out of bounds")
		}
		a.data[row*a.shape[1]+col] = value
		return nil
	default:
		return errors.New("Arei of unsupported dimension")
	}
}

// Max finds the max element in an Arei
func (a *Arei) Max() (float64, error) {

	var result float64 = math.Inf(-1)
	for i := range a.data {
		if a.data[i] > result {
			result = a.data[i]
		}
	}
	return result, nil
}

// Min finds the min element in an Arei
func (a *Arei) Min() (float64, error) {

	var result float64 = math.Inf(1)
	for i := range a.data {
		if a.data[i] < result {
			result = a.data[i]
		}
	}
	return result, nil
}
