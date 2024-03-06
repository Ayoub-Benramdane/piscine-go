package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	str1 := os.Args[1]
	str2 := os.Args[2]

	table := make(map[rune]bool)
	result := ""

	for _, char := range str1 {
		if !table[char] {
			table[char] = true
			result += string(char)
		}
	}

	for _, char := range str2 {
		if table[char] {
			fmt.Print(string(char))
			table[char] = false
		}
	}
	fmt.Println()
}
