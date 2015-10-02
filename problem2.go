package main

// problem2
// By considering the terms in the Fibonacci sequence whose values do not exceed
// four million, find the sum of the even-valued terms.
func Problem2() interface{} {
	sum := 0
	memo := make(map[int]int)

	for i := 0;; i++ {
		v := fib(i, memo)
		if v >= 4000000 {
			break;
		}
		if v % 2 == 0 {
			sum += v
		}
	}
	return sum
}

// fibonacci
// Fn = F(n-1) + F(n-2)
// n = nth fib
// memo is a map used to memoize previously calculated values.
func fib(n int, memo map[int]int) int {

	if n <= 1 {
		return n
	}

	var n1 int
	var n2 int
	var isPresent bool

	n1, isPresent = memo[n - 1]
	if !isPresent {
		n1 = fib(n - 1, memo)
		memo[n - 1] = n1
	}

	n2, isPresent = memo[n - 2]
	if !isPresent {
		n2 = fib(n - 2, memo)
		memo[n - 2] = n2
	}

	return n1 + n2

}