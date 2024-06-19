package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	for i := 0; i < len(base); i++ {
		for j := i + 1; j < len(base); j++ {
			if base[j] == base[i] {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}
	str := ""
	nb := nbr
	if nbr < 0 {
		nb = -nbr
	}
	for nb > 0 {
		n := nb % len(base)
		str = string(base[n]) + str
		nb /= len(base)
	}
	if nbr < 0 {
		str = "-" + str
	}
	for _, c := range str {
		z01.PrintRune(c)
	}
}
