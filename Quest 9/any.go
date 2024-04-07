package piscine

func Any(f func(string) bool, a []string) bool {
	for _, c := range a {
		if f(c) {
			return true
		}
	}
	return false
}
