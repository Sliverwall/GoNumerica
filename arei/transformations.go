package arei

import "math"

// Transformations take an existing Arei and apply a transformation to create a new Arei

// Transform takes an Arei and a given transformation function and applies to each element
func Transform(a *Arei, transformation func(float64) float64) (*Arei, error) {
	resultData := make([]float64, len(a.data))

	for i := range a.data {
		resultData[i] = transformation(a.data[i])
	}

	return &Arei{
		shape: a.shape,
		data:  resultData,
	}, nil
}

// Exp takes each element in an Arei with base euler's number
func Exp(a *Arei) (*Arei, error) {

	resultData := make([]float64, len(a.data))

	for i := range a.data {
		resultData[i] = math.Exp(a.data[i])
	}

	return &Arei{
		shape: a.shape,
		data:  resultData,
	}, nil
}

// Sign creates a new arei with flipped signs on each element
func Sign(a *Arei) (*Arei, error) {

	resultData := make([]float64, len(a.data))

	for i := range a.data {
		resultData[i] = a.data[i] * -1
	}

	return &Arei{
		shape: a.shape,
		data:  resultData,
	}, nil
}
