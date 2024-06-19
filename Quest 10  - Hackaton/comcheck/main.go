package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i < len(os.Args); i++ {
		if os.Args[i] == "01" || os.Args[i] == "galaxy" || os.Args[i] == "galaxy 01" {
			fmt.Println("Alert!!!")
		}
	}
}
