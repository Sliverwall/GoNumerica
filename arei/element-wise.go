package arei

import (
	"errors"
)

// ElementWise applies a given function element-wise to two Areis.
func ElementWise(a, b *Arei, operation func(float64, float64) float64) (*Arei, error) {
	if !a.SameShape(b) {
		return nil, errors.New("areis must have the same Shape for element-wise operation")
	}

	resultData := make([]float64, len(a.Data))
	for i := range a.Data {
		resultData[i] = operation(a.Data[i], b.Data[i])
	}

	return &Arei{
		Shape: a.Shape,
		Data:  resultData,
	}, nil
}

// Sum performs element-wise addition of two Areis.
func Sum(a, b *Arei) (*Arei, error) {
	if !a.SameShape(b) {
		return nil, errors.New("areis must have the same Shape for element-wise operation")
	}

	resultData := make([]float64, len(a.Data))
	for i := range a.Data {
		resultData[i] = a.Data[i] + b.Data[i]
	}

	return &Arei{
		Shape: a.Shape,
		Data:  resultData,
	}, nil
}

// Sub performs element-wise subtraction of two Areis.
func Sub(a, b *Arei) (*Arei, error) {
	if !a.SameShape(b) {
		return nil, errors.New("areis must have the same Shape for element-wise operation")
	}

	resultData := make([]float64, len(a.Data))
	for i := range a.Data {
		resultData[i] = a.Data[i] - b.Data[i]
	}

	return &Arei{
		Shape: a.Shape,
		Data:  resultData,
	}, nil
}

// Multi performs element-wise multiplication of two Areis.
func Multi(a, b *Arei) (*Arei, error) {
	if !a.SameShape(b) {
		return nil, errors.New("areis must have the same Shape for element-wise operation")
	}

	resultData := make([]float64, len(a.Data))
	for i := range a.Data {
		resultData[i] = a.Data[i] * b.Data[i]
	}

	return &Arei{
		Shape: a.Shape,
		Data:  resultData,
	}, nil
}

// Div performs element-wise Division of two Areis.
func Div(a, b *Arei) (*Arei, error) {
	if !a.SameShape(b) {
		return nil, errors.New("areis must have the same Shape for element-wise operation")
	}

	resultData := make([]float64, len(a.Data))
	for i := range a.Data {
		if b.Data[i] == 0 {
			return nil, errors.New("cannot divide by 0")
		}
		resultData[i] = a.Data[i] / b.Data[i]
	}

	return &Arei{
		Shape: a.Shape,
		Data:  resultData,
	}, nil
}
