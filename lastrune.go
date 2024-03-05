package piscine

func LastRune(s string) rune {
	text := []rune(s)
	return text[len(s)-1]
}
