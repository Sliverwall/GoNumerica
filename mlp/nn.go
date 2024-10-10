package mlp

// Define the Neurel Net structure

type NeuralNet struct {

	// Define hyperparameters
	Epochs       int
	BatchSize    int
	LearningRate float64

	// Define slice to hold layers
	LayerMap []Layer
}

// Constructor for NeuralNet
func NewNeuralNet(epochs int, batchSize int, learningRate float64, layerMap []Layer) *NeuralNet {
	return &NeuralNet{
		Epochs:       epochs,       // Number of epochs to perform
		BatchSize:    batchSize,    // Number of samples to process each pass
		LearningRate: learningRate, // Step rate while performing gradient descent
		LayerMap:     layerMap,     // Slice to hold each indiviudal Layer
	}
}
