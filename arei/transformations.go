package arei

import (
	"errors"
	"math"
)

// Transformations take an existing Arei and apply a transformation to create a new Arei

// Transform takes an Arei and a given transformation function and applies to each element
func Transform(a *Arei, transformation func(float64) float64) *Arei {
	resultData := make([]float64, len(a.Data))

	for i := range a.Data {
		resultData[i] = transformation(a.Data[i])
	}

	return &Arei{
		Shape: a.Shape,
		Data:  resultData,
	}
}

// Exp takes each element in an Arei with base euler's number
func Exp(a *Arei) *Arei {

	resultData := make([]float64, len(a.Data))

	for i := range a.Data {
		resultData[i] = math.Exp(a.Data[i])
	}

	return &Arei{
		Shape: a.Shape,
		Data:  resultData,
	}
}

// Sign creates a new arei with flipped signs on each element
func Sign(a *Arei) *Arei {

	resultData := make([]float64, len(a.Data))

	for i := range a.Data {
		resultData[i] = a.Data[i] * -1
	}

	return &Arei{
		Shape: a.Shape,
		Data:  resultData,
	}
}

// MultiT takes a given x then multiplies each element of the matrix by x.
func MultiT(a *Arei, x float64) *Arei {

	resultData := make([]float64, len(a.Data))

	for i := range a.Data {
		resultData[i] = a.Data[i] * x
	}

	return &Arei{
		Shape: a.Shape,
		Data:  resultData,
	}
}

// DivT takes a given x then multiplies each element of the matrix by x.
func DivT(a *Arei, x float64) (*Arei, error) {

	// Handle x = 0
	if x == 0 {
		return nil, errors.New("x cannot be 0")
	}
	resultData := make([]float64, len(a.Data))

	for i := range a.Data {
		resultData[i] = a.Data[i] / x
	}

	return &Arei{
		Shape: a.Shape,
		Data:  resultData,
	}, nil
}
