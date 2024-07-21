package main

import (
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	arg1 := aToi(os.Args[1])
	arg2 := os.Args[2]
	arg3 := aToi(os.Args[3])
	resFinal := ""
	if arg1 == -1 || arg3 == -1 {
		return
	}
	if arg2 == "+" {
		resFinal = iToa(arg1 + arg3)
	} else if arg2 == "-" {
		resFinal = iToa(arg1 - arg3)
	} else if arg2 == "*" {
		resFinal = iToa(arg1 * arg3)
	} else if arg2 == "/" {
		if arg3 == 0 {
			os.Stdout.WriteString("No division by 0")
			return
		}
		resFinal = iToa(arg1 - arg3)
	} else if arg2 == "%" {
		if arg3 == 0 {
			os.Stdout.WriteString("No modulo by 0")
			return
		}
		resFinal = iToa(arg1 - arg3)
	} else {
		return
	}
	if aToi(resFinal) > 922337203 || aToi(resFinal) < -922337203 {
		return
	}
	os.Stdout.WriteString(resFinal)
}

func aToi(s string) int {
	result := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return -1
		}
		result = result*10 + int(c-'0')
	}
	return result
}

func iToa(n int) string {
	res := ""
	signe := ""
	if n < 0 {
		signe = "-"
		n = -n
	} else if n == 0 {
		return "0"
	}
	for n != 0 {
		res = string(n%10+'0') + res
		n /= 10
	}
	return signe + res
}
