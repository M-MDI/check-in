package main

import (
	"github.com/01-edu/z01"
)

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21})
}
func Myprint(str string) {
	for i := 0; i < len(str); i++ {
		z01.PrintRune(rune(str[i]))
	}
}

func PrintMemory(arr [10]byte) {

	res := ""
	x := "0123456789abcdef"
	
	for i := 0 ; i < len(arr) ; i++{
		res += string(x[int(arr[i]/16)])
		res += string(x[int(arr[i]%16)])
			if i == 3 || i == 7 {
				res += "\n"
			} else {
				res += " "
			}
		}
		Myprint(res)
		z01.PrintRune('\n')	
		for i := 0 ; i < len(arr) ; i++ {
			if arr[i] > 23 && arr[i] < 127 {
				Myprint(string(arr[i]))
			} else {
				Myprint(".")
			}
		} 
		z01.PrintRune('\n')
}
