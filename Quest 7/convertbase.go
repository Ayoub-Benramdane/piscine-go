package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	s := ""
	n := 0
	in := 0
	res := ""
	for _, v := range nbr {
		for i, x := range baseFrom {
			if v == x {
				in = i
			}
		}
		n = n*len(baseFrom) + in
	}
	for n > 0 {
		s += string(baseTo[n%len(baseTo)])
		n /= len(baseTo)
	}
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}
	return res
}
