package piscine

func NRune(s string, n int) rune {
	text := []rune(s)
	if n-1 < 0 {
		return 0
	} else if n > len(s) {
		return 0
	}
	return text[n-1]
}
