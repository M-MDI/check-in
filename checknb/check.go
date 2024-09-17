package main


import (
	"fmt"
)

func main() {
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello1"))
}

func CheckNumber(s string) bool {
	
	res := true
	
	for i := 0 ; i <len(s) ;i++ {
		if s[i] < '0' ||  s[i] > '9' {
			res = true
		} else {
			res = false
		}
	}
return res
}