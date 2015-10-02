package main

// problem 7
// What is the 10 001st prime number?
func Problem7() interface{} {
	var primeIndex int = 1000
	for i := primes[len(primes) - 1];; i++ {
		if isLikelyPrime(i) {
			primeIndex ++
			if primeIndex == 10001 {
				return i
			}
		}
	}
	return 0
}