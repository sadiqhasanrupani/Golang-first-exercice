package main

import "fmt"

func countPalindromicSubsequence(s string) int {
	n := len(s)
	// result := 0
	seenPalindromes := make(map[string]bool)

	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if s[i] == s[k] && s[i] != s[j] {
					palindrome := string(s[i]) + string(s[j]) + string(s[k])
					seenPalindromes[palindrome] = true
				}
			}
		}
	}

	return len(seenPalindromes)
}

func main() {
	// Example 1
	inputString1 := "aabca"
	result1 := countPalindromicSubsequence(inputString1)
	fmt.Printf("Example 1: %d\n", result1)

	// Example 2
	inputString2 := "adc"
	result2 := countPalindromicSubsequence(inputString2)
	fmt.Printf("Example 2: %d\n", result2)

	// Example 3
	inputString3 := "bbcbaba"
	result3 := countPalindromicSubsequence(inputString3)
	fmt.Printf("Example 3: %d\n", result3)
}
