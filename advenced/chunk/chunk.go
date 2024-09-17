package main

import "fmt"

func Chunk(slice []int, size int) {
	
	sl := []int{}
	
	var res [][]int
	
	if size == 0 {
		return
	}
	
	for i := 0; i < len(slice); i++ {
		if i%size == 0 {
			for j := i; j < i+size && j < len(slice); j++ {
				sl = append(sl, slice[j])
			}
			res = append(res, sl)
		}
	}
	fmt.Println(res)
}

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}
