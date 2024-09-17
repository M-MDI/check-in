package main

import (
	"fmt"
	"os"
)

func inter(s1, s2 string) string {
	seen := make(map[rune]bool)
	result := ""

	// Create a map to store characters of s2
	s2Chars := make(map[rune]bool)
	for _, char := range s2 {
		s2Chars[char] = true
	}

	// Iterate through s1 and check if the character exists in s2
	for _, char := range s1 {
		if s2Chars[char] && !seen[char] {
			result += string(char)
			seen[char] = true // Mark character as seen to avoid duplicates
		}
	}

	return result
}

func main() {
	if len(os.Args) != 3 {
		return
	}

	fmt.Println(inter(os.Args[1], os.Args[2]))
}
