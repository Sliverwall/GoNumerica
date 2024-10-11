package mlp

import "log"

func Test_1() {
	// Test building a layer

	shapes := [][]int{
		// Column 1 = number of neruons, column 2 = number of connections (1 connection per neuron in previous layer)
		{10, 1},  //input layer
		{20, 10}, //hidden layer-> 20 neurons connecting to 10 input layers
		{20, 20},
		{2, 20}, // Output layer -> 2 output neurons connecting to 20 hidden neurons
	}

	activationFunctions := []string{
		"None", // Input not being activated
		"relu", // Hidden will use relu
		"relu",
		"softmax", // Output will use softmax
	}
	// build layers
	layerMap := BuildLayers(shapes, activationFunctions)

	// connect to nn
	model := NewNeuralNet(1, 32, 0.1, layerMap)
	// Loop through map and paste some metrics

	log.Println(len(layerMap))
	for i := range len(model.LayerMap) {
		layer := model.LayerMap[i]

		log.Println(i, layer.LayerType, layer.Weights.Shape, layer.ActivationFunction)
	}
}
