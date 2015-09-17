package main

// problem 6
// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
func Problem6() int {
	var sumOfSquares int = 0
	var squareOfSums int = 0
	for i := 1; i <= 100; i++ {
		sumOfSquares += i * i
		squareOfSums += i
	}
	squareOfSums = squareOfSums * squareOfSums
	return squareOfSums - sumOfSquares
}