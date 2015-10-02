package main

import "sort"

// problem 3
// What is the largest prime factor of the number 600851475143
func Problem3() interface{} {
	checkValue := 600851475143
	var primeFactors []int

	for {
		val, err := getFirstPrimeFactor(checkValue)
		if err != nil {
			break
		}
		primeFactors = append(primeFactors, val)
		checkValue = checkValue / val
	}
	sort.Ints(primeFactors)
	return primeFactors[len(primeFactors) - 1]
}