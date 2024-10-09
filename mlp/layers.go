package mlp

import "GoNumerica/arei"

type Layer struct {
	// Store shape
	Shape []int
	// Store incoming weights and biases
	Weights arei.Arei
	Biases  arei.Arei
	// Store outputs and activated outputs
	Outputs     arei.Arei
	Activations arei.Arei
	// Store the cost gradient in respect to weights and biases
	CgWeights arei.Arei
	CgBiases  arei.Arei
}

func (l *Layer) InitParams() {
	// Use full shape for weights. m = number of neruons, n = number of previous layer's neurons
	l.Weights, _ = arei.Zeros(l.Shape)
	// Biases are always flat vectors. m = number of neurons, n = 1
	l.Biases, _ = arei.Zeros((l.Shape[0],1))
}
