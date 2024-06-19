package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println()
		return
	}
	str := ""
	for i := 0; i < len(args); i++ {
		if i == len(args)-1 {
			str += args[i]
			break
		}
		str += args[i] + " "
	}
	runeStr := []rune(str)
	vowel := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	depart := len(runeStr) - 1
	for i := 0; i < len(runeStr); i++ {
		for j := 0; j < len(vowel); j++ {
			if vowel[j] == runeStr[i] && i < depart {
				depart = mirror(runeStr, vowel, i, depart)
				break
			}
		}
	}
	newStr := string(runeStr)
	fmt.Println(newStr)
}

func mirror(runeStr, vowel []rune, indice, depart int) int {
	for i := depart; i >= 0; i-- {
		depart--
		for j := 0; j < len(vowel); j++ {
			if vowel[j] == runeStr[i] {
				runeStr[i], runeStr[indice] = runeStr[indice], runeStr[i]
				return depart
			}
		}
	}
	return depart
}
