package arei

import (
	"errors"
	"fmt"
)

// Functions to create special kind of matrixes

// Ns creates an Arei of n value based on the provided shape. Defaults to 0 if no value given
func Ns(shape []int, n ...float64) (*Arei, error) {
	if len(shape) == 0 {
		return nil, errors.New("shape cannot be empty")
	}
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
	if len(shape) == 0 {
		return nil, errors.New("shape cannot be empty")
	}
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

// Permutation outputs for a given matrix of shape m x n, a described perumutation
func Permutation(shape []int, insturction [][]int) (*Arei, error) {
	if len(shape) == 0 {
		return nil, errors.New("shape cannot be empty")
	}
	row := shape[0]
	col := shape[1]
	m_m := row == 1 && col > 0
	m_n := row > 1 && col > 0
	if m_m {
		// Return error for vector (scalar 1)
		return nil, errors.New("cannot permutate a vector")
	} else if m_n {
		// Use the minimum of rows and columns to form a square identity matrix
		// The number of rows determines the identity matrix size
		identityData := make([]float64, row*row)
		for i := 0; i < row; i++ {
			identityData[i*row+i] = 1.0
		}
		// new shape will be the size_size
		finalShape := []int{row, row}
		// Form base permutation as a identity matrix
		P := &Arei{Shape: finalShape, Data: identityData}

		// Loop through the instruction array
		for i := 0; i < len(insturction); i++ {
			// Through through each instruction. Each row should only have 2 cols
			if len(insturction[i]) != 2 {
				return nil, errors.New("instruction must be of length 2, i.e. {i,j}")
			}

			replaceRow := insturction[i][0]
			withRow := insturction[i][1]
			// Allow negative indexing
			if replaceRow < 0 {
				replaceRow += P.Shape[0]
			}
			if withRow < 0 {
				withRow += P.Shape[0]
			}

			// Check if instruction is out of bounds
			if replaceRow < 0 || replaceRow >= P.Shape[0] || withRow < 0 || withRow >= P.Shape[0] {
				errorMsg := fmt.Sprintf("instruction index out of bounds {%d,%d}", insturction[i][0], insturction[i][1])
				return nil, errors.New(errorMsg)
			}
			// Set replacement row to all 0
			for k := 0; k < row; k++ {
				P.SetIndex(0, replaceRow, k)
			}

			// Set replacement location to 1
			P.SetIndex(1, replaceRow, withRow)
		}

		// Return nil for now
		return P, nil
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
