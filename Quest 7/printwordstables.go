package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, c := range a {
		for _, d := range c {
			z01.PrintRune(d)
		}
		z01.PrintRune('\n')
	}
}
