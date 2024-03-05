package piscine

func IsNumeric(s string) bool {
	text := []rune(s)
	for i := 0; i < len(text); i++ {
		if text[i] < '0' || text[i] > '9' {
			return false
		}
	}
	return true
}
