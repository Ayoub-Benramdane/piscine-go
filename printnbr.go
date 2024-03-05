package piscine

import "github.com/01-edu/z01"

func check(x int) {
	y := '0'
	if x == 0 {
		z01.PrintRune(y)
		return
	}
	for i := 1; i <= (x % 10); i++ {
		y++
	}
	if x/10 != 0 {
		check(x / 10)
	}
	z01.PrintRune(y)
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	check(n)
}
