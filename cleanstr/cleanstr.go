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
		if char ==' ' || char =='\t' {
			if i+1 < lenght && (inputString[i+1]==' ' || inputString[i+1] =='\t'){
				continue
			}
		}
		result+=string(char) 
	}
	if result[0] ==' ' || result[0] =='\t'{
		result = result[1:]
	}
	if result[len(result)-1] ==' ' || result[len(result)-1] =='\t'{
		result = result[:len(result)-1] 
	}
fmt.Println(result)
}
















/*func main() {
	var result string
	inputString := os.Args[1]
	

	if len(os.Args) != 2 {
		fmt.Println("usage : go run main.go Hello World from GO")
		return
	}
	for i := 0; i < len(inputString); i++ {
		char := inputString[i]
		if char == ' ' || char == '\t' {
			if i+1 < len(inputString) && (inputString[i+1] == ' ' || inputString[i+1] == '\t') {
				continue
			}
		}
		result += string(char)
	}
	if result[0] == ' ' || result[0] == '\t' {
		result = result[1:]
	}
	if result[len(result)-1]  == ' ' || result[len(result)-1] == '\t' {
		result = result[:len(result)-1] 
	}
	fmt.Println(result)
}*/
