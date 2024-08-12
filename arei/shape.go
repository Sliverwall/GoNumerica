package arei

import "errors"

// Functions to create special kind of matrixes

// Ns creates a Arei of n value based on the provided Shape. Defaults to 0 if no value given
func Ns(Shape []int, n ...float64) (*Arei, error) {

	// Default n value to 0
	nValue := 0.0

	// If n value provided, set it to nValue
	if len(n) > 0 {
		nValue = n[0]
	}

	if len(Shape) == 0 {
		return nil, errors.New("shape cannot be empty")
	}

	size := 1
	for _, dim := range Shape {
		if dim <= 0 {
			return nil, errors.New("Shape dimensions must be positive integers")
		}
		size *= dim
	}

	// Create zeros Data place holder
	Data := make([]float64, size)

	if nValue != 0.0 {
		for i := range Data {
			Data[i] = nValue
		}
	}

	return &Arei{Shape: Shape, Data: Data}, nil
}

// Identity creates an identity Arei with the same Shape as the given Arei
func Identity(a *Arei) (*Arei, error) {
	switch len(a.Shape) {
	case 1:
		// Identity vector (scalar 1)
		identityData := make([]float64, a.Shape[0])
		for i := range identityData {
			identityData[i] = 1.0
		}
		return &Arei{Shape: a.Shape, Data: identityData}, nil
	case 2:
		if a.Shape[0] != a.Shape[1] {
			return nil, errors.New("cannot create an identity matrix for a non-square matrix")
		}
		identityData := make([]float64, len(a.Data))
		for i := 0; i < a.Shape[0]; i++ {
			identityData[i*a.Shape[0]+i] = 1.0
		}
		return &Arei{Shape: a.Shape, Data: identityData}, nil
	default:
		return nil, errors.New("invalid Shape for identity Arei")
	}
}

// Zeros creates a Arei of zeros based on the provided Shape.
func Zeros(Shape []int) (*Arei, error) {
	if len(Shape) == 0 {
		return nil, errors.New("Shape cannot be empty")
	}

	size := 1
	for _, dim := range Shape {
		if dim <= 0 {
			return nil, errors.New("Shape dimensions must be positive integers")
		}
		size *= dim
	}

	zeroData := make([]float64, size)
	return &Arei{Shape: Shape, Data: zeroData}, nil
}
