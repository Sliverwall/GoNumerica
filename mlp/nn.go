package mlp

import (
	"GoNumerica/arei"
	"log"
)

// Define the Neurel Net structure

type NeuralNet struct {

	// Define hyperparameters
	Epochs       int
	BatchSize    int
	LearningRate float64

	// Define slice to hold layers
	LayerMap []*Layer
}

// Constructor for NeuralNet
func NewNeuralNet(epochs int, batchSize int, learningRate float64, layerMap []*Layer) *NeuralNet {
	return &NeuralNet{
		Epochs:       epochs,       // Number of epochs to perform
		BatchSize:    batchSize,    // Number of samples to process each pass
		LearningRate: learningRate, // Step rate while performing gradient descent
		LayerMap:     layerMap,     // Slice to hold each indiviudal Layer
	}
}

// forward_prop method to have neural net perform a forward propagation
func (nn *NeuralNet) forwardProp() {

	// Init number of layers
	numLayers := len(nn.LayerMap)

	// Loop through each layer, exclude input layer
	for i := 1; i < numLayers; i++ {
		// Get current layer
		layer := nn.LayerMap[i]
		// Get previous layer
		previousLayer := nn.LayerMap[i-1]
		// Calculate output
		var err error
		layer.Outputs, err = arei.Sum((layer.Weights.Dot(previousLayer.Activations)), layer.Biases)
		if err != nil {
			log.Fatal(err)
		}
		// Use layer's activation function to set layer's activation matrix
		layer.Activations = layer.Activate(layer.Outputs)

	}
}

// backwardProp travels from output layer until the last hidden layer to calculate all needed cost gradients
func (nn *NeuralNet) backwardProp(yHat, y *arei.Arei) {
	numLayer := len(nn.LayerMap)

	// Start from last layer, then move backwards until last layer before input
	for i := (numLayer - 1); i > 0; i-- {
		layer := nn.LayerMap[i]
		leftLayer := nn.LayerMap[i-1]

		// Check layer type
		if layer.LayerType == OutputLayer {
			// If output layer, use cross entropy since softmax is used
			// Step 1
			layer.CgActivations = layer.CrossEntropy(yHat, y)

			// Step 2
			layer.CgWeights = leftLayer.Activations.T().Dot(layer.CgActivations)
		} else {
			// Else if not use the deriv of the activation function
			rightLayer := nn.LayerMap[i+1]
			layer.CgActivations = rightLayer.CgActivations.Dot(rightLayer.Weights.T())
		}

	}
}
