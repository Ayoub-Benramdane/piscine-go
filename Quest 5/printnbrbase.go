package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	str := ""
	strFinal := ""
	if len(base) < 2 || handlingError(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	if nbr < 0 {
		nbr = -nbr
		strFinal = "-"
	}
	for nbr > 0 {
		str = string(base[nbr%len(base)]) + str
		nbr /= len(base)
	}
	strFinal += str
	for _, c := range strFinal {
		z01.PrintRune(c)
	}
}

func handlingError(base string) bool {
	for i := 0; i < len(base); i++ {
		for j := i + 1; j < len(base); j++ {
			if base[j] == base[i] || base[i] == '-' || base[i] == '+' || base[j] == '-' || base[j] == '+' {
				return true
			}
		}
	}
	return false
}
