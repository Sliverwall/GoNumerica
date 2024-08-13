package arei

import (
	"fmt"
	"log"
)

func Test_1() {
	// Create a vector (1D Arei)
	x, err := NewArei([]float64{1, 2, 3})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(x)

	// Create another vector
	y, _ := NewArei([]float64{2, 3, 4})

	log.Println("Test_ 1 of dot product of two equal length vectors")
	log.Println("x: ", x, " y: ", y)
	// Get dot product, vector * vector will be a scalar
	result, _ := DotProduct(x, y)
	log.Println(result)

}

func Test_2() {
	// Create a vector
	x, _ := NewArei([]float64{5, 6, 7})

	// scalar element
	y, _ := NewArei([]float64{5})

	log.Println("Test_ 2 of dot product of vector and scalar")
	log.Println("x: ", x, " y: ", y)

	// Get dot product. scalar * vector will be a vector
	result, _ := DotProduct(x, y)
	log.Println(result)
}

func Test_3() {
	// Create matrix
	A, _ := NewArei([][]float64{
		{1, 2, 3},
		{2, 3, 3},
		{9, 2, 0},
	})
	// Test_ 3 statement
	log.Println("Test_ 3 uses Identity function to create an I of A then NatrixProduct to find their product.")
	// Confirm matrix A
	log.Println("Matrix A:", A)
	// Use Identity function to create matrix I of A
	I, _ := Identity(A)

	// Confirm identity matrix
	log.Println("Matrix I:", I)

	// Get matrix product, which should be just A
	result, _ := MatrixProduct(A, I)

	log.Println("Result matrix (A * I):", result)
}

func Test_4() {

	// Test_ statement
	log.Println("Test_4 tests creation of identity matrix and zeros Arei")

	// None Sqaure matrix
	A, _ := NewArei([][]float64{
		{1, 2},
		{2, 2},
		{3, 3},
	})

	log.Println("Matrix A non-sqaure:", A)

	// Sqaure matrix
	B, _ := NewArei([][]float64{
		{2, 3, 1},
		{2, 2, 2},
		{2, 1, 1},
	})
	log.Println("Matrix B sqaure", B)

	// Test Identity for A

	aI, err := Identity(A)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("I of A:", aI)
	}

	// Test Identity for B
	bI, err := Identity(B)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("I of B:", bI)
	}

	// Test Zeros Arei
	zeroA, _ := Zeros([]int{2, 2})

	log.Println("2x2 Zeros matrix:", zeroA)

	// Test vector zeros Arei

	zeroX, _ := Zeros([]int{1, 5})

	log.Println("1x5 Zeros vector", zeroX)

}

func Test_5() {
	// Test Arei methods
	log.Println("Test_5 for Arei methods")

	// Check same Shape
	A, _ := Zeros([]int{3, 3})
	B, _ := Zeros([]int{3, 3})
	C, _ := Zeros([]int{2, 5})

	log.Println("A:", A, "\nB:", B, "\nC:", C)
	log.Println("Testing A.SameShape(*Arei)")
	log.Println("A same Shape as B:", A.SameShape(B))
	log.Println("A same Shape as C", A.SameShape(C))

	log.Println("Testing A.Reshape([]int{1,9})")

	newDim := []int{1, 9}
	// Try to catch error. Will fail if all elements can be redistruted to new form.
	err := A.Reshape(newDim)
	if err != nil {
		log.Println(err)
	} else {
		A.Reshape(newDim)
	}
	log.Println("New A:", A)
}

func Test_6() {
	// Test Arei transpose

	// Matrix transpose
	A, _ := Zeros([]int{2, 5})

	log.Println("A", A)

	A.Transpose()
	log.Println("A tranposed", A)

	// Vector transpose
	X, _ := Zeros([]int{1, 5})

	log.Println("X", X)
	X.Transpose()
	log.Println("X tranposed", X)

}

func Test_7() {
	// Test few element-wise functions
	log.Println("Test_7 test a few element-wise functions")
	A, _ := NewArei([][]float64{
		{1, 2, 3},
		{2, 2, 2},
		{0, 2, -1},
	})

	log.Println("Matrix A:", A)

	// Apply Exp function
	A_E, _ := Exp(A)

	log.Println("Matrix A with Exp applied", A_E)

	// Test Max on A. Result should be 3
	maxA, _ := A.Max()

	log.Println("Max element in Matrix A:", maxA)

	// Test Min on A. Result should b -1
	minA, _ := A.Min()

	log.Println("Min element in Matrix A:", minA)
}

