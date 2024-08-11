package notes

// Introduction summary
// A major application of linear algebra is to solving systems of linear equations.
// This lecture presents three ways of thinking about these systems.
// The “row method” focuses on the individual equations,
// the “column method” focuses on combining the columns, and the “matrix method”
// is an even more compact and powerful way of describing systems of linear equations.
// URL https://ocw.mit.edu/courses/18-06-linear-algebra-spring-2010/resources/lecture-1-the-geometry-of-linear-equations/
// Ax = b
// A = matrix, x = vector, b = result vector

func Example1() []int {
	A := [][]int{
		//Col
		{2, 5}, //Row
		{1, 3},
	}
	x := []int{1, 2}

	var b []int

	// Calculate matrix manually
	aTopLeft := A[0][0] * x[0]
	aBottomLeft := A[1][0] * x[0]
	aTopRight := A[0][1] * x[1]
	aBottomRight := A[1][1] * x[1]

	aTop := aTopLeft + aTopRight
	aBottom := aBottomLeft + aBottomRight

	b = append(b, aTop, aBottom)
	return b
}
