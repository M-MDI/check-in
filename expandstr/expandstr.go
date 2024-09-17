package main

import (
	"fmt"
	"os"
)

func main() {
	var result string
	inputString := os.Args[1]
	lenght := len(inputString)

	if len(os.Args) != 2 {
		fmt.Println("usage : go run main.go Hello World from GO")
		return
	}

	for i := 0; i < len(inputString); i++ {
		char := inputString[i]
		if char == ' ' || char == '\t' {
			if i+1 < lenght && (inputString[i+1] == ' ' || inputString[i+1] == '\t') {
				continue
			} else {
				result += "  "
			}
		}
		result += string(char)
	}
	if len(result) >= 3 && result[:3] == "   " {
		result = result[3:]
	}
	if len(result) >= 3 && result[:len(result)-3] == "   " {
		result = result[:len(result)-3]
	}
	fmt.Println(result)
}
