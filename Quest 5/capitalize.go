package piscine

func Capitalize(s string) string {
	result := ""
	isFirst := true
	for i := 0; i < len(s); i++ {
		if isFirst {
			if s[i] >= 'a' && s[i] <= 'z' {
				result += string(s[i] - 32)
			} else {
				result += string(s[i])
			}
			isFirst = false
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			result += string(s[i] + 32)
		} else {
			result += string(s[i])
		}
		if s[i] < '0' || (s[i] > '9' && s[i] < 'A') || (s[i] > 'Z' && s[i] < 'a') || s[i] > 'z' {
			isFirst = true
		}
	}
	return result
}
