package main

import (
	"fmt"
	"os"
)

func hiddenp(s1, s2 string) int {
	if s1 == "" {
		return 1
	}

	j := 0
	for i := 0; i < len(s1); i++ {
		for j < len(s2) && s2[j] != s1[i] {
			j++
		}
		if j == len(s2) { // If we reached the end of s2 and didn't find the character
			return 0
		}
		j++
	}

	return 1
}

func main() {
	if len(os.Args) != 3 {
		return
	}

	result := hiddenp(os.Args[1], os.Args[2])
	fmt.Println(result)
}
