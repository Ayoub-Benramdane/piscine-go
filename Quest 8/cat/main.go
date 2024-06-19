package main

import (
	"github.com/01-edu/z01"
	"io"
	"os"
)

func printS(str string) {
	for _, c := range str {
		z01.PrintRune(c)
	}
}
func main() {
	if len(os.Args) == 1 {
		io.Copy(os.Stdout, os.Stdin)
	} else {
		args := os.Args[1:]
		for _, c := range args {
			file, err := os.Open(c)
			if err != nil {
				printS("ERROR: ")
				printS(err.Error())
				z01.PrintRune('\n')
				os.Exit(1)
			}
			arr, _ := io.ReadAll(file)
			printS(string(arr))
			file.Close()
		}
	}
}
