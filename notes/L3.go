package notes

import (
	"GoNumerica/arei"
	"log"
)

// Video 5

// A = LU now becomes PA = LU. Where P is a permutation matrix that does the necessary row exchanges
// in cases where we get 0s at the pivot. In the case where this is not necessary, P = I.

func Example8() {
	// LU=PA Matrix Decomposition/Factoring video

	case1, _ := arei.NewArei([][]float64{
		{-1, 0, 3},
		{2, 1, 3},
		{1, 1, 2},
	})
	// print row 1 of case1

	log.Println(case1)
	log.Println(case1.Shape)

	// Get L and U from A assuming P == I

	L, U, _ := arei.Elimination(case1)

	log.Println("U:", U, " L:", L)

	switchRows, _ := arei.NewArei([][]float64{
		{0, 0, 1},
		{0, 1, 0},
		{1, 0, 0},
	})

	case2, _ := arei.MatrixProduct(switchRows, case1)

	log.Println("Case 2:")
	case2.Frame()
	L2, U2, _ := arei.Elimination(case2)

	log.Println("L of Case 2:")
	L2.Frame()
	log.Println("U of Case 2:")
	U2.Frame()

	testCase2, _ := arei.MatrixProduct(L2, U2)

	log.Println("Test case2 result:")
	testCase2.Frame()

	// Test video's solution

	L3, _ := arei.NewArei([][]float64{
		{1, 0, 0},
		{2, 1, 0},
		{-1, -1, 1},
	})

	U3, _ := arei.NewArei([][]float64{
		{1, 1, 2},
		{0, -1, -1},
		{0, 0, 4},
	})

	testCaseVideo, _ := arei.MatrixProduct(L3, U3)
	log.Println("Video test case:")
	testCaseVideo.Frame()
}

func Example9() {
	// Working with non-sqaure matrices

	A, _ := arei.NewArei([][]float64{
		{-15, 30, -10, -5},
		{25, 28, 34, 35},
		{5, -7, 4, 3},
	})

	P, _ := arei.NewArei([][]float64{
		{0, 0, 1},
		{1, 0, 0},
		{0, 1, 0},
	})

	case1, _ := arei.MatrixProduct(P, A)

	log.Println("case 1:")
	case1.Frame()

	L, U, err := arei.Elimination(case1)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("L:")
	L.Frame()

	log.Println("U:")
	U.Frame()
}

func Example10() {
	// Why Matrices are import video

	// 150 students and zombie in a school. Every hour 20% of humans turn into zombies and 10% of zombies cured
	// humans : 0.8(h) + 0.1(z)
	// zombies : 0.2(h) + 0.9(z)

	// Markov Matrix
	// 2x2
	A, err := arei.NewArei([][]float64{
		{0.8, 0.1},
		{0.2, 0.9},
	})
	if err != nil {
		log.Fatal(err)
	}

	setTime := 40

	// multiply A by itself x times to get the markov matrix needed for any given timepoint
	for i := 0; i < setTime; i++ {
		A, err = arei.MatrixProduct(A, A)
		if err != nil {
			log.Fatal(err)
		}

	}
	// vector of humans and zombies at t = 0
	t0, _ := arei.NewArei([][]float64{
		{150},
		{150},
	})

	at_hour, _ := arei.MatrixProduct(A, t0)
	log.Println("Students/Zombies at T =", setTime)
	at_hour.Frame()
}
