package numtheory

import "math"

// Functions to use for primes

func IsPrime(x int) bool {
	// Get the floor of the sqaure root. Then check if x divides eveningly into any number less than or equal to that floor.

	// Check 0 and 1
	if x == 0 || x == 1 {
		return false
	}
	threshold := int(math.Floor(math.Sqrt(float64(x))))

	// Skip 1
	for i := 2; i < threshold; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
