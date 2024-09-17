package main

import "fmt"

func main() {
	fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(RetainFirstHalf("A"))
	fmt.Println(RetainFirstHalf(""))
	fmt.Println(RetainFirstHalf("Hello World"))
}

func RetainFirstHalf(str string) string {
	count := len(str)
	midcount := len(str) / 2
	result := ""

	if count == 0 {
		return ""
	}
	if count == 1 {
		return str
	}
	for i := 0; i < midcount; i++ {
		result += string(str[i])
	}
	return result
}
