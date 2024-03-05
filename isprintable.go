package piscine

func IsPrintable(s string) bool {
	text := []rune(s)
	for i := 0; i < len(text); i++ {
		if text[i] < ' ' || text[i] > '~' {
			return false
		}
	}
	return true
}
