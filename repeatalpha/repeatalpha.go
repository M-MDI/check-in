package main

import "fmt"

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}

func RepeatAlpha(s string) string {
    result := ""
    length := len(s)
    index  := 0

    for i := 0; i < length; i++ {
        // Check if it's an alphabetic character
        if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
            // Calculate the alphabetical index
            if s[i] >= 'a' && s[i] <= 'z' {
                index = int(s[i] - 'a' + 1)
            } else {
                index = int(s[i] - 'A' + 1)
            }
            // Repeat the character based on its index
            for j := 0; j < index; j++ {
                result += string(s[i])
            }
        } else {
            // Non-alphabetic characters are added as is
            result += string(s[i])
        }
    }
    
    return result
}