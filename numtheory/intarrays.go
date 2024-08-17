package numtheory

import "errors"

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
