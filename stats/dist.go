package stats

import (
	"GoNumerica/arei"
	"math"
)

// Stdev takes an Arei and a feature index then return the sample standard devation of the selected feature
func Stdev(X *arei.Arei, featueIndex int) float64 {
	// Get the length of feature
	n := X.Shape[0]

	// Get the average value in the feature
	avg := Mean(X, featueIndex)

	// init sum
	sum := 0.0

	// Summation loop
	for i := 0; i < n; i++ {
		// Index current entry on feature column
		xi, _ := X.Index(i, featueIndex)
		// Add to the some the indvidual (x - the average) sqaured
		sum += math.Pow((xi - avg), 2)
	}

	// Final term
	s := math.Sqrt((1 / (float64(n) - 1)) * sum)

	return s
}

// Var takes an Arei and a feature index then return the sample variance of the selected feature
func Var(X *arei.Arei, featueIndex int) float64 {
	// Get the length of feature
	n := X.Shape[0]

	// Get the average value in the feature
	avg := Mean(X, featueIndex)

	// init sum
	sum := 0.0

	// Summation loop
	for i := 0; i < n; i++ {
		// Index current entry on feature column
		xi, _ := X.Index(i, featueIndex)
		// Add to the some the indvidual (x - the average) sqaured
		sum += math.Pow((xi - avg), 2)
	}

	// Final term
	s := (1 / (float64(n) - 1)) * sum

	return s
}

// Corr takes an Arei and two feature indices then return the covarance of two selected features
func CoVar(X *arei.Arei, xIndex, yIndex int) float64 {
	// Get the length of feature
	n := X.Shape[0]

	// Get the average values in the features
	xAvg := Mean(X, xIndex)
	yAvg := Mean(X, yIndex)

	// init sum
	sum := 0.0

	// Summation loop
	for i := 0; i < n; i++ {
		// Index current entry on feature columns
		xi, _ := X.Index(i, xIndex)
		yi, _ := X.Index(i, yIndex)
		// Add to the summ the difference between each entriy's value and their feature's average
		sum += (xi - xAvg) * (yi - yAvg)
	}

	// Final term
	s := (1 / (float64(n) - 1)) * sum

	return s
}

// Corr takes an Arei and two feature indices then return the Pearsons correlation of two selected features
func Corr(X *arei.Arei, xIndex, yIndex int) float64 {
	// Pearsons
	sigmaX := Stdev(X, xIndex)
	sigamY := Stdev(X, yIndex)

	// If there is no variance in a feature, there will be no correlation
	if sigmaX == 0 || sigamY == 0 {
		return 0
	}
	r := CoVar(X, xIndex, yIndex) / (sigmaX * sigamY)

	return r
}

// Rsd takes an Arei and a feature index, then returns the relative standard deviation of the selected feature
func Rsd(X *arei.Arei, featureIndex int) float64 {
	sigma := Stdev(X, featureIndex)
	mean := Mean(X, featureIndex)

	return sigma / mean

}
