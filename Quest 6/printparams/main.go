package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	for i := 0; i < len(arg); i++ {
		for _, v := range arg[i] {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}
