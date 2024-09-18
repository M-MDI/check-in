package main

import "fmt"


func CanJump(nums []uint)bool{

	for i := 0 ; i < len(nums) ;i++{
		i += int(nums[i])
		if i == len(nums)-1{
			return true 
		} 
	}
	return false
	}

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1)) // true

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2)) // false

	input3 := []uint{0}
	fmt.Println(CanJump(input3)) // true
}
