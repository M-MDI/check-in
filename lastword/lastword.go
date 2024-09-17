package main

import "fmt"

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum     "))
	fmt.Print(LastWord(" "))
}

func LastWord(s string) string {
n := len(s)
	// Start from the end of the string and skip trailing spaces
	i := n - 1
	for i >= 0 && s[i] == ' ' {
		i--
	}
	// If there are no non-space characters, return just a newline
	if i < 0 {
		return "\n"
	}
	// Find the start of the last word
	end := i
	for i >= 0 && s[i] != ' ' {
		i--
	}
	// Return the last word followed by a newline
	return s[i+1:end+1] + "\n"
}
