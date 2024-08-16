package arei

import (
	"errors"
	"fmt"
	"math"
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

// MatrixPow computes A^n (matrix A raised to the power n)
func MatrixPow(A *Arei, n int) (*Arei, error) {
	// Base case: A^1 is A
	if n == 1 {
		return A, nil
	}

	// Recursively compute A^(n/2)
	halfPower, err := MatrixPow(A, n/2)
	if err != nil {
		return nil, err
	}

	// Multiply A^(n/2) by itself
	result, err := MatrixProduct(halfPower, halfPower)
	if err != nil {
		return nil, err
	}

	// If n is odd, multiply by A one more time
	if n%2 != 0 {
		result, err = MatrixProduct(result, A)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Elimination performs Gaussian elimination on the given Arei and returns L, U, P, and number of row swaps
func Elimination(a *Arei) (*Arei, *Arei, *Arei, int, error) {
	// Check if the input is a 2D Arei
	if len(a.Shape) != 2 {
		return nil, nil, nil, 0, errors.New("arei must be a 2D matrix")
	}

	// Create a copy of the input Arei
	u, err := a.Copy()
	if err != nil {
		return nil, nil, nil, 0, fmt.Errorf("failed to copy arei: %w", err)
	}

	rows, cols := u.Shape[0], u.Shape[1]

	// Initialize L and P as identity matrices
	l, err := Identity(u.Shape)
	if err != nil {
		return nil, nil, nil, 0, fmt.Errorf("failed to create identity matrix: %w", err)
	}

	p, err := Identity(u.Shape)
	if err != nil {
		return nil, nil, nil, 0, fmt.Errorf("failed to create identity matrix: %w", err)
	}

	// Track the number of row swaps
	rowSwaps := 0

	// Perform Gaussian elimination with row swapping
	for i := 0; i < rows-1; i++ {
		// Pivot element
		pivot, _ := u.Index(i, i)

		// If the pivot is zero, find a row below to swap
		if pivot == 0 {
			rowSwapped := false
			for j := i + 1; j < rows; j++ {
				newPivot, _ := u.Index(j, i)
				if newPivot != 0 {
					// Swap rows i and j in the U matrix
					u.SwapRows(i, j)
					// Swap rows i and j in the L matrix (before column i)
					l.SwapRows(i, j)
					// Swap rows i and j in the permutation matrix P
					p.SwapRows(i, j)
					rowSwaps++ // Increment row swap counter
					pivot = newPivot
					rowSwapped = true
					break
				}
			}
			// If no row was found to swap, the matrix might be singular
			if !rowSwapped {
				return nil, nil, p, rowSwaps, errors.New("cannot perform elimination, singular matrix detected")
				// Any two rows or columns are identical
				// All elements in a row or column are zero
			}
		}

		for j := i + 1; j < rows; j++ {
			// Calculate the factor to eliminate the current row
			factor, _ := u.Index(j, i)
			factor = factor / pivot

			// Store the factor in L
			l.SetIndex(factor, j, i)

			// Subtract the scaled pivot row from the current row
			for k := i; k < cols; k++ {
				// Rx = Rx - (factor * Ri)
				Rx_k, _ := u.Index(j, k)
				Ri_k, _ := u.Index(i, k)
				newValue := Rx_k - factor*Ri_k
				u.SetIndex(newValue, j, k)
			}
		}
	}

	return l, u, p, rowSwaps, nil
}

// Determinant calculates the determinant of a matrix using Gaussian elimination
func Determinant(a *Arei) (float64, error) {
	_, u, _, rowSwaps, err := Elimination(a)
	if err != nil {
		return 0, err
	}

	// Calculate the product of the diagonal elements of U
	det := 1.0
	rows := u.Shape[0]

	for i := 0; i < rows; i++ {
		diagonalElement, _ := u.Index(i, i)
		det = det * diagonalElement
	}

	// Adjust the sign of the determinant based on the number of row swaps
	if rowSwaps%2 != 0 {
		det = -det
	}

	return det, nil
}

// Cofactor returns the cofactor matrix of a given matrix
func Cofactor(a *Arei) (*Arei, error) {
	// Check if the input is a 2D Arei
	if len(a.Shape) != 2 {
		return nil, errors.New("arei must be a 2D matrix")
	}

	m, n := a.Shape[0], a.Shape[1]

	// Initialize a new Arei for the cofactor matrix
	cofactorData := make([]float64, m*n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// Compute the minor by removing row i and column j
			minor, err := RemoveRow(a, i)
			if err != nil {
				return nil, err
			}
			minor, err = RemoveColumn(minor, j)
			if err != nil {
				return nil, err
			}

			// Compute the determinant of the minor matrix
			det, err := Determinant(minor)
			if err != nil {
				return nil, err
			}

			// Calculate the cofactor value
			sign := math.Pow(-1, float64(i+j))
			cofactorData[i*n+j] = sign * det
		}
	}

	// Return the cofactor matrix
	return &Arei{
		Shape: []int{m, n},
		Data:  cofactorData,
	}, nil
}

// Inverse takes a 2D arei and returns its inverse, if possible
func Inverse(a *Arei) (*Arei, error) {
	// A^-1 = 1/|A| * CT

	// Get the determinant of a
	det, err := Determinant(a)
	if err != nil {
		return nil, err
	}

	// Check if determinant is 0
	if det == 0 {
		return nil, errors.New("arei cannot be inverted due to determinant being 0")
	}
	// Get the cofactor of a
	c, err := Cofactor(a)
	if err != nil {
		return nil, err
	}

	// Tranpose cofactor
	c.Transpose()

	// Multiply each element of the transposed cofactor by 1/determinant of a
	inverse := MultiT(c, 1/det)

	// Return inverse
	return inverse, nil

}
