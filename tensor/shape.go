package tensor

import "errors"

// Functions to create special kind of matrixes

// Identity creates an identity tensor with the same shape as the given tensor
func Identity(tensor *Tensor) (*Tensor, error) {
	switch len(tensor.shape) {
	case 1:
		// Identity vector (scalar 1)
		identityData := make([]float64, tensor.shape[0])
		for i := range identityData {
			identityData[i] = 1.0
		}
		return &Tensor{shape: tensor.shape, data: identityData}, nil
	case 2:
		if tensor.shape[0] != tensor.shape[1] {
			return nil, errors.New("cannot create an identity matrix for a non-square matrix")
		}
		identityData := make([]float64, tensor.shape[0]*tensor.shape[1])
		for i := 0; i < tensor.shape[0]; i++ {
			identityData[i*tensor.shape[0]+i] = 1.0
		}
		return &Tensor{shape: tensor.shape, data: identityData}, nil
	default:
		return nil, errors.New("invalid shape for identity tensor")
	}
}
