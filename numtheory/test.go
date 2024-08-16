package numtheory

import (
	"GoNumerica/arei"
	"log"
)

// Test module

func Test_1() {
	// Test fib sequence

	n := 20
	fibSequence := fibSeq(n)
	log.Println(fibSequence)

	nthFib := fib(n)
	log.Println("Fib int at n=", n, nthFib)

	// Convert sequence into float64
	floatFigSeq := ConvertIntArrToFloat64Arr(fibSequence)

	figSeqArei, _ := arei.NewArei(floatFigSeq)
	m := 4
	figSeqArei.Reshape([]int{m, len(floatFigSeq) / m})
	figSeqArei.Frame()

}
