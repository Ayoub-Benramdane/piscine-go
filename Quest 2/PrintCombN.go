package main

import (
	"github.com/01-edu/z01"
)

func main() {
	PrintCombN(1)
	PrintCombN(3)
	PrintCombN(9)
}

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	// Helper function to generate combinations
	var generateComb func(current string, start int)
	combinations := []string{}

	generateComb = func(current string, start int) {
		if len(current) == n {
			combinations = append(combinations, current)
			return
		}
		for i := start; i <= 9; i++ {
			generateComb(current+string(rune('0'+i)), i+1)
		}
	}

	generateComb("", 0)

	// Print the combinations using z01.PrintRune
	for i, comb := range combinations {
		for _, char := range comb {
			z01.PrintRune(char)
		}
		if i < len(combinations)-1 {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
