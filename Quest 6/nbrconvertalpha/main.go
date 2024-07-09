package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	character := 96
	str := ""
	if len(args) == 0 {
		return
	}
	for i := 0; i < len(args); i++ {
		if i == 0 && args[i] == "--upper" {
			character -= 32
			continue
		}
		num := atoi(args[i])
		if num < 1 || num > 26 {
			str += " "
		} else {
			str += string(character + num)
		}
	}
	for _, c := range str {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func atoi(arg string) int {
	nb := 0
	for _, c := range arg {
		nb = nb*10 + int(c-'0')
	}
	return nb
}
