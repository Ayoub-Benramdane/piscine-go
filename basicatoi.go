package piscine

func BasicAtoi(s string) int {
	result := 0
	for _, c := range s {
		result = result*10 + int(c-'0')
	}
	return result
}
