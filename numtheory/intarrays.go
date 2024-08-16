package numtheory

// Module to work on int arrays. These are indepedent, but may overlap with arei functions
func MatrixProduct(A, B [][]int) [][]int {

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
func MatrixPow(A [][]int, n int) [][]int {
	// Base case: A^1 is A
	if n == 1 {
		return A
	}

	// Recursively compute A^(n/2)
	halfPower := MatrixPow(A, n/2)
	// Multiply A^(n/2) by itself
	B := MatrixProduct(halfPower, halfPower)

	// If n is odd, multiply by A one more time
	if n%2 != 0 {
		B = MatrixProduct(B, A)
	}

	return B
}
