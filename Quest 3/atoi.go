package piscine

func Atoi(s string) int {
	result := 0
	t := 1
	for i, c := range s {
		if c == '-' && i == 0 {
			t = -1
			continue
		} else if c == '+' && i == 0 {
			continue
		} else if c < '0' || c > '9' {
			return 0
		}
		result = result*10 + int(c-'0')
	}
	return result * t
}
