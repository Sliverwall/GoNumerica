package stats

import (
	"GoNumerica/arei"
)

// module for basic math functions

// Mean takes an Arei and a feature index then return the mean of the selected feature
func Mean(X *arei.Arei, featureIndex int) float64 {
	// init sum
	sum := 0.0

	// Get length of the feature
	n := X.Shape[0]

	// Get the sum of the feature
	for i := 0; i < n; i++ {
		value, _ := X.Index(i, featureIndex)
		sum += value
	}

	// Return sum over length
	return (sum / float64(n))
}
