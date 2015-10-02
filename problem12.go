package main

import (
	"math"
)

// Problem12
// What is the value of the first triangle number to have over five hundred divisors?
func Problem12() interface{} {

	var index int = 1
	var triangle int = 0
	for {

		triangle += index
		factors := calculateFactors(triangle)
		if len(factors) > 500 {
			return triangle
		}

		index++
	}
	return 0
}

func calculateFactors(n int) []int {
	var factors []int

	for i := 1; i <= int(math.Sqrt(float64(n))+1); i++ {
		if n%i == 0 {
			factors = append(factors, i)
			factors = append(factors, n/i)
		}
	}

	return factors
}
