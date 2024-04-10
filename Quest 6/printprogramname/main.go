package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[0][2:]
	for _, v := range arg {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}