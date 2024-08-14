package arei

import (
	"errors"
)

// Functions to create special kind of matrixes

// Ns creates a Arei of n value based on the provided shape. Defaults to 0 if no value given
func Ns(shape []int, n ...float64) (*Arei, error) {

	// Default n value to 0
	nValue := 0.0

	// If n value provided, set it to nValue
	if len(n) > 0 {
		nValue = n[0]
	}

	if len(shape) == 0 {
		return nil, errors.New("shape cannot be empty")
	}

	size := 1
	for _, dim := range shape {
		if dim <= 0 {
			return nil, errors.New("shape dimensions must be positive integers")
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

	return &Arei{Shape: shape, Data: Data}, nil
}

// Identity outputs for a given matrix of shape m x n, an m x m identity matrix.
func Identity(shape []int) (*Arei, error) {

	row := shape[0]
	col := shape[1]
	m_m := row == 1 && col > 0
	m_n := row > 1 && col > 0
	if m_m {
		// Identity vector (scalar 1)
		identityData := make([]float64, row)
		for i := range identityData {
			identityData[i] = 1.0
		}
		return &Arei{Shape: shape, Data: identityData}, nil
	} else if m_n {
		// Use the minimum of rows and columns to form a square identity matrix
		// The number of rows determines the identity matrix size
		identityData := make([]float64, row*row)
		for i := 0; i < row; i++ {
			identityData[i*row+i] = 1.0
		}
		// new shape will be the size_size
		finalShape := []int{row, row}
		return &Arei{Shape: finalShape, Data: identityData}, nil
	} else {
		return nil, errors.New("invalid shape for identity arei")
	}
}

// Zeros creates a Arei of zeros based on the provided shape.
func Zeros(shape []int) (*Arei, error) {
	if len(shape) == 0 {
		return nil, errors.New("shape cannot be empty")
	}

	size := 1
	for _, dim := range shape {
		if dim <= 0 {
			return nil, errors.New("shape dimensions must be positive integers")
		}
		size *= dim
	}

	zeroData := make([]float64, size)
	return &Arei{Shape: shape, Data: zeroData}, nil
}
