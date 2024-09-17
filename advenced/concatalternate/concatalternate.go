package main

import "fmt"

func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}
func ConcatAlternate(slice1, slice2 []int) []int {
    len1, len2 := len(slice1), len(slice2)

    maxLen := len1

    Alen := len1+len2

    if len2 > len1 {
        maxLen = len2
    }

    result := make([]int, 0, Alen)
   
    for i := 0; i < maxLen; i++ {
        if len2 > len1 {
            if i < len2 {
                result = append(result, slice2[i])
            }
            if i < len1 {
                result = append(result, slice1[i])
            }
        } else {
            if i < len1 {
                result = append(result, slice1[i])
            }
            if i < len2 {
                result = append(result, slice2[i])
            }
        }
    }
    return result
}