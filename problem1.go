package main

// problem1
// sum all multiples of 3 and 5 under 1000
func Problem1() interface{} {
	sum := 0
	for i := 0; i < 1000; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			sum += i
		}
	}
	return sum
}