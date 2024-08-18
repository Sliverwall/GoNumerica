package geometry

import (
	"GoNumerica/interfaces"
	"math"
)

// Module for Eculidean geomtric functions

// Hypotenuse calculates the hypotenuse of a right triangle
// given a vector x
func Hypotenuse[T interfaces.Num](x ...T) float64 {
	// Sqaure each scalar to get the sum of the right handside
	// c^2 = i -> n, Sum(x[i]*x[i])
	c_c := 0.0
	for i := 0; i < len(x); i++ {
		c_c += (float64(x[i]) * float64(x[i]))
	}

	c := math.Sqrt(float64(c_c))
	return c
}
