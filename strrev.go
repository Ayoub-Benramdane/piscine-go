package piscine

func StrRev(s string) string {
	text := []rune(s)

	start := 0
	end := len(text) - 1

	for start < end {
		text[start], text[end] = text[end], text[start]
		start++
		end--
	}

	return string(text)
}
