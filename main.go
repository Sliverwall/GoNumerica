package main

import (
	"log"
)

// Define helper functions
func matrixSum(A, B [][]int) [][]int {
	// Element wise sum two matrices
	C := [][]int{
		{0, 0},
		{0, 0},
	}
	m, n := 2, 2
	// Perform matrix sum
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			C[i][j] = A[i][j] + B[i][j]
		}
	}
	return C
}
func matrixProduct(A, B [][]int) [][]int {

	C := [][]int{
		{0, 0},
		{0, 0},
	}
	m, n := 2, 2
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
	return C
}
func fib(n int) [][]int {

	fibMatrix := [][]int{
		{1, 1},
		{1, 0},
	}

	placeHolder := [][]int{
		{0, 0},
		{0, 0},
	}
	for i := 0; i < n-1; i++ {
		C := matrixProduct(fibMatrix, fibMatrix)

		placeHolder = matrixSum(C, placeHolder)
	}

	return placeHolder

}
func main() {
	log.Println("Hello GoNumerica")

	// Print examples
	log.Println(fib(1))
}
