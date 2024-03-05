package piscine

func IsAlpha(s string) bool {
	text := []rune(s)
	for i := 0; i < len(text); i++ {
		if (text[i] < 'A' || text[i] > 'Z') && (text[i] < 'a' || text[i] > 'z') && (text[i] < '0' || text[i] > '9') {
			return false
		}
	}
	return true
}
