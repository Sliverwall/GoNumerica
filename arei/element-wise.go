package arei

import (
	"errors"
	"math"
)

// Sum performs element-wise addition of two Areis.
func Sum(a, b *Arei) (*Arei, error) {
	if !a.SameShape(b) {
		return nil, errors.New("areis must have the same shape to add")
	}

	resultData := make([]float64, len(a.data))
	for i := range a.data {
		resultData[i] = a.data[i] + b.data[i]
	}

	return &Arei{
		shape: a.shape,
		data:  resultData,
	}, nil
}

// Sub performs element-wise subtraction of two Areis.
func Sub(a, b *Arei) (*Arei, error) {
	if !a.SameShape(b) {
		return nil, errors.New("areis must have the same shape to add")
	}

	resultData := make([]float64, len(a.data))
	for i := range a.data {
		resultData[i] = a.data[i] - b.data[i]
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

// Max finds the max element in an Arei
func Max(a *Arei) (float64, error) {

	var result float64 = math.Inf(-1)
	for i := range a.data {
		if a.data[i] > result {
			result = a.data[i]
		}
	}
	return result, nil
}

// Min finds the min element in an Arei
func Min(a *Arei) (float64, error) {

	var result float64 = math.Inf(1)
	for i := range a.data {
		if a.data[i] < result {
			result = a.data[i]
		}
	}
	return result, nil
}
