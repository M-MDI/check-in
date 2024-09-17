package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Myprint(str string) {
	for i := 0; i < len(str); i++ {
		z01.PrintRune(rune(str[i]))
	}
}

func main() {
	if len(os.Args) != 3 {
		return
	}
	first, second := os.Args[1], os.Args[2]
	j := 0
	for i := 0; i < len(first); i++ {
		found := false
		for j < len(second) {
			if first[i] == second[j] {
				found = true
				j++
				break
			}
			j++
		}
		if !found {
			return
		}
	}

	Myprint(first)
}
