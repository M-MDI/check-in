package main

import (
	"fmt"

)

func main() {
    fmt.Println(Itoa(12345))
    fmt.Println(Itoa(0))
    fmt.Println(Itoa(-1234))
    fmt.Println(Itoa(987654321))
}
func Itoa(n int) string {
	result := ""
	sign := ""
	
	if n == 0 {
		return "0"
	}

	if n < 0 {
		sign = "-"
		n = -n
	}

	for n > 0 {
		digit := n % 10
		result = string(digit+'0') + result
		n /= 10
	}

	return sign + result
}





//decimal to hex ´
/*
value decimal / 16 = C +  







*/