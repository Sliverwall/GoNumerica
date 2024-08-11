package tensor

import "errors"

// Sum performs element-wise addition of two Tensors.
func Sum(a, b *Tensor) (*Tensor, error) {
	if !a.sameShape(b) {
		return nil, errors.New("tensors must have the same shape to add")
	}

	resultData := make([]float64, len(a.data))
	for i := range a.data {
		resultData[i] = a.data[i] + b.data[i]
	}

	return &Tensor{
		shape: a.shape,
		data:  resultData,
	}, nil
}

// Sub performs element-wise subtraction of two Tensors.
func Sub(a, b *Tensor) (*Tensor, error) {
	if !a.sameShape(b) {
		return nil, errors.New("tensors must have the same shape to add")
	}

	resultData := make([]float64, len(a.data))
	for i := range a.data {
		resultData[i] = a.data[i] - b.data[i]
	}

	return &Tensor{
		shape: a.shape,
		data:  resultData,
	}, nil
}
