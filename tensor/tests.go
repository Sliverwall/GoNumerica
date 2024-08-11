package tensor

import (
	"fmt"
	"log"
)

func Test_1() {
	// Create a vector (1D Tensor)
	x, err := NewTensor([]float64{1, 2, 3})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(x)

	// Create another vector
	y, _ := NewTensor([]float64{2, 3, 4})

	log.Println("Test_ 1 of dot product of two equal length vectors")
	log.Println("x: ", x, " y: ", y)
	// Get dot product, vector * vector will be a scalar
	result, _ := DotProduct(x, y)
	log.Println(result)

}

func Test_2() {
	// Create a vector
	x, _ := NewTensor([]float64{5, 6, 7})

	// scalar element
	y, _ := NewTensor([]float64{5})

	log.Println("Test_ 2 of dot product of vector and scalar")
	log.Println("x: ", x, " y: ", y)

	// Get dot product. scalar * vector will be a vector
	result, _ := DotProduct(x, y)
	log.Println(result)
}

func Test_3() {
	// Create matrix
	A, _ := NewTensor([][]float64{
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
	log.Println("Test_4 tests creation of identity matrix and zeros tensor")

	// None Sqaure matrix
	A, _ := NewTensor([][]float64{
		{1, 2},
		{2, 2},
		{3, 3},
	})

	log.Println("Matrix A non-sqaure:", A)

	// Sqaure matrix
	B, _ := NewTensor([][]float64{
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

	// Test Zeros tensor
	zeroA, _ := Zeros([]int{2, 2})

	log.Println("2x2 Zeros matrix:", zeroA)

	// Test vector zeros tensor

	zeroX, _ := Zeros([]int{1, 5})

	log.Println("1x5 Zeros vector", zeroX)

}
