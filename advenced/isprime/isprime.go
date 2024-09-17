package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsPrime(5))   // true
	fmt.Println(IsPrime(4))   // false
	fmt.Println(IsPrime(13))  // true
	fmt.Println(IsPrime(1))   // false
	fmt.Println(IsPrime(100)) // false
}

// IsPrime checks if a number is prime
func IsPrime(nb int) bool {
	// 1 is not a prime number
	if nb < 2 {
		return false
	}

	// 2 is the only even prime number
	if nb == 2 {
		return true
	}

	// Any other even number is not prime
	if nb%2 == 0 {
		return false
	}
	for i := 3; i*i <= nb; i += 2 {
		if nb%i == 0 {
			return false
		}
	}

	// If no divisors were found, the number is prime
	return true
}
