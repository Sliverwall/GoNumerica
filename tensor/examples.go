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

	y, _ := NewTensor([]float64{2, 3, 4})

	log.Println("Example 1 of dot product of two equal length vectors")
	log.Println("x: ", x, " y: ", y)
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

	result, _ := DotProduct(x, y)
	log.Println(result)
}
