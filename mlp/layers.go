package mlp

import (
	"GoNumerica/arei"
)

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
	// Store layer type
	LayerType int
}

// Make enum to keep track of types of layers
const (
	InputLayer = iota
	HiddenLayer
	OutputLayer
)

// --------------Init Methods-------------------------//
// Constructor for Layer type
func NewLayer(shape []int, activationFunction string, layerType int) *Layer {

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
		LayerType:   layerType,
	}
}

// BuildLayers takes in a ks2d array of shapes and 1d array of activationFunctions to build a LayerMap for neural net structure.
func BuildLayers(shapes [][]int, activationFunctions []string) []*Layer {
	// Init return layerMap to fit shapes height
	numLayers := len(activationFunctions)
	layerMap := make([]*Layer, numLayers)

	// Declare layerType since it inits in switch statement
	var layerType int
	// Loop through each row
	for i := range numLayers {
		// use switch case to init layer type
		switch i {
		// First layer should be input type
		case 0:
			layerType = InputLayer
		// Last layer should be an output
		case numLayers - 1:
			layerType = OutputLayer

		// Otherwise hidden
		default:
			layerType = HiddenLayer
		}

		// Should be 2 columns: 0th for row dimensions and 1th for column dimensions
		shape := []int{shapes[i][0], shapes[i][1]}   // Extract shape
		activationFunction := activationFunctions[i] // Extract activation function
		layer := NewLayer(shape, activationFunction, layerType)
		layerMap[i] = layer
	}

	// Return completed layerMap
	return layerMap
}

// --------------Activation Functions-------------------------//
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

// Activate uses a layer's activation type to determine which element-wise activation function to apply to z
func (l *Layer) Activate(z *arei.Arei) *arei.Arei {

	var a *arei.Arei
	switch l.ActivationFunction {
	case "relu":
		// Element-wise function. Checks if element > 0, element, otherwise 0
		a = arei.Maximum(z, 0)
	case "softmax":
		a = arei.SoftMax(z, 0)
	}
	// Return activated arei
	return a
}
