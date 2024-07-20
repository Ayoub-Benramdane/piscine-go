package piscine

func AtoiBase(s string, base string) int {
	n := 0
	nb := 0
	if len(base) < 2 || handlingError(base) {
		return 0
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := 0; j < len(base); j++ {
			if s[i] == base[j] {
				n += j * Puissance(len(base), nb)
				nb++
				break
			} else if j == len(base)-1 && base[j] != s[i] {
				return 0
			}
		}
	}
	return n
}

func handlingError(base string) bool {
	for i := 0; i < len(base); i++ {
		if base[i] == '-' || base[i] == '+' {
			return true
		}
		for j := i + 1; j < len(base); j++ {
			if base[j] == base[i] {
				return true
			}
		}
	}
	return false
}

func Puissance(base, nb int) int {
	n := 1
	for nb > 0 {
		n *= base
		nb--
	}
	return n
}
