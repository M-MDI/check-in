package main

import "fmt"

func main() {
	fmt.Println(SaveAndMiss("123456789", 3))                     // Output: 123789
	fmt.Println(SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3))     // Output: abcghimnostuz
	fmt.Println(SaveAndMiss("", 3))                              // Output: (empty string)
	fmt.Println(SaveAndMiss("hello you all ! ", 0))              // Output: hello you all ! 
	fmt.Println(SaveAndMiss("what is your name?", 0))            // Output: what is your name?
	fmt.Println(SaveAndMiss("go Exercise Save and Miss", -5))    // Output: go Exercise Save and Miss
}


func SaveAndMiss(arg string, num int) string {
	if num <= 0 {
		return arg
	}

	result := ""
	for i := 0; i < len(arg); i += 2 * num {
		if i+num < len(arg) {
			result += arg[i : i+num]
		} else {
			result += arg[i:]
		}
	}

	return result
}
