package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	alpha := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	alpha1 := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	str := ""
	if len(args) == 0 {
		return
	}
	for i := 0; i < len(args); i++ {
		if args[i] == "--upper" && i == 0 {
			alpha = alpha1
			continue
		}
		arg := atoi(args[i])
		if arg < 0 || arg > len(alpha) {
			str += " "
			continue
		}
		str += string(alpha[arg])
	}
	for _, c := range str {
		z01.PrintRune(c)
	}
}

func atoi(arg string) int {
	nb := 0
	for _, c := range arg {
		nb = nb*10 + int(c-'0')
	}
	return nb - 1
}
