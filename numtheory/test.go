package numtheory

import (
	"log"
	"time"
)

// Test module

func Test_1() {
	// Test fib sequence

	n := 10000000
	startTime := time.Now()
	fibn := Fib(n)
	endTime := time.Now()

	exeTime := endTime.Sub(startTime)
	log.Println("n/fibN/time", n, fibn, exeTime)
}
