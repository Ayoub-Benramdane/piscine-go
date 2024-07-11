package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	for i := 0; i < len(arg); i++ {
		if i < len(arg)-1 && arg[i+1] < arg[i] {
			arg[i], arg[i+1] = arg[i+1], arg[i]
			i = 0
		}
	}
	for _, v := range arg {
		for _, c := range v {
			z01.PrintRune(c)
			z01.PrintRune('\n')
		}
	}
}
