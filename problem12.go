package main

func Problem12() int {
	var index int = 0
	for {
		triangle := generateNthTriangle(index)
		factors := calculateFactors(triangle)
		if len(factors) > 500 {
			return index
		}
		index++
	}
	return 0
}

func generateNthTriangle(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}