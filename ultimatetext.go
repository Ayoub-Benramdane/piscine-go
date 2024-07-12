package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	x := os.Args[1:]
	for i :=0;i<len(x);i++{
		ultimateprint(x[i])
	}
}
func ultimateprint(s string) {
	var p []rune
    for i := 33; i <= 126; i++ {
        p = append(p, rune(i))
    }
	str := []rune(s)
	for i := 0; i < len(str); i++ {
		for j := p[0]; j <= str[i]; j++ {
			fmt.Printf("%c", j)
			time.Sleep(10 * time.Millisecond)
			if j != str[i] {
				clearScreen()
			}
		}
	}
	fmt.Printf(" ")
}
func clearScreen() {
	fmt.Print("\b \b")
}
