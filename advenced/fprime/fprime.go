package main

import (
	"fmt"
	"os"
	"strconv"
)

func fprime(n int) string {
	if n <= 1 {
		return ""
	}

	result := ""
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			if result != "" {
				result += "*"
			}
			result += strconv.Itoa(i)
			n /= i
		}
	}
	if n > 1 {
		if result != "" {
			result += "*"
		}
		result += strconv.Itoa(n)
	}

	return result
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	// Convert argument to integer
	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num <= 0 {
		return
	}

	// Get prime factors and print
	result := fprime(num)
	if result != "" {
		fmt.Println(result)
	}
}

