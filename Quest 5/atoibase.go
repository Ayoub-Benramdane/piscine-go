package piscine

func AtoiBase(s string, base string) int {
	n := 0
	nb := 0
	for _, c := range base {
		if (c < '0' || c > '9') && (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') {
			return n
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := 0; j < len(base); j++ {
			if s[i] == base[j] {
				n += j * Puissance(len(base), nb)
				nb++
			}
		}
	}
	return n
}

func Puissance(base, nb int) int {
	n := 1
	if nb == 0 {
		return n
	}
	for nb > 0 {
		n *= base
		nb--
	}
	return n
}
