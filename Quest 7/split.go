package piscine

func Split(s, sep string) []string {
	str := ""
	slice := []string{}
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			str += string(s[i])
			slice = append(slice, str)
		}
		if i+len(sep) < len(s) && s[i:i+len(sep)] == sep {
			slice = append(slice, str)
			str = ""
			i += len(sep) - 1
			continue
		}
		str += string(s[i])
	}
	return slice
}
