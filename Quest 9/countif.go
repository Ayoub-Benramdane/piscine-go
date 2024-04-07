package piscine

func CountIf(f func(string) bool, tab []string) int {
	n := 0
	for _, c := range tab {
		if f(c) == true {
			n++
		}
	}
	return n
}
