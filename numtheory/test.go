package numtheory

import (
	"log"
	"time"
)

// Test module

func Test_1() {
	// Test fib sequence

	n := 1000000
	startTime := time.Now()
	nthFib := Fib(n)
	endTime := time.Now()

	exeTime := endTime.Sub(startTime)
	log.Println("n/fibN/time", n, nthFib, exeTime)
}
