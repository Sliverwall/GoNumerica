package mlutils

import (
	"GoNumerica/arei"
	"log"
)

func Test_1() {
	// Test label encoder

	// Init labels between 0 and 9 ints
	labels := []int{0, 0, 0, 1, 2, 2, 3, 4, 5, 5, 5, 6, 7, 7, 8, 7, 9, 9, 5, 4, 4, 3, 0}
	// Create 1d arei with labels and name as y
	y, _ := arei.NewArei(labels)

	log.Println("Original labels:")
	y.Frame()
	oneHotY, _ := LabelEncode(y)
	// Print oneHotY
	log.Println("Encoded labels:")
	oneHotY.Frame()

}
