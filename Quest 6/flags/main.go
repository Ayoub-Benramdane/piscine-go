package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 || args[0] == "-h" || args[0] == "--help" {
		help()
		return
	}
	str := ""
	strInst := ""
	order := false
	for _, c := range args {
		if len(c) > 1 && c[:1] == "-" {
			if c == "--order" || c == "-o" {
				order = true
			} else if len(c) > 9 && c[:9] == "--insert=" {
				strInst += c[9:]
			} else if len(c) > 3 && c[:3] == "-i=" {
				strInst += c[3:]
			} else {
				help()
				return
			}
			continue
		}
		str += c
	}
	str += strInst
	if order {
		orderString(str)
	} else {
		fmt.Println(str)
	}
}

func orderString(str string) {
	runeStr := []rune(str)
	for i := 0; i < len(runeStr); i++ {
		if i < len(runeStr)-1 && runeStr[i] > runeStr[i+1] {
			runeStr[i], runeStr[i+1] = runeStr[i+1], runeStr[i]
			i = -1
		}
	}
	fmt.Println(string(runeStr))
}

func help() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("         This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("         This flag will behave like a boolean, if it is called it will order the argument.")
}
