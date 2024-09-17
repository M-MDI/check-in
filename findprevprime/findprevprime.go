package main

import "fmt"

func main() {
	fmt.Println(FindPrevPrime(5))  // Output: 5
	fmt.Println(FindPrevPrime(4))  // Output: 3
	fmt.Println(FindPrevPrime(10)) // Output: 7
	fmt.Println(FindPrevPrime(1))  // Output: 0
}

func FindPrevPrime(nb int) int {
	if nb < 2 {
		return 0
	}

	for i := nb; i >= 2; i-- {
		isPrime := true
		if i == 2 {
			return 2
		}

		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			return i
		}
	}

	return 0
}
