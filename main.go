package main

import (
	"GoNumerica/numtheory"
	"log"
)

func main() {
	log.Println("Hello GoNumerica")

	// dataset, err := arei.ReadDataFile("data/iris_converted.data")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(dataset.Shape)
	// // Extract feature column from dataset
	// features, err := arei.Column(dataset, -1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("Feature column:")
	// features.Frame()

	// input, err := arei.RemoveColumn(dataset, -1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("Input Data:")
	// input.Frame()

	numtheory.Test_2()
}
