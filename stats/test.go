package stats

import "log"

func Test_1() {
	// test Stdev function
	X := []int{5, 3, 2, 2, 1}

	std := Stdev(X)

	log.Println(std)
}
