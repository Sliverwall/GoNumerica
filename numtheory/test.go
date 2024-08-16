package numtheory

import (
	"GoNumerica/arei"
	"log"
)

// Test module

func Test_1() {
	// Test fib sequence

	n := 10
	fibSequence := FibSeq(n)
	log.Println(fibSequence)

	nthFib := Fib(n)
	log.Println("Fib int at n=", n, nthFib)

	// Convert sequence into float64
	floatFigSeq := ConvertIntArrToFloat64Arr(fibSequence)

	// Create an arei matrix from the figsequence
	figSeqArei, _ := arei.NewArei(floatFigSeq)
	m := 4
	figSeqArei.Reshape([]int{m, len(floatFigSeq) / m})
	figSeqArei.Frame()

}
