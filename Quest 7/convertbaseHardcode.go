package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	strFinal := ""
	for _, c := range nbr {
		if (c < '0' || c > '9') && (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') {
			return ""
		}
	}
	if len(baseTo) == 2 {
		if len(baseFrom) == 10 {
			strFinal = FromDec(nbr, baseTo)
		} else if len(baseFrom) == 8 {
			strFinal = OctToBin(nbr, baseFrom)
		} else if len(baseFrom) == 16 {
			strFinal = HexToBin(nbr, baseFrom)
		}
	} else if len(baseTo) == 8 {
		if len(baseFrom) == 2 {
			strFinal = BinToOct(nbr, baseFrom)
		} else if len(baseFrom) == 10 {
			strFinal = FromDec(nbr, baseTo)
		} else if len(baseFrom) == 16 {
			strFinal = HexToOct(nbr, baseFrom)
		}
	} else if len(baseTo) == 10 {
		strFinal = toDecimal(nbr, baseFrom)
	} else if len(baseTo) == 16 {
		if len(baseFrom) == 2 {
			strFinal = BinToHex(nbr, baseFrom)
		} else if len(baseFrom) == 10 {
			strFinal = FromDec(nbr, baseTo)
		} else if len(baseFrom) == 8 {
			strFinal = OctToHex(nbr, baseFrom)
		}
	} else {
		return ""
	}
	return strFinal
}

func FromDec(nbr, baseTo string) string {
	liste := []rune("0123456789abcdef")
	str := ""
	nb := atoi(nbr)
	for nb > 0 {
		str = string(liste[nb%len(baseTo)]) + str
		nb /= len(baseTo)
	}
	return str
}

func atoi(arg string) int {
	nb := 0
	for _, c := range arg {
		nb = nb*10 + int(c-'0')
	}
	return nb
}

func OctToBin(nbr, baseFrom string) string {
	liste := []string{"000", "001", "010", "011", " 100", "101", "110", "111"}
	str := ""
	for i := 0; i < len(nbr); i++ {
		for j := 0; j < len(baseFrom); j++ {
			if nbr[i] == baseFrom[j] {
				str += liste[j]
				break
			}
		}
	}
	if str[0] == '0' {
		str = CleanStr(str)
	}
	return str
}

func HexToBin(nbr, baseFrom string) string {
	liste := []string{"0000", "0001", "0010", "0011", "0100", "0101", "0110", "0111", "1000", "1001", "1010", "1011", "1100", "1101", "1110", "1111"}
	str := ""
	for i := 0; i < len(nbr); i++ {
		for j := 0; j < len(baseFrom); j++ {
			if nbr[i] == baseFrom[j] {
				str += liste[j]
				break
			}
		}
	}
	if str[0] == '0' {
		str = CleanStr(str)
	}
	return str
}

func BinToOct(nbr, baseFrom string) string {
	liste := []string{"000", "001", "010", "011", "100", "101", "110", "111"}
	liste1 := []rune("0123456789abcdef")
	str := ""
	indice := ""
	for i := len(nbr) - 1; i >= 0; i-- {
		indice = string(nbr[i]) + indice
		if len(indice) == 3 || i == 0 {
			if len(indice) == 1 && indice != "0" {
				indice = "00" + indice
			} else if len(indice) == 2 && indice != "00" {
				indice = "0" + indice
			}
			for j, c := range liste {
				if c == indice {
					str = string(liste1[j]) + str
				}
			}
			indice = ""
		}
	}
	if str[0] == '0' {
		str = CleanStr(str)
	}
	return str
}

func HexToOct(nbr, baseFrom string) string {
	str := HexToBin(nbr, baseFrom)
	str = BinToOct(str, "01")
	return str
}

func toDecimal(nbr, baseFrom string) string {
	n := 0
	nb := len(nbr) - 1
	for i := 0; i < len(nbr); i++ {
		for j := 0; j < len(baseFrom); j++ {
			if nbr[i] == baseFrom[j] {
				n += j * Puissance(len(baseFrom), nb)
				nb--
				break
			}
		}
	}
	str := itoa(n)
	return str
}

func Puissance(base, nb int) int {
	n := 1
	if nb == 0 {
		return n
	}
	for nb > 0 {
		n *= base
		nb--
	}
	return n
}

func itoa(n int) string {
	res := ""
	if n == 0 {
		res += "0"
	}
	for n != 0 {
		res = string(n%10+'0') + res
		n /= 10
	}
	return res
}

func BinToHex(nbr, baseFrom string) string {
	liste := []string{"0000", "0001", "0010", "0011", "0100", "0101", "0110", "0111", "1000", "1001", "1010", "1011", "1100", "1101", "1110", "1111"}
	liste1 := []rune("0123456789abcdef")
	str := ""
	indice := ""
	for i := len(nbr) - 1; i >= 0; i-- {
		indice = string(nbr[i]) + indice
		if len(indice) == 4 || i == 0 {
			if len(indice) == 1 && indice != "0" {
				indice = "000" + indice
			} else if len(indice) == 2 && indice != "00" {
				indice = "00" + indice
			} else if len(indice) == 3 && indice != "000" {
				indice = "0" + indice
			}
			for j, c := range liste {
				if c == indice {
					str = string(liste1[j]) + str
				}
			}
			indice = ""
		}
	}
	if str[0] == '0' {
		str = CleanStr(str)
	}
	return str
}

func OctToHex(nbr, baseFrom string) string {
	str := OctToBin(nbr, baseFrom)
	str = BinToHex(str, "01")
	return str
}

func CleanStr(str string) string {
	for i, c := range str {
		if c != '0' {
			str = str[i:]
			break
		}
	}
	return str
}
