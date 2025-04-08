package solution

func FifthAndSkip(str string) string {
	if len(str) == 0 {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Input\n"
	}
	
	result := ""
	count := 0
	skipNext := false
	
	for _, char := range str {
		if skipNext {
			skipNext = false
			continue
		}
		result += string(char)
		count++
		if count == 5 {
			result += " "
			count = 0
			skipNext = true
		}
	}
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}
	
	return result + "\n"
}