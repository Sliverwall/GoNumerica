package notes

import (
	"log"
)

func MatrixProduct(A, B [][]float64) [][]float64 {
	// Try to get A*B
	numACols := len(A[0])
	numARows := len(A)

	numBRows := len(B)
	numBCols := len(B[0])

	// Step 1, check if shape works
	if numACols == numBRows {

		log.Println("Shape works, A cols = B rows. A cols = ", numACols, " B rows = ", numBRows)
		log.Println("A rows = ", numARows, " B cols = ", numBCols)
	} else {
		return [][]float64{{0}}

	}
	// Step 2 result matrix C with zeros
	C := make([][]float64, numARows)
	for i := range C {
		C[i] = make([]float64, numBCols)
	}

	// Step 3 Loop through each row of A
	for i := range numARows {
		// Step 4 Loop through each column of B
		for j := range numBCols {
			sum := 0.0
			// Step 5 Loop through each element of the current row of A and the current column of B
			for k := range numACols {
				sum += A[i][k] * B[k][j]
			}
			// Step 6 assign sum to spot in result matrix C.
			C[i][j] = sum
		}
	}
	return C
}

// DotProduct calculates the dot product of two vectors (slices).
func DotProduct(a, b []float64) []float64 {
	// get lenght of vectors so it only needs to be calc'd once
	lengthA := len(a)
	lengthB := len(b)

	// init result component as a []float64
	product := []float64{0.0}

	// Check if the result is invalid

	if lengthA != 1 && lengthB != 1 && lengthA != lengthB {
		log.Println("invalid inputs")
		return []float64{0.0}
	}

	// Check if it is vector * vector or scalar * vector
	if lengthA == 1 && lengthB != 1 {

		for i := range lengthB {
			hold := a[0] * b[i]
			product = append(product, hold)
		}
	} else if lengthA != 1 && lengthB == 1 {
		for i := range lengthA {
			hold := a[i] * b[0]
			product = append(product, hold)

		}
	} else {

		// for case when lenthA and lengthB are equal
		for i := range lengthA {
			product[0] += a[i] * b[i]
		}
	}
	return product
}
