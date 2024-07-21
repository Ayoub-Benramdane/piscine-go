package piscine

func JumpOver(str string) string {
	if len(str) < 3 {
		return "\n"
	}
	var result string
	for i := 2; i < len(str); i += 3 {
		result += string(str[i])
	}
	return result + "\n"
}
