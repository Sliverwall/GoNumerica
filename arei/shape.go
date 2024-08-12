package arei

import "errors"

// Functions to create special kind of matrixes

// Identity creates an identity Arei with the same shape as the given Arei
func Identity(a *Arei) (*Arei, error) {
	switch len(a.shape) {
	case 1:
		// Identity vector (scalar 1)
		identityData := make([]float64, a.shape[0])
		for i := range identityData {
			identityData[i] = 1.0
		}
		return &Arei{shape: a.shape, data: identityData}, nil
	case 2:
		if a.shape[0] != a.shape[1] {
			return nil, errors.New("cannot create an identity matrix for a non-square matrix")
		}
		identityData := make([]float64, len(a.data))
		for i := 0; i < a.shape[0]; i++ {
			identityData[i*a.shape[0]+i] = 1.0
		}
		return &Arei{shape: a.shape, data: identityData}, nil
	default:
		return nil, errors.New("invalid shape for identity Arei")
	}
}

// Zeros creates a Arei of zeros based on the provided shape.
func Zeros(shape []int) (*Arei, error) {
	if len(shape) == 0 {
		return nil, errors.New("shape cannot be empty")
	}

	size := 1
	for _, dim := range shape {
		if dim <= 0 {
			return nil, errors.New("shape dimensions must be positive integers")
		}
		size *= dim
	}

	zeroData := make([]float64, size)
	return &Arei{shape: shape, data: zeroData}, nil
}
