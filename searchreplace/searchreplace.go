package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	inputString := os.Args[1]
	charToReplace := os.Args[2]
	replacementChar := os.Args[3]

	if len(charToReplace) != 1 || len(replacementChar) != 1 {
		return
	}

	result := ""
	
	for i := 0; i < len(inputString); i++ {
		if inputString[i] == charToReplace[0] {
			result += replacementChar
		} else {
			result += string(inputString[i])
		}
	}
	fmt.Println(result)
}
