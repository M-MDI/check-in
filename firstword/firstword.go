package main

import "fmt"

func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("hello   .........  bye"))
}

func FirstWord(s string) string {
	/*
	result := ""

		for i := 0; i < len(s); i++ {
			if s[i] == ' ' {
					break
				}
			result += string(s[i])
		} 
	return result + "\n"
	*/
	result := ""
	end := len(s) - 1
/*	// Skip trailing spaces
	for end >= 0 && s[end] == ' ' {
		end--
	}
*/
	// Collect the last word
	for i := end; i >= 0; i-- {
		if s[i] == ' ' {
			break
		}
		result = string(s[i]) + result
	}

	return result + "\n"
}

