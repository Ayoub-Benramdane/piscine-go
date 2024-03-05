package piscine

func Capitalize(s string) string {
	result := make([]byte, len(s))

	isFirst := true

	for i := 0; i < len(s); i++ {
		char := s[i]
		if isFirst && ((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')) {
			if char >= 'a' && char <= 'z' {
				result[i] = char - 'a' + 'A'
			} else {
				result[i] = char
			}
			isFirst = false
		} else {
			if char >= 'A' && char <= 'Z' {
				result[i] = char - 'A' + 'a'
			} else {
				result[i] = char
			}
		}

		if char < '0' || (char > '9' && char < 'A') || (char > 'Z' && char < 'a') || char > 'z' {
			result[i] = char
			isFirst = true
		}
	}
	return string(result)
}
