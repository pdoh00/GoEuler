package main

import "strconv"

// problem 4
// Find the largest palindrome made from the product of two 3-digit numbers.
func Problem4() int {

	var tempBiggestPalindrome int
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			product := i * j
			if isPalindrome(strconv.Itoa(product)) && product > tempBiggestPalindrome {
				tempBiggestPalindrome = product
			}
		}
	}
	return tempBiggestPalindrome
}

// isPalindrome determines if the string s is a palindrome.
func isPalindrome(s string) bool {
	for i := 0; i <= len(s) / 2; i++ {
		if s[i] != s[len(s) - (i + 1)] {
			return false
		}
	}
	return true
}