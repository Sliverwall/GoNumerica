package numtheory

// Module to store functions that generate int sequences

// fibSeq takes a given n, then returns
func fibSeq(n int) []int {
	if n <= 0 {
		return []int{}
	}

	// Create a slice to hold the Fibonacci sequence
	fibSequence := make([]int, n)

	// Initialize the first two Fibonacci numbers
	if n > 0 {
		fibSequence[0] = 0
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

// Fib returns the nth element in a fib sequence. No intermediate array is created.
func fib(n int) int {
	// Manually return first two results in the fib sequence fib[0] = 0, fib[1] = 1
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	// Keep track of the value of fib[i-2] and fib[i-1] starting at i = 2
	n0 := 0
	n1 := 1
	// initialize the result variable as an int
	var result int
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
