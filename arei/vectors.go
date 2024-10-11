package arei

import "errors"

// DotProduct calculates the dot product or scalar-vector product of two Areis.
func DotProduct(a, b *Arei) (*Arei, error) {
	// Check that both Areis are 1D
	if a.Shape[1] != 1 || b.Shape[1] != 1 {
		return nil, errors.New("both Areis must be 1D for dot product")
	}

	// Check if data is present in both Areis
	if len(a.Data) == 0 || len(b.Data) == 0 {
		return nil, errors.New("input Areis cannot be empty")
	}

	lengthA, lengthB := a.Shape[0], b.Shape[0]

	// Initialize result Arei
	result := &Arei{}

	if lengthA == 1 && lengthB != 1 {
		// Scalar-vector multiplication
		result = &Arei{
			Shape: []int{lengthB},
			Data:  make([]float64, lengthB),
		}
		for i := 0; i < lengthB; i++ {
			result.Data[i] = a.Data[0] * b.Data[i]
		}
	} else if lengthA != 1 && lengthB == 1 {
		// Vector-scalar multiplication
		result = &Arei{
			Shape: []int{lengthA},
			Data:  make([]float64, lengthA),
		}
		for i := 0; i < lengthA; i++ {
			result.Data[i] = a.Data[i] * b.Data[0]
		}
	} else if lengthA == 1 && lengthB == 1 {
		// Vector-vector dot product
		sum := 0.0
		for i := 0; i < lengthA; i++ {
			sum += a.Data[i] * b.Data[i]
		}
		result = &Arei{
			Shape: []int{1},
			Data:  []float64{sum},
		}
	} else {
		return nil, errors.New("invalid dimensions for dot product")
	}

	return result, nil
}
