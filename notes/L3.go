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

	// Get A from MatrixProduct of LU
	A, _ := arei.MatrixProduct(L, U)

	log.Println("A from LU:", A)

}
