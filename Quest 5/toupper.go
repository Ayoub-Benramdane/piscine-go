package piscine

func ToUpper(s string) string {
	text := []rune(s)
	result := ""
	for i := 0; i < len(text); i++ {
		if text[i] >= 'a' && text[i] <= 'z' {
			result += string(text[i] - 32)
		} else {
			result += string(text[i])
		}
	}
	return result
}
