package main

import (
	"fmt"
)

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelCasE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
	fmt.Println(CamelToSnakeCase(""))
}

func CamelToSnakeCase(s string) string {
	result := ""
	lenght := len(s)

	if lenght < 1 {
		return ""
	}
	for i, r := range s {
		if i > 0 && s[i] >= 'A' && s[i] <= 'Z' {
			if (s[i-1] >= 'A' && s[i-1] <= 'Z') || i == len(s)-1 {
				return s
			}
			result += "_"
		}
		result += string(r)
	}
	return result
}
