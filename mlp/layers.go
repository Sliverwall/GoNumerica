package mlp

import "GoNumerica/arei"

// Define layer structure for neurel net

type Layer struct {
	// Store shape
	Shape []int
	// Store activation function used as string
	ActivationFunction string
	// Store incoming weights and biases
	Weights *arei.Arei
	Biases  *arei.Arei
	// Store outputs and activated outputs
	Outputs     *arei.Arei
	Activations *arei.Arei
	// Store the cost gradient in respect to weights and biases
	CgWeights *arei.Arei
	CgBiases  *arei.Arei
}

// Constructor for Layer type
func NewLayer(shape []int, activationFunction string) *Layer {
	// Use full shape for weights. m = number of neruons, n = number of previous layer's neurons
	weights, _ := arei.Zeros(shape)
	// Biases are always flat vectors. m = number of neurons, n = 1
	biasShape := []int{shape[0], 1}
	biases, _ := arei.Zeros(biasShape)

	// Return the layer object
	return &Layer{
		Shape:              shape,
		ActivationFunction: activationFunction,
		Weights:            weights,
		Biases:             biases,
		// Initialize parts dependent on input as nil
		Outputs:     nil,
		Activations: nil,
		CgWeights:   nil,
		CgBiases:    nil,
	}
}

// ReLU activation function
func (l *Layer) ReLU(a *arei.Arei) *arei.Arei {
	// Element-wise function. Checks if element > 0, element, otherwise 0
	activatedL := arei.Maximum(a, 0)
	return activatedL
}

// Deriv ReLU activation function
func (l *Layer) DerivReLU(a *arei.Arei) *arei.Arei {
	// Element-wise function. Checks if element > 0, 1, otherwise 0
	derivActivatedL := arei.Compare(a, 0)

	return derivActivatedL
}
