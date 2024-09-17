package main

import (
	"fmt"
)

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}

func WeAreUnique(str1, str2 string) int {
    if str1 == "" && str2 == "" {
        return -1
    }

    chars := make(map[rune]bool)
    unique := make(map[rune]bool)

    for _, char := range str1 + str2 {
        if !chars[char] {
            chars[char] = true
            unique[char] = true
        } else {
            delete(unique, char)
        }
    }

    return len(unique)
}