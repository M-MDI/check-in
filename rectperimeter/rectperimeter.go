package main

import (
	"fmt"
)

func main() {
	fmt.Println(RectPerimeter(10, 2))
	fmt.Println(RectPerimeter(434343, 898989))
	fmt.Println(RectPerimeter(10, -2))
}

func RectPerimeter(w, h int) int {
	
	count := 2*(w+h)

	return count
}