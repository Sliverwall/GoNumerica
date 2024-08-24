package stats

import (
	"GoNumerica/interfaces"
	"math"
)

// Stdev takes any 1D array then returns the standard deviation
func Stdev[T interfaces.NumArray](X T) float64 {
	// Standard deviation formula

	// Get the length of dataset
	n := len(X)

	// Get the average value in the dataset
	avg := mean(X)

	// init sum
	sum := 0.0

	// Summation loop
	for i := 0; i < n; i++ {

		xi := interfaces.Index(X, i).(float64)
		// Add to the some the indvidual (x - the average) sqaured
		sum += math.Pow((xi - avg), 2)
	}

	// Final term
	s := math.Sqrt((1 / (float64(n) - 1)) * sum)

	return s
}
