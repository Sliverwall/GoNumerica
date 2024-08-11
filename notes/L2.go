package notes

import "log"

// Method used in this lecutre: Elimination

func Example2() {
	// Matrix with x,y,z. three equations
	A := [][]float64{
		{1, 2, 1},
		{3, 8, 1},
		{0, 4, 1},
	}
	log.Println(A)
	// Elimnation involves finding the pivot, must be non-zero
	// Then find the multiplication value that will have aRi + Rj == 0
	// Here 1 is the first pivot, 1 * -3(a here) gives -3(a) + 3 == 0
	// Then we transform Rj into 0,2,-2

	// Ax represents x[0] * A[][0] + x[1] * A[][1]....

	// {1,0,0},
	// {0,1,0},
	// {0,0,1}, is the identitiy matrix. This times A is A.

	// Permutation exchange can be expressed as by row
	// {0,0,1},
	// {0,1,0},
	// {1,0,0}, Now row 0 is row 2 and row 2 is row 0.

	// Permutation exchange can also be expressed by column.
	// [a b]		[b a]		{0,1}
	// [c d] can be [d c] using {1,0} on the right side

}

func Example3() {
	// Practice from organic chem tutor on how to multiply matrixes

	// 1 row, 3 cols is 1x3
	A := [][]float64{
		{3, 1, 4},
	}

	// 3 rows, 2 cols.
	B := [][]float64{
		{4, 3},
		{2, 5},
		{6, 8},
	}

	// To multiple two matrixes, A must have equal coumns to B rows

	log.Println("Row Length of A: ", len(A), " Row Length of B: ", len(B))
	log.Println("Col Length of A: ", len(A[0]), " Col Length of B: ", len(B[0]))
	log.Println("Row 0 for A: ", A[0], " Row 0 for B: ", B[0])
	log.Println("First element in row 0 for A: ", A[0][0], " First element in row 0 for B: ", B[0][0])

	// Try to get A*B

	// Step 1, check if shape works
	numACols := len(A[0])
	numARows := len(A)

	numBRows := len(B)
	numBCols := len(B[0])

	if numACols == numBRows {

		log.Println("Shape works, A cols = B rows. A cols = ", numACols, " B rows = ", numBRows)
		log.Println("A rows = ", numARows, " B cols = ", numBCols)
	} else {
		log.Println("Shape does NOT Work, A cols != B rows. A cols = ", numACols, " B rows = ", numBRows)
		log.Println("A rows = ", numARows, " B cols = ", numBCols)
	}
	// Initialize result matrix C with zeros
	C := make([][]float64, numARows)
	for i := range C {
		C[i] = make([]float64, numBCols)
	}

	// Loop through each row of A
	for i := 0; i < numARows; i++ {
		// Loop through each column of B
		for j := 0; j < numBCols; j++ {
			sum := 0.0
			// Loop through each element of the current row of A and the current column of B
			for k := 0; k < numACols; k++ {
				sum += A[i][k] * B[k][j]
			}
			C[i][j] = sum
		}
	}

	log.Println(C)
}

func Example4() {

	// Take first matrix pair from Lecture 1 to test DotProduct func

	// // 2x2
	// A := [][]float64{
	// 	//Col
	// 	{2, 5}, //Row
	// 	{1, 3},
	// }
	// // x := []int{1, 2} turn this into a matrix now

	// // 1x2
	// x := [][]float64{
	// 	{1, 2},
	// }

	// log.Println(DotProduct(x, A))

	// 3x2
	a := [][]float64{
		{4, 1},
		{6, 2},
		{8, 3},
	}

	// 2 x 3
	b := [][]float64{
		{5, 3, 2},
		{1, 1, 1},
	}

	log.Println(MatrixProduct(a, b))
}

func Example5() {

	// Multiplying matrix by I
	// 3x3
	A := [][]float64{
		{1, 2, 3},
		{2, 3, 3},
		{9, 2, 0},
	}

	// 3x3
	I := [][]float64{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}

	result := MatrixProduct(A, I)

	log.Println(result)

}

func Example6() {
	// Dot product example between two vectors
	x := []float64{1.0, 2.0, 3.0}

	y := []float64{2.0, 3.0, 3.0}

	result := DotProduct(x, y)

	log.Println(result)
}

func Example7() {

	// Dot product
	a := []float64{2.0, 3.0}

	b := []float64{5.0, -4.0}

	// With two vectors
	ab := DotProduct(a, b)

	// With one scalar and one vector
	ab_a := DotProduct(ab, a)

	log.Println(ab)
	log.Println(ab_a)
}
