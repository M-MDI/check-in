package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsPrime(2))   // true
	fmt.Println(IsPrime(4))   // false
	fmt.Println(IsPrime(13))  // true
	fmt.Println(IsPrime(1))   // false
	fmt.Println(IsPrime(100)) // false
}

// IsPrime checks if a number is prime
func IsPrime(nb int) bool {
	if nb < 2 {
		return false
	}
	for i := 2; i <= nb; i += 2 {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
