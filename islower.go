package piscine

func IsLower(s string) bool {
	for i := 0; i < len(s); i++ {
		text := []rune(s)
		if text[i] < 'a' || text[i] > 'z' {
			return false
		}
	}
	return true
}
