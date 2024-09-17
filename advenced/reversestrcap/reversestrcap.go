package main

import (
	"fmt"
	"os"
	"strings"
)

func reverseStrCap() {
	for i := 1; i < len(os.Args); i++ {
		words := strings.Fields(os.Args[i])
		for j := 0; j < len(words); j++ {
			word := words[j]
			word = strings.ToLower(word)
			if len(word) > 0 {
				word = word[:len(word)-1] + strings.ToUpper(string(word[len(word)-1]))
			}
			words[j] = word
		}
		fmt.Println(strings.Join(words, " "))
	}
}

func main() {
	if len(os.Args) > 1 {
		reverseStrCap()
	}
}
