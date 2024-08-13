package arei

import "errors"

// Functions to create special kind of matrixes

// Ns creates a Arei of n value based on the provided Shape. Defaults to 0 if no value given
func Ns(Shape []int, n ...float64) (*Arei, error) {

	// Default n value to 0
	nValue := 0.0

	// If n value provided, set it to nValue
	if len(n) > 0 {
		nValue = n[0]
	}

	if len(Shape) == 0 {
		return nil, errors.New("shape cannot be empty")
	}

	size := 1
	for _, dim := range Shape {
		if dim <= 0 {
			return nil, errors.New("Shape dimensions must be positive integers")
		}
		size *= dim
	}

	// Create zeros Data place holder
	Data := make([]float64, size)

	if nValue != 0.0 {
		for i := range Data {
			Data[i] = nValue
		}
	}

	return &Arei{Shape: Shape, Data: Data}, nil
}

// Identity outputs for a given matrix of shape m x n, an m x m identity matrix.
func Identity(a *Arei) (*Arei, error) {
	switch len(a.Shape) {
	case 1:
		// Identity vector (scalar 1)
		identityData := make([]float64, a.Shape[0])
		for i := range identityData {
			identityData[i] = 1.0
		}
		return &Arei{Shape: a.Shape, Data: identityData}, nil
	case 2:
		// Use the minimum of rows and columns to form a square identity matrix
		size := a.Shape[0] // The number of rows determines the identity matrix size
		identityData := make([]float64, size*size)
		for i := 0; i < size; i++ {
			identityData[i*size+i] = 1.0
		}
		return &Arei{Shape: []int{size, size}, Data: identityData}, nil
	default:
		return nil, errors.New("invalid Shape for identity Arei")
	}
}

// Zeros creates a Arei of zeros based on the provided Shape.
func Zeros(Shape []int) (*Arei, error) {
	if len(Shape) == 0 {
		return nil, errors.New("Shape cannot be empty")
	}

	size := 1
	for _, dim := range Shape {
		if dim <= 0 {
			return nil, errors.New("Shape dimensions must be positive integers")
		}
		size *= dim
	}

	zeroData := make([]float64, size)
	return &Arei{Shape: Shape, Data: zeroData}, nil
}
