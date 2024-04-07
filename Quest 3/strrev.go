package piscine

func StrRev(s string) string {
	text := ""
	for i := len(s)-1; i >= 0; i-- {
		text += string(s[i])
	}
	return text
}
