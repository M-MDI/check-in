package main

import "os"

func main() {
	// If the number of arguments is not 2, print a newline and return
	if len(os.Args) != 3 {
		println()
		return
	}

	// Concatenate the two strings into one
	s := os.Args[1] + os.Args[2]

	// Create a slice to track characters that have already been printed
	var result []byte

	// Iterate through the concatenated string
	for i := 0; i < len(s); i++ {
		// Check if the character has already been added to result
		exists := false
		for j := 0; j < len(result); j++ {
			if s[i] == result[j] {
				exists = true
				break
			}
		}
		// If the character doesn't exist in result, add it and print it
		if !exists {
			result = append(result, s[i])
			print(string(s[i]))
		}
	}

	// Print a newline at the end
	println()
}
