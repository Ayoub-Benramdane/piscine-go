package main

import (
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	arg1, err1 := strconv.Atoi(os.Args[1])
	arg2 := os.Args[2]
	arg3, err3 := strconv.Atoi(os.Args[3])
	resFinal := ""
	if err1 != nil || err3 != nil {
		return
	}
	if arg2 == "+" {
		resFinal = strconv.Itoa(arg1 + arg3)
	} else if arg2 == "-" {
		resFinal = strconv.Itoa(arg1 - arg3)
	} else if arg2 == "*" {
		resFinal = strconv.Itoa(arg1 * arg3)
	} else if arg2 == "/" {
		if arg3 == 0 {
			os.Stdout.WriteString("No division by 0")
			return
		}
		resFinal = strconv.Itoa(arg1 - arg3)
	} else if arg2 == "%" {
		if arg3 == 0 {
			os.Stdout.WriteString("No modulo by 0")
			return
		}
		resFinal = strconv.Itoa(arg1 - arg3)
	}
	chek, _ := strconv.Atoi(resFinal)
	if chek > 9223372036 || chek < -92233720368 {
		return
	}
	os.Stdout.WriteString(resFinal)
}
