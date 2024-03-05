package piscine

func TrimAtoi(s string) int {
	result := 0
	t := 1
	for _, c := range s {
		if c >= '0' && c <= '9' {
			t = 1
			break
		} else if c == '-' {
			t = -1
			break
		}
	}
	for _, c := range s {
		if c >= '0' && c <= '9' {
			result = result*10 + int(c-'0')
		}
	}
	return result * t
}
