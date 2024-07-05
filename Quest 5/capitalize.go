package piscine

func isAlphaNum(a rune) bool {
	return (a < '0' || (a > '9' && a < 'A') || (a > 'Z' && a < 'a') || a > 'z')
}

func Capitalize(s string) string {
	runes := []rune(s)
	for i := range runes {
		if runes[0] >= 'a' && runes[0] <= 'z' {
			runes[0] -= 32
		} else if i > 0 && isAlphaNum(runes[i-1]) && runes[i] >= 'a' && runes[i] <= 'z' {
			runes[i] -= 32
		} else if i > 0 && !isAlphaNum(runes[i-1]) && runes[i] >= 'A' && runes[i] <= 'Z' {
			runes[i] += 32
		}
	}
	return string(runes)
}
