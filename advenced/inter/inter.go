package main

import (
	"fmt"
	"os"
)

func inter(s1, s2 string) string {
	result := ""
	s2Chars := make(map[rune]int)
		for i := 0; i < len(s1); i++ {
			char := rune(s1[i])
			for _, char2 := range s2 {
				if char == char2 {
					s2Chars[char]++
						if s2Chars[char] == 1 {
							result += string(char)
						}
				i++
			}
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
