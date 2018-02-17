package diffsquares

import (
	"math"
)

const square = 2.0

// SquareOfSums calculates the square of the sum of the first n numbers
func SquareOfSums(n int) int {
	return int(math.Pow(float64(n*(n+1)/2), square))
}

// SumOfSquares squares each of the first n numbers and adds them up
func SumOfSquares(n int) int {
	var sum float64
	for i := 1; i <= n; i++ {
		sum += math.Pow(float64(i), square)
	}
	return int(sum)
}

// Difference returns the square of sums minus the sum of the squares.
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
