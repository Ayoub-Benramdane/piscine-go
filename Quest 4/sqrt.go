package main

import (
	"fmt"
)

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(3))
}

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}
	for result := 0; result*result <= nb; result++ {
		if result*result == nb {
			return result
		}
	}
	return 0
}
