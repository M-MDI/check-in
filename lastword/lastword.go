package main

import "fmt"

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum     "))
	fmt.Print(LastWord(" "))
}

func LastWord(s string) string {
n := len(s)

	i := n - 1
	for i >= 0 && s[i] == ' ' {
		i--
	}
	if i < 0 {
		return "\n"
	}
	
	end := i
	for i >= 0 && s[i] != ' ' {
		i--
	}
	// Return the last word followed by a newline
	return s[i+1:end+1] + "\n"
}
