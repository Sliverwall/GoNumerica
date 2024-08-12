package arei

import (
	"errors"
)

// ElementWise applies a given function element-wise to two Areis.
func ElementWise(a, b *Arei, operation func(float64, float64) float64) (*Arei, error) {
	if !a.SameShape(b) {
		return nil, errors.New("areis must have the same shape for element-wise operation")
	}

	resultData := make([]float64, len(a.data))
	for i := range a.data {
		resultData[i] = operation(a.data[i], b.data[i])
	}

	return &Arei{
		shape: a.shape,
		data:  resultData,
	}, nil
}

// Sum performs element-wise addition of two Areis.
func Sum(a, b *Arei) (*Arei, error) {
	if !a.SameShape(b) {
		return nil, errors.New("areis must have the same shape for element-wise operation")
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
		return nil, errors.New("areis must have the same shape for element-wise operation")
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

// Multi performs element-wise multiplication of two Areis.
func Multi(a, b *Arei) (*Arei, error) {
	if !a.SameShape(b) {
		return nil, errors.New("areis must have the same shape for element-wise operation")
	}

	resultData := make([]float64, len(a.data))
	for i := range a.data {
		resultData[i] = a.data[i] * b.data[i]
	}

	return &Arei{
		shape: a.shape,
		data:  resultData,
	}, nil
}

// Div performs element-wise Division of two Areis.
func Div(a, b *Arei) (*Arei, error) {
	if !a.SameShape(b) {
		return nil, errors.New("areis must have the same shape for element-wise operation")
	}

	resultData := make([]float64, len(a.data))
	for i := range a.data {
		if b.data[i] == 0 {
			return nil, errors.New("cannot divide by 0")
		}
		resultData[i] = a.data[i] / b.data[i]
	}

	return &Arei{
		shape: a.shape,
		data:  resultData,
	}, nil
}
