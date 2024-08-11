package tensor

import (
	"fmt"
	"log"
)

func Example1() {
	// Create a vector (1D Tensor)
	x, err := NewTensor([]float64{1, 2, 3})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(x)

	// Create another vector
	y, _ := NewTensor([]float64{2, 3, 4})

	log.Println("Example 1 of dot product of two equal length vectors")
	log.Println("x: ", x, " y: ", y)
	// Get dot product, vector * vector will be a scalar
	result, _ := DotProduct(x, y)
	log.Println(result)

}

func Example2() {
	// Create a vector
	x, _ := NewTensor([]float64{5, 6, 7})

	// scalar element
	y, _ := NewTensor([]float64{5})

	log.Println("Example 2 of dot product of vector and scalar")
	log.Println("x: ", x, " y: ", y)

	// Get dot product. scalar * vector will be a vector
	result, _ := DotProduct(x, y)
	log.Println(result)
}

func Example3() {
	// Create matrix
	A, _ := NewTensor([][]float64{
		{1, 2, 3},
		{2, 3, 3},
		{9, 2, 0},
	})
	// Example 3 statement
	log.Println("Example 3 uses Identity function to create an I of A then NatrixProduct to find their product.")
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
