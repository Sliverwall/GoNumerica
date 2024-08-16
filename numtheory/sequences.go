package numtheory

// Module to store functions that generate int sequences

//--------------------------- FIBONACCI METHODS ---------------------------------------

// Fib returns the nth element in a fib sequence using matrix exponentiation. No intermediate array is created.
func Fib(n int) uint64 {
	if n < 2 {
		return uint64(n)
	}
	fibMatrix := [][]int{
		{1, 1},
		{1, 0},
	}
	//the [0][0] position will be = to F(n+1)
	result := MatrixPow(fibMatrix, n-1)

	return uint64(result[0][0])
}

// FibIter returns the nth element in a fib sequence using iteration techinque. No intermediate array is created.
func FibIter(n int) uint64 {
	// Manually set seed values for index 0 and 1 as fib[0] = 1, fib[1] = 1
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	// Keep track of the value of fib[i-2] and fib[i-1] starting at i = 2
	var n0, n1 uint64 = 1, 1
	// initialize the result variable as an int
	var result uint64
	// Start at i= 2 because first two elements solved manually
	for i := 2; i < n; i++ {
		// the nth result of fib sequence is the sum of the previous two elements when i >= 2
		result = n1 + n0
		// Move up our fib[i-2] and fib[i-1] trackers without storing the sequence
		n0 = n1
		n1 = result
	}
	// Return the result after looping n times
	return result
}

// FibSeq takes a given n using the iter method, then returns the sequence
func FibSeq(n int) []uint64 {
	if n <= 0 {
		return []uint64{}
	}

	// Create a slice to hold the Fibonacci sequence
	fibSequence := make([]uint64, n)

	// Manually set seed values for index 0 and 1 as fib[0] = 1, fib[1] = 1
	if n > 0 {
		fibSequence[0] = 1
	}
	if n > 1 {
		fibSequence[1] = 1
	}

	// After i = 0 and i = 1, the ith input will be the sum of the last two elements
	for i := 2; i < n; i++ {
		// i = (i-1) + (i-2)
		fibSequence[i] = fibSequence[i-1] + fibSequence[i-2]
	}

	// Return the int array of the sequence
	return fibSequence
}
