package piscine

func IsUpper(s string) bool {
	for i := 0; i < len(s); i++ {
		text := []rune(s)
		if text[i] < 'A' || text[i] > 'Z' {
			return false
		}
	}
	return true
}
