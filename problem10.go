package main

// problem 10
// Find the sum of all the primes below two million.
func Problem10() interface{} {
	sum := 0
	//initialize sum with all the known primes
	for _, num := range primes {
		sum += num
	}

	//loop through all the numbers above the last know prime and two million
	for i := primes[len(primes) - 1] + 1; i < 2000000; i++ {
		if isLikelyPrime(i) {
			sum += i
		}
	}
	return sum
}