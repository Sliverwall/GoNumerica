package arei

import "errors"

// Sum performs element-wise addition of two Areis.
func Sum(a, b *Arei) (*Arei, error) {
	if !a.sameShape(b) {
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
	if !a.sameShape(b) {
		return nil, errors.New("Areis must have the same shape to add")
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
