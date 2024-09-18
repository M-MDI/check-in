package main

import (
	"fmt"
	"strconv"
)

func fprime(n int) string {
	if n <= 1 {
		return ""
	}

	result := ""
	for i := 2; i <= n && n != 0; i++ {
		if n%i == 0 {
			if result != "" {
				result += "*"
			}
			result += strconv.Itoa(i)
			n /= i
			i--
		}
	}
	return result
}

func main() {
	fmt.Println(fprime(225225))
}
