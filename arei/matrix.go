package arei

import "errors"

// MatrixProduct performs matrix multiplication of two Tensors representing matrices.
func MatrixProduct(A, B *Arei) (*Arei, error) {
	// Check that A and B are 2D tensors
	if len(A.shape) != 2 || len(B.shape) != 2 {
		return nil, errors.New("both tensors must be 2D for matrix multiplication")
	}

	numARows, numACols := A.shape[0], A.shape[1]
	numBRows, numBCols := B.shape[0], B.shape[1]

	// Check if matrix dimensions are compatible for multiplication
	if numACols != numBRows {
		return nil, errors.New("number of columns in A must be equal to the number of rows in B")
	}

	// Initialize result matrix C with zeros
	C := &Arei{
		shape: []int{numARows, numBCols},
		data:  make([]float64, numARows*numBCols),
	}

	// Perform matrix multiplication
	for i := 0; i < numARows; i++ {
		for j := 0; j < numBCols; j++ {
			sum := 0.0
			for k := 0; k < numACols; k++ {
				sum += A.data[i*numACols+k] * B.data[k*numBCols+j]
			}
			C.data[i*numBCols+j] = sum
		}
	}

	return C, nil
}
