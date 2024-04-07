package piscine

func Rot14(s string) string {
	text := []rune(s)
	for i, c := range s {
		if c >= 'a' && c <= 'l' || c >= 'A' && c <= 'L' {
			text[i] += 14
		} else if c >= 'm' && c <= 'z' || c >= 'M' && c <= 'Z' {
			text[i] -= 12
		}
	}
	return string(text)
}
