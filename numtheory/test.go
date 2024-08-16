package numtheory

import "log"

// Test module

func Test_1() {
	// Test fib sequence

	n := 10
	figSequence := fibSeq(n)
	log.Println(figSequence)

	nthFib := fib(n)
	log.Println("Fib int at n=", n, nthFib)
}
