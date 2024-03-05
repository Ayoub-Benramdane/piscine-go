package piscine

func AlphaCount(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		text := []rune(s)
		if (text[i] >= 'A' && text[i] <= 'Z') || (text[i] >= 'a' && text[i] <= 'z') {
			count++
		}
	}
	return count
}