func Test_8() {
	// Test flatten and transpose and count methods
	log.Println("Test_8 test flatten, transpose, and count methods")

	// Create Matrix with Ns
	N, _ := Ns([]int{3, 6}, 10)
	log.Println(N)
	// Create matrix
	A, _ := Zeros([]int{3, 6})
	log.Println(A)
	// Test Flatten on A
	A.Flatten()
	log.Println(A)

	// Test Flatten on vector

	X, _ := NewArei([]float64{1, 2, 3, 4})
	// confirm error check prevents crash
	X.Flatten()
	log.Println(X)

	// Test transpose on vector
	X.Transpose()
	log.Println(X)
	log.Println(X.Shape, X.Shape[0], X.Shape[1])

	// Test count on matrix and vector
	log.Println("A count:", A.Count(), " X count:", X.Count())
}

func Test_9() {
	// Test element-wise functions

	log.Println("Test_9 to test Multi, Div, Elemetwise, and Sign")

	A, _ := NewArei([][]float64{
		{1, 2, 3, 4},
		{4, 3, 2, 3},
	})

	B, _ := NewArei([][]float64{
		{1, 2, 3, 4},
		{4, 3, 2, 3},
	})

	log.Println("A:", A, "B:", B)

	// Test multi and div
	AB_multi, _ := Multi(A, B)
	AB_div, _ := Div(A, B)

	log.Println("AB_multi:", AB_multi)
	log.Println("AB_div:", AB_div)

	A_minus, _ := Sign(A)
	log.Println("-A", A_minus)

	// Assign function to variable to use in elementwise
	divide := func(a, b float64) float64 {
		return a / b
	}
	AB_elementwise, _ := ElementWise(A, B, divide)

	log.Println("AB_elementwise(divide):", AB_elementwise)
}

func Test_10() {
	// Test Search functions

	// Test matrix
	A, _ := NewArei([][]float64{
		{3, 3, 3},
		{1, 3, 2},
		{3, 2, 1},
	})

	log.Println("matrix A:", A)
	maxIndexes, _ := WhereMax(A)

	// Should be {2,1,0}
	log.Println("Max indices of A along each row:", maxIndexes)

	// Now a vector
	X, _ := NewArei([]float64{1, 2, 3})

	log.Println("vector X:", X)

	maxIndexesVector, _ := WhereMax(X)
	log.Println("Max index of X", maxIndexesVector)

	// Find min indexes along each row
	minIndexes, _ := WhereMin(A)

	log.Println("Min indices of A along each row:", minIndexes)

	// Find min index of vector
	minIndexVector, _ := WhereMin(X)
	log.Println("Min index of X", minIndexVector)

	// conditional where true if element >= 3
	cond := func(element float64) bool {
		if element >= 3.0 {
			return true
		} else {
			return false
		}
	}
	threeIndexes, _ := Where(A, cond)
	log.Println(threeIndexes)

	// Test Index method
	value, _ := A.Index(0, 2)
	log.Println("A.Index(0,2) is (row=0,col=2) == A[0][2]:", value)
	// Test set index
	A.SetIndex(10.0, 0, 2)
	value, _ = A.Index(0, 2)
	log.Println("A.Index(0,2) after A.SetIndex(10.0,0,2):", value)

	// Test vector indexing
	valueX, _ := X.Index(2)

	log.Println("X.Index(2) for just column in 1D arei:", valueX)

	// Test vector set index

	X.SetIndex(10.0, 2)
	valueX, _ = X.Index(2)

	log.Println("X.Index(2) after X.SetIndex(10.0,2):", valueX)

}

func Test_11() {
	// Test frame print format

	log.Println("Test Arei's Frame() method for printing Areis")
	A, _ := Ns([]int{2, 4}, 4.0)
	log.Println("Print A regular:", A)

	log.Println("Print with A.Frame()")
	A.Frame()

	X, _ := Ns([]int{1, 4}, 4.0)
	log.Println("Print X regular:", X)

	log.Println("Print with X.Frame()")
	X.Frame()
}

func Test_12() {
	// Test Row and Column search functions

	A, _ := NewArei([][]float64{
		{1, 2, 4},
		{1, 2, 3},
		{3, 2, 1},
	})

	log.Println("A:")
	A.Frame()
	row0, _ := Row(A, 0)
	log.Println("Row0:", row0)
	col2, _ := Column(A, 2)
	log.Println("Col2:", col2)

	rowNeg1, _ := Row(A, -1)
	colNeg2, _ := Column(A, -2)
	log.Println("Row-1", rowNeg1)
	log.Println("col-2", colNeg2)

	// Test negative indexing
	indexNeg1Neg1, _ := A.Index(-2, -1)
	log.Println("A.Index(-2,-1)", indexNeg1Neg1)

	log.Println("SetIndex(10,-2,-1)")
	A.SetIndex(10.0, -2, -1)
	indexNeg1Neg1, _ = A.Index(-2, -1)
	log.Println("A.Index(-2,-1)", indexNeg1Neg1)

}
