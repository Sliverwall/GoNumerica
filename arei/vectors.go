package arei

import "errors"

// DotProduct calculates the dot product or scalar-vector product of two Areis.
func DotProduct(a, b *Arei) (*Arei, error) {
	// Check that both Areis are 1D
	if len(a.shape) != 1 || len(b.shape) != 1 {
		return nil, errors.New("both Areis must be 1D for dot product")
	}

	lengthA, lengthB := a.shape[0], b.shape[0]

	// Initialize result Arei
	result := &Arei{}

	if lengthA == 1 && lengthB != 1 {
		// Scalar-vector multiplication
		result = &Arei{
			shape: []int{lengthB},
			data:  make([]float64, lengthB),
		}
		for i := 0; i < lengthB; i++ {
			result.data[i] = a.data[0] * b.data[i]
		}
	} else if lengthA != 1 && lengthB == 1 {
		// Vector-scalar multiplication
		result = &Arei{
			shape: []int{lengthA},
			data:  make([]float64, lengthA),
		}
		for i := 0; i < lengthA; i++ {
			result.data[i] = a.data[i] * b.data[0]
		}
	} else if lengthA == lengthB {
		// Vector-vector dot product
		sum := 0.0
		for i := 0; i < lengthA; i++ {
			sum += a.data[i] * b.data[i]
		}
		result = &Arei{
			shape: []int{1},
			data:  []float64{sum},
		}
	} else {
		return nil, errors.New("invalid dimensions for dot product")
	}

	return result, nil
}
