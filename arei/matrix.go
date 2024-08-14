package arei

import (
	"errors"
	"fmt"
)

// MatrixProduct performs matrix multiplication of two Arei representing matrices.
func MatrixProduct(A, B *Arei) (*Arei, error) {
	// Check that A and B are 2D areis
	if len(A.Shape) != 2 || len(B.Shape) != 2 {
		return nil, errors.New("both areis must be 2D for matrix multiplication")
	}

	numARows, numACols := A.Shape[0], A.Shape[1]
	numBRows, numBCols := B.Shape[0], B.Shape[1]

	// Check if matrix dimensions are compatible for multiplication
	if numACols != numBRows {
		return nil, errors.New("number of columns in A must be equal to the number of rows in B")
	}

	// Initialize result matrix C with zeros
	C := &Arei{
		Shape: []int{numARows, numBCols},
		Data:  make([]float64, numARows*numBCols),
	}

	// Perform matrix multiplication
	for i := 0; i < numARows; i++ {
		for j := 0; j < numBCols; j++ {
			sum := 0.0
			for k := 0; k < numACols; k++ {
				sum += A.Data[i*numACols+k] * B.Data[k*numBCols+j]
			}
			C.Data[i*numBCols+j] = sum
		}
	}

	return C, nil
}

// Elimination performs Gaussian elimination on the given Arei and returns L and U
func Elimination(a *Arei) (*Arei, *Arei, error) {
	// Check if the input is a 2D Arei
	if len(a.Shape) != 2 {
		return nil, nil, errors.New("arei must be a 2D matrix")
	}

	// Create a copy of the input Arei
	u, err := a.Copy()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to copy arei: %w", err)
	}

	rows, cols := u.Shape[0], u.Shape[1]
	if rows > cols {
		return nil, nil, errors.New("cannot perform elimination on a tall matrix")
	}

	// Initialize L as an identity matrix
	l, err := Identity(u.Shape)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create identity matrix: %w", err)
	}

	// Perform Gaussian elimination
	for i := 0; i < rows-1; i++ {
		// Pivot element
		pivot, _ := u.Index(i, i)
		if pivot == 0 {
			return nil, nil, errors.New("pivot is zero, cannot eliminate")
		}

		for j := i + 1; j < rows; j++ {
			// Calculate the factor to eliminate the current row
			factor, _ := u.Index(j, i)
			factor = factor / pivot

			// Store the factor in L
			l.SetIndex(factor, j, i)

			// Subtract the scaled pivot row from the current row
			for k := i; k < cols; k++ {
				// Rx = Rx - (factorRx * Ri)
				Rx_k, _ := u.Index(j, k)
				Ri_k, _ := u.Index(i, k)
				newValue := Rx_k - factor*Ri_k
				u.SetIndex(newValue, j, k)
			}
		}
	}

	return l, u, nil
}
