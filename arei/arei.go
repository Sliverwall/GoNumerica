package arei

import (
	"errors"
	"fmt"
	"math"
)

// Arei represents a multi-dimensional array, including vectors and matrices, for float64 Data.
type Arei struct {
	Shape []int
	Data  []float64
}

// NewArei creates a new Arei based on the provided Data.
func NewArei(Data interface{}) (*Arei, error) {
	var Shape []int
	var flatData []float64

	switch v := Data.(type) {
	case []float64:
		Shape = []int{len(v)}
		flatData = v
	case [][]float64:
		if len(v) == 0 {
			return nil, errors.New("data cannot be empty")
		}
		Shape = []int{len(v), len(v[0])}
		flatData = make([]float64, 0, Shape[0]*Shape[1])
		for _, row := range v {
			if len(row) != Shape[1] {
				return nil, errors.New("all rows must have the same number of columns")
			}
			flatData = append(flatData, row...)
		}
	default:
		return nil, errors.New("unsupported Data type")
	}

	return &Arei{
		Shape: Shape,
		Data:  flatData,
	}, nil
}

// String returns a string representation of the Arei.
func (a *Arei) String() string {
	switch len(a.Shape) {
	// String representation for a vector (1D array).
	case 1:
		return fmt.Sprintf("%v", a.Data)
	// String representation for a matrix (2D array).
	case 2:
		rows, cols := a.Shape[0], a.Shape[1]
		result := "["
		for i := 0; i < rows; i++ {
			// Slice the Data to get the current row, and format it.
			row := a.Data[i*cols : (i+1)*cols]
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

// Frame prints an arei in matrix form
func (a *Arei) Frame() {
	switch len(a.Shape) {
	// String for vector form
	case 1:
		fmt.Printf("%v", a.Data)
	// String for Matrix form.
	case 2:
		rows, cols := a.Shape[0], a.Shape[1]
		result := ""
		for i := 0; i < rows; i++ {
			// Avoid extra space at final row
			if i+1 == rows {
				result += fmt.Sprintf("%v", a.Data[i*cols:(i+1)*cols])
			} else {
				result += fmt.Sprintf("%v\n", a.Data[i*cols:(i+1)*cols])
			}
		}
		fmt.Println(result)
	default:
		fmt.Println("Arei of unsupported dimension")
	}
}

// SameShape checks if two Areis have the same Shape.
func (a *Arei) SameShape(other *Arei) bool {
	if len(a.Shape) != len(other.Shape) {
		return false
	}

	for i, dim := range a.Shape {
		if dim != other.Shape[i] {
			return false
		}
	}

	return true
}

// Reshape changes the Shape of the Arei, keeping the Data intact.
func (a *Arei) Reshape(newShape []int) error {
	size := 1
	for _, dim := range newShape {
		size *= dim
	}

	if size != len(a.Data) {
		return errors.New("new Shape must have the same number of elements as the original")
	}

	a.Shape = newShape
	return nil
}

// Transpose takes each row in the Arei and repositions it as a column
func (a *Arei) Transpose() {
	transposedData := make([]float64, len(a.Data))
	var rows, cols int
	if len(a.Shape) == 2 {
		rows, cols = a.Shape[0], a.Shape[1]
	} else if len(a.Shape) == 1 {
		rows, cols = 1, a.Shape[0]
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposedData[j*rows+i] = a.Data[i*cols+j]
		}
	}
	a.Shape = []int{cols, rows}
	a.Data = transposedData

}

// Flatten takes an Arei and forces it to be 1xN
func (a *Arei) Flatten() error {
	// Check if vector
	if len(a.Shape) == 1 {
		return errors.New("cannot flatten a 1D arei")
	}
	transposedData := make([]float64, len(a.Data))
	// Rows must be 1 and cols will equal the current rows * cols
	rows, cols := 1, (a.Shape[0] * a.Shape[1])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposedData[j*rows+i] = a.Data[i*cols+j]
		}
	}
	a.Shape = []int{rows, cols}
	a.Data = transposedData
	return nil
}

// Count returns the number of elements in the arei
func (a *Arei) Count() int {
	return len(a.Data)
}

// Index returns the element at the specified index for 1D Areis or at the specified row and column for 2D Areis.
func (a *Arei) Index(indices ...int) (float64, error) {
	switch len(a.Shape) {
	case 1:
		if len(indices) != 1 {
			return 0, errors.New("1D Arei requires exactly 1 index")
		}
		if indices[0] < 0 || indices[0] >= a.Shape[0] {
			return 0, errors.New("index out of bounds")
		}
		return a.Data[indices[0]], nil
	case 2:
		if len(indices) != 2 {
			return 0, errors.New("2D Arei requires exactly 2 indices")
		}

		// Allow negative indexing
		if indices[0] < 0 {
			indices[0] += a.Shape[0]
		}
		if indices[1] < 0 {
			indices[1] += a.Shape[1]
		}
		row, col := indices[0], indices[1]
		if row < 0 || row >= a.Shape[0] || col < 0 || col >= a.Shape[1] {
			return 0, errors.New("index out of bounds")
		}
		return a.Data[row*a.Shape[1]+col], nil
	default:
		return 0, errors.New("Arei of unsupported dimension")
	}
}

// SetIndex sets the element at the specified index for 1D Areis or at the specified row and column for 2D Areis.
func (a *Arei) SetIndex(value float64, indices ...int) error {
	switch len(a.Shape) {
	case 1:
		if len(indices) != 1 {
			return errors.New("1D Arei requires exactly 1 index")
		}
		if indices[0] < 0 || indices[0] >= a.Shape[0] {
			return errors.New("index out of bounds")
		}
		a.Data[indices[0]] = value
		return nil
	case 2:
		if len(indices) != 2 {
			return errors.New("2D Arei requires exactly 2 indices")
		}
		// Allow negative indexing
		if indices[0] < 0 {
			indices[0] += a.Shape[0]
		}
		if indices[1] < 0 {
			indices[1] += a.Shape[1]
		}
		row, col := indices[0], indices[1]
		if row < 0 || row >= a.Shape[0] || col < 0 || col >= a.Shape[1] {
			return errors.New("index out of bounds")
		}
		a.Data[row*a.Shape[1]+col] = value
		return nil
	default:
		return errors.New("Arei of unsupported dimension")
	}
}

// Max finds the max element in an Arei
func (a *Arei) Max() (float64, error) {

	var result float64 = math.Inf(-1)
	for i := range a.Data {
		if a.Data[i] > result {
			result = a.Data[i]
		}
	}
	return result, nil
}

// Min finds the min element in an Arei
func (a *Arei) Min() (float64, error) {

	var result float64 = math.Inf(1)
	for i := range a.Data {
		if a.Data[i] < result {
			result = a.Data[i]
		}
	}
	return result, nil
}

// Copy creates a new arei with same shape and values as input
func (a *Arei) Copy() (*Arei, error) {
	if len(a.Shape) == 0 {
		return nil, errors.New("shape cannot be empty")
	}

	// Create a new slice for the data to ensure it's not shared
	dataCopy := make([]float64, len(a.Data))
	copy(dataCopy, a.Data)

	// Create a new slice for the shape to ensure it's not shared
	shapeCopy := make([]int, len(a.Shape))
	copy(shapeCopy, a.Shape)

	return &Arei{
		Shape: shapeCopy,
		Data:  dataCopy,
	}, nil
}
