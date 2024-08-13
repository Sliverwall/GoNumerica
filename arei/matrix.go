package arei

import (
	"errors"
)

// MatrixProduct performs matrix multiplication of two Arei representing matrices.
func MatrixProduct(A, B *Arei) (*Arei, error) {
	// Check that A and B are 2D tensors
	if len(A.Shape) != 2 || len(B.Shape) != 2 {
		return nil, errors.New("both areis must be 2D for matrix multiplication")
	}

	numARows, numACols := A.Shape[0], A.Shape[1]
	numBRows, numBCols := B.Shape[0], B.Shape[1]

	// Check if matrix dimensions are compatible for multiplication
	if numACols != numBRows {
		return nil, errors.New("number of columns in A must be equal to the number of rows in B")
	}

	// Initialize result matrix C with zeros
	C := &Arei{
		Shape: []int{numARows, numBCols},
		Data:  make([]float64, numARows*numBCols),
	}

	// Perform matrix multiplication
	for i := 0; i < numARows; i++ {
		for j := 0; j < numBCols; j++ {
			sum := 0.0
			for k := 0; k < numACols; k++ {
				sum += A.Data[i*numACols+k] * B.Data[k*numBCols+j]
			}
			C.Data[i*numBCols+j] = sum
		}
	}

	return C, nil
}

// Assume it will work always for now
// Elimination takes a given aeri and outputs the upper trianglar matrix
func Elimination(a *Arei) (*Arei, *Arei, error) {
	// Assume a matrix
	if len(a.Shape) == 1 {
		return nil, nil, errors.New("arei cannot be a 1d arei")
	}

	// Create copy of input arei
	u, _ := a.Copy()
	// Number of pivots equal to number of cols - 1. Do not need to pivot last row
	nPivots := u.Shape[0] - 1

	// Create factor array to hold all needed factors
	factorArrayCounter := 0
	factorArray := make([]int, u.Shape[0])

	for i := range nPivots {
		// pivot will be along the dianole of square matrix
		pivot, _ := u.Index(i, i)

		// num factors will be equal to nPivots - current pivot index
		numFactor := nPivots - i

		// Create factor array to hold factors needed for the current row
		factorList := make([]int, numFactor)
		for factorIndex := range numFactor {
			// Get values below pivot till reaching the bottom of the matrix
			factorValue, _ := u.Index(factorIndex+1, i)

			// factorRx = Axi * Aii
			factor_pivot := int(factorValue) * int(pivot)
			factorList[factorIndex] = factor_pivot

			// Add factor to list of all factors needed
			factorArray[factorArrayCounter] = factor_pivot
			factorArrayCounter += 1
		}

		for rowIndex := range numFactor {
			// Get row to modify
			// Rx = Rx - (factorRx * Ri)
			rx, _ := Row(u, rowIndex+1+i)
			// // Debug statement see selected row before element-wise sub
			// log.Println("Row before sub", rowIndex+1+i, rx)
			ri, _ := Row(u, i)
			factorRx, _ := NewArei([]float64{float64(factorList[rowIndex])})

			factorRx_ri, _ := DotProduct(ri, factorRx)

			rx, _ = Sub(rx, factorRx_ri)
			// // Debug statement see selected row after element-wise sub
			// log.Println("Row after sub", rowIndex+1+i, rx)

			// overwrite row in u with new row's values
			for newRowIndex := range u.Shape[0] {
				setValue := rx.Data[newRowIndex]
				u.SetIndex(setValue, rowIndex+1+i, newRowIndex)
			}
		}
	}
	// Now find the lower trianglar as l
	l, _ := Identity(u)

	// Set indexes for L based on stored factorArray
	row, col := 0, 0
	for i := range factorArray {
		for j := range i {
			row = i
			col = j
			// // Debug statement
			// log.Println("i,j", "(", i, j, ") item:", factorArray[(i+j)-1], " at row,col", row, col)

			l.SetIndex(float64(factorArray[(i+j)-1]), row, col)
		}
	}

	return l, u, nil

}
