package stats

import "GoNumerica/interfaces"

// module for basic math functions

func mean[T interfaces.NumArray](X T) float64 {
	// init sum
	sum := 0.0

	// Get length of dataset
	n := len(X)

	// Get the sum of the dataset
	for i := 0; i < n; i++ {
		sum += interfaces.Index(X, i).(float64)
	}

	// Return sum over length
	return (sum / float64(n))
}
