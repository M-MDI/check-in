package main

import (
	"fmt"
)

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}
// 	(ASCII of current character + size of the string) % 127, ensuring the result falls within the ASCII range of 0 to 127
func HashCode(dec string) string {
	lenght     := len(dec)
	result     := ""
	
	for i := 0 ; i < lenght ; i++ {
		asciiValue := int(dec[i])
				if asciiValue < 32 && asciiValue > 126 {
					asciiValue = asciiValue + 33
				} 
		newascii := (asciiValue+lenght)%127
		result += string(newascii)
	}
return result
}