package stats

import (
	"GoNumerica/arei"
	"log"
)

func Test_1() {
	// test Stdev function
	X, _ := arei.NewArei([][]float64{
		//   x y z
		{1, 2, 3},
		{5, 4, 3},
		{3, 2, 3},
		{1, 2, 3},
	})

	// print mean and std of each feature

	for i := 0; i < X.Shape[1]; i++ {
		// Get mean
		mean := Mean(X, i)
		// Get stdev
		std := Stdev(X, i)
		// Get variance
		variance := Var(X, i)
		// Get RSD
		cvPercent := Rsd(X, i) * 100

		log.Println("Mean of feature", i, "=", mean, "and stdev =", std, "and var =", variance, "RSD% =", cvPercent)
	}

	// Correlation between x and y
	x_y := Corr(X, 1, 0)
	x_z := Corr(X, 0, 2)
	log.Println("Correlation between feature 0 and 1:", x_y, "And 0 and 2:", x_z)
}
