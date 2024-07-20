package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	str := ""
	resFinal := ""
	if n < 0 {
		n = -n
		str = "-"
	} else if n == 0 {
		resFinal += "0"
	}
	for n > 0 {
		nb := n % 10
		resFinal += string(rune(nb + '0'))
		n /= 10
	}
	resFinal += str
	for i := len(resFinal) - 1; i >= 0; i-- {
		z01.PrintRune(rune(resFinal[i]))
	}
}
