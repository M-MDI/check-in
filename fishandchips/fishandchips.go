package main

import "fmt"

func main() {
	fmt.Println(FishAndChips(4))
	fmt.Println(FishAndChips(9))
	fmt.Println(FishAndChips(6))
	fmt.Println(FishAndChips(16))
	fmt.Println(FishAndChips(-65468))
}

func FishAndChips(n int) string {

	if n <= 0 {
			return"error: number is negative."
	}	

		if (n%3 == 0 && n%2 == 0){
			return "fishandships"
		} else if n%2 == 0 {
			return "fish"
		} else if n%3 == 0 {
			return "ships"
		} else { 
			return "error: non divisible."		
		}
	}
