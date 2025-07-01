package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func countPalindromicSubstrings(s string) int {
	count := 0
	n := len(s)

	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			substr := s[i:j]
			if isPalindrome(substr) {
				count++
			}
		}
	}
	return count
}

func main() {
	input := "katak"
	result := countPalindromicSubstrings(input)
	fmt.Printf("Jumlah substring palindromik: %d\n", result)
}
