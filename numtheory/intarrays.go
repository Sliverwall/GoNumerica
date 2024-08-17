package numtheory

import (
	"errors"
	"math/big"
)

// Module to work on int arrays. These are indepedent, but may overlap with arei functions

// MatrixProduct multiplies [][]int matrixes
func MatrixProduct(A, B [][]int) ([][]int, error) {

	m, n := len(A), len(A[0])
	p, o := len(B), len(B[0])
	if m != p || n != o {
		return nil, errors.New("both matrices must have the same shape")
	}
	C := make([][]int, m)
	for i := range C {
		C[i] = make([]int, n)
	}
	// Perform matrix multiplication
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum := 0
			for k := 0; k < n; k++ {
				sum += A[i][k] * B[k][j]
			}
			C[i][j] = sum
		}
	}
	return C, nil
}

// MatrixPow raises a matrix to the nth power A^n
func MatrixPow(A [][]int, n int) [][]int {
	// Base case: A^1 is A
	if n == 1 {
		return A
	}

	// Recursively compute A^(n/2)
	halfPower := MatrixPow(A, n/2)
	// Multiply A^(n/2) by itself
	B, _ := MatrixProduct(halfPower, halfPower)

	// If n is odd, multiply by A one more time
	if n%2 != 0 {
		B, _ = MatrixProduct(B, A)
	}

	return B
}

// MatrixProductBig multiplies [][]*big.Int matrices
func MatrixProductBig(A, B [][]*big.Int) ([][]*big.Int, error) {
	m, n := len(A), len(A[0])
	p, o := len(B), len(B[0])
	if n != p {
		return nil, errors.New("number of columns in A must equal the number of rows in B")
	}

	// Initialize result matrix
	C := make([][]*big.Int, m)
	for i := range C {
		C[i] = make([]*big.Int, o)
		for j := range C[i] {
			C[i][j] = new(big.Int)
		}
	}

	// Perform matrix multiplication
	for i := 0; i < m; i++ {
		for j := 0; j < o; j++ {
			for k := 0; k < n; k++ {
				// C[i][j] += A[i][k] * B[k][j]
				C[i][j].Add(C[i][j], new(big.Int).Mul(A[i][k], B[k][j]))
			}
		}
	}

	return C, nil
}

// MatrixPowBig raises a matrix of big.Int values to the nth power
func MatrixPowBig(A [][]*big.Int, n int) [][]*big.Int {
	// Base case: A^1 is A
	if n == 1 {
		return A
	}

	// Recursively compute A^(n/2)
	halfPower := MatrixPowBig(A, n/2)

	// Multiply A^(n/2) by itself
	B, _ := MatrixProductBig(halfPower, halfPower)

	// If n is odd, multiply by A one more time
	if n%2 != 0 {
		B, _ = MatrixProductBig(B, A)
	}

	return B
}
