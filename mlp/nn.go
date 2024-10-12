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
func (nn *NeuralNet) backwardProp(y *arei.Arei) {
	numLayer := len(nn.LayerMap)

	// Start from last layer, then move backwards until last layer before input
	for i := (numLayer - 1); i > 0; i-- {
		layer := nn.LayerMap[i]
		leftLayer := nn.LayerMap[i-1]

		// Check layer type
		if layer.LayerType == OutputLayer {
			// If output layer, use cross entropy since softmax is used
			// Step 1
			yHat := layer.Activations
			layer.CgActivations = layer.CrossEntropy(yHat, y)
		} else {
			// Else if not use the deriv of the activation function
			rightLayer := nn.LayerMap[i+1]
			// Backprop the error to the hidden layers
			derivZ := layer.DerivActivationFunction(layer.Outputs)
			// Do element-wise multiplication of derivZ and propgated error
			layer.CgActivations, _ = arei.Multi(rightLayer.CgActivations.Dot(rightLayer.Weights.T()), derivZ)
		}
		// Step 2 cost gradients for cost function in respect to weights and biases
		layer.CgWeights = leftLayer.Activations.T().Dot(layer.CgActivations)
		layer.CgBiases = arei.RowWiseSum(layer.CgActivations)
	}
}

func (nn *NeuralNet) updateParams() {
	// Init number of layers
	numLayers := len(nn.LayerMap)

	// Loop through each layer, exclude input layer
	for i := 1; i < numLayers; i++ {
		layer := nn.LayerMap[i]

		// Get cost gradients scaled by step size
		weightStep := arei.MultiT(layer.CgWeights, nn.LearningRate)
		biasStep := arei.MultiT(layer.CgBiases, nn.LearningRate)

		// Subtract step from weight and bias to get updated weights and biases
		layer.Weights, _ = arei.Sub(layer.Weights, weightStep)
		layer.Biases, _ = arei.Sub(layer.Biases, biasStep)
	}
}

func (nn *NeuralNet) Fit(X, y *arei.Arei) {

	// Loop through epochs
	for i := range nn.Epochs {
		// Msg to confirm epoch has started
		log.Println("Epoch", i, "of", nn.Epochs, "Start")

		// Feed input data as the activation for layer 0
		nn.LayerMap[0].Activations = X // Shape shold be m x n, or the number of features x number of samples
		// Perform forward pass
		nn.forwardProp()
		// Perform backward pass, Should contain ground truth y
		nn.backwardProp(y)
		// Update parameters after calculating the cost gradient of error in respect to each parameter
		nn.updateParams()
		// Msg to confirm epoch has ended
		log.Println("Epoch", i, "of", nn.Epochs, "End")
	}
}

func (nn *NeuralNet) Predict(X *arei.Arei) *arei.Arei {
	// Set input layer's activation as sample
	nn.LayerMap[0].Activations = X
	// Do forward pass
	nn.forwardProp()
	// Return activated output
	outputLayer := nn.LayerMap[len(nn.LayerMap)-1]
	return outputLayer.Activations
}
