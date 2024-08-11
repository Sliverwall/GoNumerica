package tensor

import (
	"errors"
	"fmt"
)

// Tensor represents a multi-dimensional array, including vectors and matrices, for float64 data.
type Tensor struct {
	shape []int
	data  []float64
}

// NewTensor creates a new Tensor based on the provided data.
// It automatically determines the shape based on the type of data.
func NewTensor(data interface{}) (*Tensor, error) {
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

	return &Tensor{
		shape: shape,
		data:  flatData,
	}, nil
}

// sameShape checks if two Tensors have the same shape.
func (t *Tensor) sameShape(other *Tensor) bool {
	if len(t.shape) != len(other.shape) {
		return false
	}

	for i, dim := range t.shape {
		if dim != other.shape[i] {
			return false
		}
	}

	return true
}

// Reshape changes the shape of the Tensor, keeping the data intact.
func (t *Tensor) Reshape(newShape []int) error {
	size := 1
	for _, dim := range newShape {
		size *= dim
	}

	if size != len(t.data) {
		return errors.New("new shape must have the same number of elements as the original")
	}

	t.shape = newShape
	return nil
}

// String returns a string representation of the Tensor.
func (t *Tensor) String() string {
	switch len(t.shape) {
	case 1:
		return fmt.Sprintf("Vector: %v", t.data)
	case 2:
		rows, cols := t.shape[0], t.shape[1]
		result := "\n"
		for i := 0; i < rows; i++ {
			result += fmt.Sprintf("%v\n", t.data[i*cols:(i+1)*cols])
		}
		return result
	default:
		return "Tensor of unsupported dimension"
	}
}
