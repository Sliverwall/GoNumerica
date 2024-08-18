package geometry

import (
	"GoNumerica/interfaces"
	"math"
)

// Module for Eculidean geomtric functions

// Hypotenuse calculates the hypotenuse of a right triangle
// given sides a and b, for int or floats
func Hypotenuse[T interfaces.Num](a, b T) float64 {
	// Convert the input values to float64 for calculations
	c_c := a*a + b*b
	c := math.Sqrt(float64(c_c))
	return c
}
