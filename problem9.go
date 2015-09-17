package main

// problem 9
// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.
func Problem9() int {
	for a := 0; a < 1000; a++ {
		for b := 0; b < 1000; b++ {
			for c := 0; c < 1000; c++ {
				if a < b && b < c &&
				a + b + c == 1000 &&
				a * a + b * b == c * c {
					return a * b * c
				}
			}
		}
	}
	return 0
}