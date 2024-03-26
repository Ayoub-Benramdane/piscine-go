package piscine

import "github.com/01-edu/z01"

func IsNegative(nb int) {
	if nb < 0 {
		z01.PrintRune('t')
	} else {
		z01.PrintRune('f')
	}
	z01.PrintRune('\n')
}
