package main

import (
    "fmt"
)

func main() {
    fmt.Println(ZipString("YouuungFellllas"))
    fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
    fmt.Println(ZipString("Helloo Therre!"))
}

func ZipString(s string) string {
    if len(s) == 0 {
        return ""
    }

    var result string
    count := 1
    current := s[0]

    for i := 1; i < len(s); i++ {
        if s[i] == current {
            count++
        } else {
            result += fmt.Sprintf("%d%c", count, current)
            count = 1
            current = s[i]
        }
    }

    result += fmt.Sprintf("%d%c", count, current)
    return result
}