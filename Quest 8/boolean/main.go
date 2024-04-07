package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	args := os.Args[1:]
	nbr := len(args)
	if isEven(nbr){
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}
