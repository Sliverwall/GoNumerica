package numtheory

import (
	"log"
	"time"
)

// Test module

func Test_1() {
	// Test fib sequence

	n := 1000
	startTime := time.Now()
	fibn := Fib(n)
	endTime := time.Now()

	exeTime := endTime.Sub(startTime)
	log.Println("n/fibN/time", n, fibn, exeTime)
}

func Test_2() {
	// Find all primes 1 to 100

	primes := make([]int, 0)
	for i := 2; i < 101; i++ {
		if IsPrime(i) {
			primes = append(primes, i)
		}
	}

	log.Println(primes)

	n := 9742
	log.Println(n, "Is Prime:", IsPrime(n))
}
