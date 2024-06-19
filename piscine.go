package main

//---------  Quest II  ---------//

/* Q2-1
func main() {
	for i := 'a'; i <= 'z'; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
*/

/* Q2-2
func main() {
	for i := 'z'; i >= 'a'; i-- {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
*/

/* Q2-3
func main() {
	for i := '0'; i <= '9'; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
*/

/* Q2-4
func main() {
	IsNegative(1)
	IsNegative(0)
	IsNegative(-1)
}

func IsNegative(nb int) {
	if nb < 0 {
		z01.PrintRune('T')
	} else  {
		z01.PrintRune('F')
	}
	z01.PrintRune('\n')
}
*/

/* Q2-5
func main() {
	for i := '0'; i <= '7'; i++ {
		for j := i+1; j <= '8'; j++ {
			for k := j+1; k <= '9'; k++ {
				if i < '7' {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
					z01.PrintRune(',')
					z01.PrintRune(' ')
				} else {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
				}
			}
		}
	}
}
*/

/* Q2-6
func main() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := '0'; k <= '9'; k++ {
				for l := '0'; l <= '9'; l++ {
					if i <= k && j < l {
						if i == '9' && j == '8' && k == '9' && l == '9' {
							z01.PrintRune(i)
							z01.PrintRune(j)
							z01.PrintRune(' ')
							z01.PrintRune(k)
							z01.PrintRune(l)
						} else {
							z01.PrintRune(i)
							z01.PrintRune(j)
							z01.PrintRune(' ')
							z01.PrintRune(k)
							z01.PrintRune(l)
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
}
*/

/* Q2-7
func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(1323)
	z01.PrintRune('\n')
}

func PrintNbr(n int) {
	res := []rune{}
	val := '0'
	if n == 0 {
		z01.PrintRune('0')
	}
	if n < 0 {
		n = -n
		z01.PrintRune('-')
	}
	for n%10 != 0 {
		for i := 1; i <= n%10; i++ {
			val++
		}
		res = append(res, val)
		val = '0'
		n /= 10
	}
	for i := len(res)-1; i >= 0 ; i-- {
		z01.PrintRune(res[i])
	}
}
*/

/* Q2-8
func main() {
	PrintCombN(1)
	PrintCombN(3)
	PrintCombN(9)
}

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}
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
*/

//---------  Quest III  ---------//

/* Q3-1
func main() {
	n := 0
	PointOne(&n)
	fmt.Println(n)
}

func PointOne(n *int) {
	*n = 1
}
*/

/* Q3-2
func main() {
	a := 0
	b := &a
	n := &b
	UltimatePointOne(&n)
	fmt.Println(a)
}

func UltimatePointOne(n ***int) {
	***n = 1
}
*/

/* Q3-3
func main() {
	a := 13
	b := 2
	var div int
	var mod int
	DivMod(a, b, &div, &mod)
	fmt.Println(div)
	fmt.Println(mod)
}

func DivMod(a int, b int, div *int, mod *int) {
	*div = a/b
	*mod = a%b
}
*/

/* Q3-4
func main() {
	a := 13
	b := 2
	UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}

func UltimateDivMod(a *int, b *int) {
	div := *a / *b
	mod := *a % *b
	*a = div
	*b = mod
}
*/

/* Q3-5
func main() {
	PrintStr("Hello World!")
}

func PrintStr(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}
*/

/* Q3-6
func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
}

func StrLen(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}
*/

/* Q3-7
func main() {
	a := 0
	b := 1
	Swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}

func Swap(a *int, b *int) {
	*a, *b = *b, *a
}
*/

/* Q3-8
func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}

func StrRev(s string) string {
	res := ""
	for i := len(s)-1; i >= 0; i-- {
		res += string(s[i])
	}
	return res
}
*/

/* Q3-9
func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}

func BasicAtoi(s string) int {
	res := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			res = res * 10 + int(c - '0')
		}
	}
	return res
}
*/

/* Q3-10
func main() {
	fmt.Println(BasicAtoi2("12345"))
	fmt.Println(BasicAtoi2("0000000012345"))
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2("Hello World!"))
}

func BasicAtoi2(s string) int {
	res := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			res = res * 10 + int(c - '0')
		} else {
			res = 0
			break
		}
	}
	return res
}
*/

/* Q3-11
func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}

func Atoi(s string) int {
	res := 0
	signe := 1
	for i, c := range s {
		if c == '+' && i == 0 {
			continue
		} else if c == '-' && i == 0 {
			signe = -1
			continue
		} else if c >= '0' && c <= '9' {
			res = res * 10 + int(c - '0')
		} else {
			res = 0
			break
		}
	}
	return res * signe
}
*/

/* Q3-12
func main() {
	s := []int{5, 4, 3, 2, 1, 0}
	SortIntegerTable(s)
	fmt.Println(s)
}

func SortIntegerTable(table []int) {
	for i := 0; i < len(table); i++ {
		for j := i; j < len(table); j++ {
			if table[j] < table[i] {
				table[j], table[i] = table[i], table[j]
			}
		}
	}
}
*/

//---------  Quest IV  ---------//

/* Q4-1
func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}

func IterativeFactorial(nb int) int {
	res := 1
	if nb < 0 {
		return 0
	}
	for i := nb; i > 1; i-- {
		res *= i
	}
	if res < 0 {
		return 0
	}
	return res
}
*/

/* Q4-2
func main() {
	arg := 22
	fmt.Println(RecursiveFactorial(arg))
}

func RecursiveFactorial(nb int) int {
	res := 1
	if nb < 0 {
		return 0
	}
	if nb > 0 {
		res = nb * RecursiveFactorial(nb - 1)
	}
	if res < 0 {
		return 0
	}
	return res
}
*/

/* Q4-3
func main() {
	fmt.Println(IterativePower(4, 3))
}

func IterativePower(nb int, power int) int {
	res := 1
	if power < 0 {
		return 0
	}
	for i := 0 ; i < power; i++ {
		res *= nb
	}
	return res
}
*/

/* Q4-4
func main() {
	fmt.Println(RecursivePower(4, 3))
}

func RecursivePower(nb int, power int) int {
	res := 1
	if power < 0 {
		return 0
	}
	if power > 0 {
		res = nb * RecursivePower(nb, power- 1)
	}
	return res
}
*/

/* Q4-5
func main() {
	arg1 := 4
	fmt.Println(Fibonacci(arg1))
}

func Fibonacci(index int) int {
	res := 1
	if index == 0 {
		return 0
	}
	if index < 0 {
		return -1
	}
	if index > 1 {
		res = Fibonacci(index-2) + Fibonacci(index-1)
	}
	return res
}
*/

/* Q4-6
func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(3))
}

func Sqrt(nb int) int {
	res := 0
	for i := 1; i <= nb; i++ {
		if i*i == nb {
			res = i
		}
	}
	return res
}
*/

/* Q4-7
func main() {
	fmt.Println(IsPrime(5))
	fmt.Println(IsPrime(4))
}

func IsPrime(nb int) bool {
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
*/

/* Q4-8
func main() {
	fmt.Println(FindNextPrime(5))
	fmt.Println(FindNextPrime(4))
}

func FindNextPrime(nb int) int {
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return FindNextPrime(nb + 1)
		}
	}
	return nb
}
*/

//---------  Quest V  ---------//

/* Q5-1
func main() {
	z01.PrintRune(FirstRune("Hello!"))
	z01.PrintRune(FirstRune("Salut!"))
	z01.PrintRune(FirstRune("Ola!"))
	z01.PrintRune('\n')
}

func FirstRune(s string) rune {
	return rune(s[0])
}
*/

/* Q5-2
func main() {
	z01.PrintRune(LastRune("Hello!"))
	z01.PrintRune(LastRune("Salut!"))
	z01.PrintRune(LastRune("Ola!"))
	z01.PrintRune('\n')
}

func LastRune(s string) rune {
	return rune(s[len(s)-1])
}
*/

/* Q5-3
func main() {
	z01.PrintRune(NRune("Hello!", 3))
	z01.PrintRune(NRune("Salut!", 2))
	z01.PrintRune(NRune("Bye!", -1))
	z01.PrintRune(NRune("Bye!", 5))
	z01.PrintRune(NRune("Ola!", 4))
	z01.PrintRune('\n')
}

func NRune(s string, n int) rune {
	if n-1 < 0 || n > len(s) {
		return 0
	}
	return rune(s[n-1])
}
*/

/* Q5-4
func main() {
	fmt.Println(Compare("Hello!", "Hello!"))
	fmt.Println(Compare("Salut!", "lut!"))
	fmt.Println(Compare("Ola!", "Ol"))
}

func Compare(a, b string) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}
*/

/* Q5-5
func main() {
	s := "Hello 78 World!    4455 /"
	nb := AlphaCount(s)
	fmt.Println(nb)
}

func AlphaCount(s string) int {
	count := 0
	for _, c := range s {
		if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
			count++
		}
	}
	return count
}
*/

/* Q5-6
func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
}

func Index(s string, toFind string) int {
	for i := 0; i < len(s)-len(toFind); i++ {
		if s[i:i+len(toFind)]==toFind {
			return i
		}
	}
	return -1
}
*/

/* Q5-7
func main() {
	fmt.Println(Concat("Hello!", " How are you?"))
}

func Concat(str1 string, str2 string) string {
	return str1 + str2
}
*/

/* Q5-8
func main() {
	fmt.Println(IsUpper("HELLO"))
	fmt.Println(IsUpper("HELLO!"))

}

func IsUpper(s string) bool {
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			continue
		} else {
			return false
		}
	}
	return true
}
*/

/* Q5-9
func main() {
	fmt.Println(IsLower("hello"))
	fmt.Println(IsLower("hello!"))

}

func IsLower(s string) bool {
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			continue
		} else {
			return false
		}
	}
	return true
}
*/

/* Q5-10
func main() {
	fmt.Println(IsAlpha("Hello! How are you?"))
	fmt.Println(IsAlpha("HelloHowareyou"))
	fmt.Println(IsAlpha("What's this 4?"))
	fmt.Println(IsAlpha("Whatsthis4"))

}

func IsAlpha(s string) bool {
	for _, c := range s {
		if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}
*/

/* Q5-11
func main() {
	fmt.Println(IsNumeric("010203"))
	fmt.Println(IsNumeric("01,02,03"))
}

func IsNumeric(s string) bool {
	for _, c := range s {
		if c >= '0' && c <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}
*/

/* Q5-12
func main() {
	fmt.Println(IsPrintable("Hello"))
	fmt.Println(IsPrintable("Hello\n"))

}

func IsPrintable(s string) bool {
	for _, c := range s {
		if c >= ' ' && c <= '~' {
			continue
		} else {
			return false
		}
	}
	return true
}
*/

/* Q5-13
func main() {
	fmt.Println(ToUpper("Hello! How are you?"))
}

func ToUpper(s string) string {
	res := ""
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			res += string(c - 32)
		} else {
			res += string(c)
		}
	}
	return res
}
*/

/* Q5-14
func main() {
	fmt.Println(ToLower("Hello! How are you?"))
}

func ToLower(s string) string {
	res := ""
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			res += string(c + 32)
		} else {
			res += string(c)
		}
	}
	return res
}
*/

/* Q5-15
func main() {
	PrintNbrInOrder(321)
	PrintNbrInOrder(0)
	PrintNbrInOrder(321)
}

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	res := [10]int{}
	for n != 0 {
		res[n%10]++
		n /= 10
	}
	for i := 0; i < len(res); i++ {
		for res[i] > 0 {
			z01.PrintRune(rune(i) + '0')
			res[i]--
		}
	}
}
*/

/* Q5-16
func main() {
	fmt.Println(TrimAtoi("12345"))
	fmt.Println(TrimAtoi("str123ing45"))
	fmt.Println(TrimAtoi("012 345"))
	fmt.Println(TrimAtoi("Hello World!"))
	fmt.Println(TrimAtoi("sd+x1fa2W3s4"))
	fmt.Println(TrimAtoi("sd-x1fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1-fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1+fa2W3s4"))
}

func TrimAtoi(s string) int {
	res := 0
	signe := 1
	for _, c := range s {
		if c >= '0' && c <= '9' {
			break
		} else if c == '-' {
			signe = -1
			break
		}
	}
	for _, c := range s {
		if c >= '0' && c <= '9' {
			res = res*10 + int(c-'0')
		}
	}
	return res * signe
}
*/

/* Q5-17
func main() {
	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
}

func Capitalize(s string) string {
	res := ""
	isFirst := true
	for _, c := range s {
		if c >= 'a' && c <= 'z' && isFirst {
			res += string(c - 32)
			isFirst = false
		} else if c >= 'A' && c <= 'Z' && !isFirst {
			res += string(c + 32)
			isFirst = false
		} else {
			res += string(c)
			isFirst = false
		}
		if !(c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9') {
			isFirst = true
		}
	}
	return res
}
*/

/* Q5-18
func main() {
	elems := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(BasicJoin(elems))
}

func BasicJoin(elems []string) string {
	res := ""
	for i := 0; i < len(elems); i++ {
		res += elems[i]
	}
	return res
}
*/

/* Q5-19
func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(Join(toConcat, ":"))
}

func Join(strs []string, sep string) string {
	res := ""
	for i := 0; i < len(strs); i++ {
		if i != len(strs)-1 {
			res += strs[i] + sep
		} else {
			res += strs[i]
		}
	}
	return res
}
*/

/* Q5-20
func main() {
	PrintNbrBase(125, "0123456789")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "01")
	z01.PrintRune('\n')
	PrintNbrBase(125, "0123456789ABCDEF")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "choumi")
	z01.PrintRune('\n')
	PrintNbrBase(125, "aa")
	z01.PrintRune('\n')
}

func PrintNbrBase(nbr int, base string) {
	for i := 0; i < len(base); i++ {
		for j := i + 1; j < len(base); j++ {
			if base[j] == base[i] {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}
	str := ""
	nb := nbr
	if nbr < 0 {
		nb = -nbr
	}
	for nb > 0 {
		n := nb % len(base)
		str = string(base[n]) + str
		nb /= len(base)
	}
	if nbr < 0 {
		str = "-" + str
	}
	for _, c := range str {
		z01.PrintRune(c)
	}
}
*/

/* Q5-21
func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}

func AtoiBase(s string, base string) int {
	n := 0
	nb := 0
	for _, c := range base {
		if (c < '0' || c > '9') && (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') {
			return n
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := 0; j < len(base); j++ {
			if s[i] == base[j] {
				n += j * Puissance(len(base), nb)
				nb++
			}
		}
	}
	return n
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
*/

//---------  Quest VI  ---------//

/* Q6-1
func main() {
	arg := os.Args[0][2:]
	for _, v := range arg {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
*/

/* Q6-2
func main() {
	arg := os.Args
	for i := 1; i < len(arg); i++ {
		for _, v := range arg[i] {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}
*/

/* Q6-3
func main() {
	arg := os.Args
	for i := len(arg) - 1; i >= 1; i-- {
		for _, v := range arg[i] {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}
*/

/* Q6-4
func main() {
	arg := os.Args
	for i := 1; i < len(arg); i++ {
		for j := i + 1; j < len(arg); j++ {
			if arg[i] > arg[j] {
				arg[i], arg[j] = arg[j], arg[i]
				i = 0
				break
			}
		}
		for _, v := range arg[i] {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}
*/

/* Q6-5
func main() {
	args := os.Args[1:]
	alpha := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	alpha1 := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	str := ""
	if len(args) == 0 {
		return
	}
	for i := 0; i < len(args); i++ {
		if args[i] == "--upper" && i == 0 {
			alpha = alpha1
			continue
		}
		arg := atoi(args[i])
		if arg < 0 || arg > len(alpha) {
			str += " "
			continue
		}
		str += string(alpha[arg])
	}
	for _, c := range str {
		z01.PrintRune(c)
	}
}

func atoi(arg string) int {
	nb := 0
	for _, c := range arg {
		nb = nb*10 + int(c-'0')
	}
	return nb - 1
}
*/

/* Q6-6
func main() {
	args := os.Args[1:]
	if len(args) == 0 || args[0] == "-h" || args[0] == "--help" {
		help()
		return
	}
	str := ""
	strInst := ""
	order := false
	for _, c := range args {
		if len(c) > 1 && c[:1] == "-" {
			if c == "--order" || c == "-o" {
				order = true
			} else if len(c) > 9 && c[:9] == "--insert=" {
				strInst += c[9:]
			} else if len(c) > 3 && c[:3] == "-i=" {
				strInst += c[3:]
			} else {
				help()
				return
			}
			continue
		}
		str += c
	}
	str += strInst
	if order {
		orderString(str)
	} else {
		fmt.Println(str)
	}
}

func orderString(str string) {
	runeStr := []rune(str)
	for i := 0; i < len(runeStr); i++ {
		for j := i + 1; j < len(runeStr); j++ {
			if runeStr[i] > runeStr[j] {
				runeStr[i], runeStr[j] = runeStr[j], runeStr[i]
			}
		}
	}
	newStr := string(runeStr)
	fmt.Println(newStr)
}

func help() {
	fmt.Println("--insert")
	fmt.Println("-i")
	fmt.Println("This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("-o")
	fmt.Println("This flag will behave like a boolean, if it is called it will order the argument.")
}
*/

/* Q6-7
func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println()
		return
	}
	str := ""
	for i := 0; i < len(args); i++ {
		if i == len(args)-1 {
			str += args[i]
			break
		}
		str += args[i] + " "
	}
	runeStr := []rune(str)
	vowel := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	depart := len(runeStr) - 1
	for i := 0; i < len(runeStr); i++ {
		for j := 0; j < len(vowel); j++ {
			if vowel[j] == runeStr[i] && i < depart {
				depart = mirror(runeStr, vowel, i, depart)
				break
			}
		}
	}
	newStr := string(runeStr)
	fmt.Println(newStr)
}

func mirror(runeStr, vowel []rune, indice, depart int) int {
	for i := depart; i >= 0; i-- {
		depart--
		for j := 0; j < len(vowel); j++ {
			if vowel[j] == runeStr[i] {
				runeStr[i], runeStr[indice] = runeStr[indice], runeStr[i]
				return depart
			}
		}
	}
	return depart
}
*/

//---------  Quest VII  ---------//

/* Q7-1
func main() {
	fmt.Println(AppendRange(5, 10))
	fmt.Println(AppendRange(10, 5))
}

func AppendRange(min, max int) []int {
	res := []int{}
	for i := min; i < max; i++ {
		res = append(res, i)
	}
	return res
}
*/

/* Q7-2
func main() {
	fmt.Println(MakeRange(5, 10))
	fmt.Println(MakeRange(10, 5))
}

func MakeRange(min, max int) []int {
	if min > max {
		return nil
	}
	res := make([]int, max-min)
	for i := 0; i < max-min; i++ {
		res[i] = min + i
	}
	return res
}
*/

/* Q7-3
func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(ConcatParams(test))
}

func ConcatParams(args []string) string {
	res := ""
	for i := 0; i < len(args); i++ {
		res += args[i]
		if i != len(args) - 1 {
			res += "\n"
		}
	}
	return res
}
*/

/* Q7-4
func main() {
	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))
}

func SplitWhiteSpaces(s string) []string {
	res := []string{}
	word := ""
	for i, c := range s {
		if i == len(s)-1 {
			word += string(c)
			res = append(res, word)
		} else if c != ' ' {
			word += string(c)
		} else {
			res = append(res, word)
			word = ""
		}
	}
	return res
}
*/

/* Q7-5
func main() {
	a := SplitWhiteSpaces("Hello how are you?")
	PrintWordsTables(a)
}

func SplitWhiteSpaces(s string) []string {
	res := []string{}
	word := ""
	for i, c := range s {
		if i == len(s)-1 {
			word += string(c)
			res = append(res, word)
		} else if c != ' ' {
			word += string(c)
		} else {
			res = append(res, word)
			word = ""
		}
	}
	return res
}

func PrintWordsTables(a []string) {
	for _, c := range a {
		for _, v := range c {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}
*/

/* Q7-6
func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}

func Split(s, sep string) []string {
	str := ""
	slice := []string{}
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			str += string(s[i])
			slice = append(slice, str)
		}
		if i+len(sep) < len(s) && s[i:i+len(sep)] == sep {
			slice = append(slice, str)
			str = ""
			i += len(sep) - 1
			continue
		}
		str += string(s[i])
	}
	return slice
}
*/

/* Q7-7
func main() {
	result := ConvertBase("101011", "01", "0123456789")
	fmt.Println(result)
}

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
*/

//---------  Quest VIII  ---------//

/* Q8-1
func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	if isEven(len(os.Args[1:])) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
*/

/* Q8-2
type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}
*/

/* Q8-3
func main() {
	if len(os.Args) == 1 {
		fmt.Println("File name missing")
		return
	} else if len(os.Args) > 2 {
		fmt.Println("Too many argiments")
		return
	}
	content, _ := ioutil.ReadFile(os.Args[1])
	fmt.Println(string(content))
}
*/

/* Q8-4
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
*/

//---------  Quest IX  ---------//

/* Q9-1
func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	ForEach(PrintNbr, a)
}

func PrintNbr(n int) {
	fmt.Print(n)
}

func ForEach(f func(int), a []int) {
	for _, c := range a {
		f(c)
	}
}
*/

/* Q9-2
func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	result := Map(IsPrime, a)
	fmt.Println(result)
}

func IsPrime(n int) bool {
	if n < 2 {
		return false
	} else if n > 2 {
		for i := 2; i < n; i++ {
			if n%i == 0 {
				return false
			}
		}
	}
	return true
}

func Map(f func(int) bool, a []int) []bool {
	res := []bool{}
	for _, c := range a {
		res = append(res, f(c))
	}
	return res
}
*/

/* Q9-3
func main() {
	a1 := []string{"Hello", "how", "are", "you"}
	a2 := []string{"This", "is", "4", "you"}

	result1 := Any(IsNumeric, a1)
	result2 := Any(IsNumeric, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}

func IsNumeric(s string) bool {
	for _, c := range s {
		if c >= '0' && c <= '9' {
			return true
		}
	}
	return false
}

func Any(f func(string) bool, a []string) bool {
	for _, c := range a {
		if f(c) {
			return true
		}
	}
	return false
}
*/

/* Q9-4
func main() {
	tab1 := []string{"Hello", "how", "are", "you"}
	tab2 := []string{"This", "1", "is", "4", "you"}
	answer1 := CountIf(IsNumeric, tab1)
	answer2 := CountIf(IsNumeric, tab2)
	fmt.Println(answer1)
	fmt.Println(answer2)
}

func IsNumeric(s string) bool {
	for _, c := range s {
		if c >= '0' && c <= '9' {
			return true
		}
	}
	return false
}

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, c := range tab {
		if f(c) {
			count++
		}
	}
	return count
}
*/

/* Q9-5
func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}

	result1 := IsSorted(f, a1)
	result2 := IsSorted(f, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}

func f(n, nb int) int {
	return n - nb
}

func IsSorted(f func(a, b int) int, a []int) bool {
	if f(a[0], a[1]) > 0 {
		for i := 1; i < len(a)-1; i++ {
			if f(a[i], a[i+1]) < 0 {
				return false
			}
		}
	} else {
		for i := 1; i < len(a)-1; i++ {
			if f(a[i], a[i+1]) > 0 {
				return false
			}
		}
	}
	return true
}
*/

/* Q9-6
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
*/

/* Q9-7
func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(result)

	fmt.Println(result)
}

func SortWordArr(a []string) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[i] {
				a[j], a[i] = a[i], a[j]
			}
		}
	}
}
*/

/* Q9-8
func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	AdvancedSortWordArr(result, Compare)
	fmt.Println(result)
}

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if Compare(a[j], a[i]) == -1 {
				a[j], a[i] = a[i], a[j]
			}
		}
	}
}

func Compare(a, b string) int {
	if a == b {
		return 0
	} else if a > b {
		return 1
	}
	return -1
}
*/

//---------  Quest X  ---------//

/* Q10-1
func main() {
	result := Rot14("Hello! How are You?")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func Rot14(s string) string {
	res := ""
	for _, c := range s {
		if c >= 'a' && c <= 'm' || c >= 'A' && c <= 'L' {
			res += string(c + 14)
		} else if c >= 'm' && c <= 'z' || c >= 'M' && c <= 'Z' {
			res += string(c - 12)
		} else {
			res += string(c)
		}
	}
	return res
}
*/

/* Q10-2
func main() {
	DescendComb()
}

func DescendComb() {
	for i := '9'; i >= '0'; i-- {
		for j := '9'; j >= '0'; j-- {
			for k := '9'; k >= '0'; k-- {
				for l := '9'; l >= '0'; l-- {
					if i >= k && j > l {
						if i == '0' && j == '1' && k == '0' && l == '0' {
							z01.PrintRune(i)
							z01.PrintRune(j)
							z01.PrintRune(' ')
							z01.PrintRune(k)
							z01.PrintRune(l)
						} else {
							z01.PrintRune(i)
							z01.PrintRune(j)
							z01.PrintRune(' ')
							z01.PrintRune(k)
							z01.PrintRune(l)
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
}
*/

/* Q10-3
func main() {
	a := []int{1, 2, 3, 1, 2, 3, 4}
	unmatch := Unmatch(a)
	fmt.Println(unmatch)
}

func Unmatch(a []int) int {
	tab := [10]int{}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[i] == a[j] {
				tab[a[i]]++
			}
		}
	}
	for i, c := range tab {
		if c%2 != 0 {
			return i
		}
	}
	return -1
}
*/

/* Q10-4
func main() {
	fmt.Println(FoodDeliveryTime("burger"))
	fmt.Println(FoodDeliveryTime("chips"))
	fmt.Println(FoodDeliveryTime("nuggets"))
	fmt.Println(FoodDeliveryTime("burger") + FoodDeliveryTime("chips") + FoodDeliveryTime("nuggets"))
}

func FoodDeliveryTime(order string) int {
	if order == "burger" {
		return 15
	} else if order == "chips" {
		return 10
	} else if order == "nuggets" {
		return 12
	} else {
		return 404
	}
}
*/

/* Q10-5
func main() {
	steps := CollatzCountdown(12)
	fmt.Println(steps)
}

func CollatzCountdown(start int) int {
	count := 0
	for start > 1 {
		if start%2 == 0 {
			start /= 2
		} else if start%2 == 1 {
			start = 3*start + 1
		}
		count++
	}
	return count
}
*/

/* Q10-6
func main() {
	middle := Abort(2, 3, 8, 5, 7)
	fmt.Println(middle)
}

func Abort(a, b, c, d, e int) int {
	res := []int{a, b, c, d, e}
	for i := 0; i < len(res)-1; i++ {
		if res[i] > res[i+1] {
			res[i], res[i+1] = res[i+1], res[i]
			i=0
		}
	}
	return res[2]
}
*/

/* Q10-7
func main() {
	summary := "Burger Water            Carrot Coffee Water Water Chips Carrot Carrot Burger Carrot Water"
	for index, element := range ShoppingSummaryCounter(summary) {
		fmt.Println(index, "=>", element)
	}
}

func ShoppingSummaryCounter(str string) map[string]int {
	res := make(map[string]int)
	word := ""
	for i, c := range str {
		if i == len(str)-1 {
			word += string(c)
			res[word]++
		} else if c != ' ' {
			word += string(c)
		} else {
			if word == "" {
				res[word] = 2
			} else {
				res[word]++
				word = ""
			}
		}
	}
	return res
}
*/

/* Q10-8
const N = 6

func main() {
	a := make([]string, N)
	a[0] = "a"
	a[2] = "b"
	a[4] = "c"

	for _, v := range a {
		fmt.Println(v)
	}

	fmt.Println("Size after compacting:", Compact(&a))

	for _, v := range a {
		fmt.Println(v)
	}
}

func Compact(ptr *[]string) int {
	res := []string{}
	for _, c := range *ptr {
		if c != "" {
			res = append(res, c)
		}
	}
	*ptr = res
	return len(res)
}
*/

/* Q10-9
func main() {
	deck := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	DealAPackOfCards(deck)
}

func DealAPackOfCards (deck []int) {
	res := ""
	resFinal := []string{}
	for i := 0; i < len(deck); i++ {
		if i == 0 {
			res += strconv.Itoa(deck[i]) + ", "
		} else if i == len(deck)-1 {
			res += strconv.Itoa(deck[i])
			resFinal = append(resFinal, res)
		} else if i%3 != 0 {
			if (i+1)%3 == 0 {
				res += strconv.Itoa(deck[i])
			} else {
				res += strconv.Itoa(deck[i]) + ", "
			}
		} else {
			resFinal = append(resFinal, res)
			res = ""
			res += strconv.Itoa(deck[i]) + ", "
		}
	}
	for i, c := range resFinal {
		fmt.Printf("Player %d: %s", i+1, c)
		fmt.Println()
	}
}
*/

/* Q10-10
func main() {
	fmt.Print(JumpOver("1010101010"))
	fmt.Print(JumpOver(""))
	fmt.Print(JumpOver("t w e l v e"))
	fmt.Print(JumpOver("12"))
}

func JumpOver(str string) string {
	res := ""
	for i, c := range str {
		if (i+1)%3 == 0 {
			res += string(c)
		}
	}
	res += "\n"
	return res
}
*/

/* Q10-11
func main() {
	fmt.Println(StringToIntSlice("A quick brown fox jumps over the lazy dog"))
	fmt.Println(StringToIntSlice("Converted this string into an int"))
	fmt.Println(StringToIntSlice("hello THERE"))
}

func StringToIntSlice(str string) []int {
	res := []int{}
	for _, c := range str {
		res = append(res, int(c))
	}
	return res
}
*/

/* Q10-12
func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(Join(toConcat, ":"))
}

func Join(strs []string, sep string) string {
	res := ""
	for i, c := range strs {
		if i == len(strs)-1 {
			res += c
		} else  {
			res += c + sep
		}
	}
	return res
}
*/

/* Q10-13
func main() {
	x := 5
	y := &x
	z := &y
	a := &z

	w := 2
	b := &w

	u := 7
	e := &u
	f := &e
	g := &f
	h := &g
	i := &h
	j := &i
	c := &j

	k := 6
	l := &k
	m := &l
	n := &m
	d := &n

	fmt.Println(***a)
	fmt.Println(*b)
	fmt.Println(*******c)
	fmt.Println(****d)

	Enigma(a, b, c, d)

	fmt.Println("After using Enigma")
	fmt.Println(***a)
	fmt.Println(*b)
	fmt.Println(*******c)
	fmt.Println(****d)
}

func Enigma(a ***int, b *int, c *******int, d ****int) {
	***a, *b, *******c, ****d = *b, ****d, ***a, *******c
}
*/

/* Q10-14
func main() {
	fmt.Println(DescendAppendRange(10, 5))
	fmt.Println(DescendAppendRange(5, 10))
}

func DescendAppendRange(max, min int) []int {
	res := []int{}
	for i := max; i > min; i-- {
		res = append(res, i)
	}
	return res
}
*/

/* Q10-15
func main() {
	slice := []string{"Pineapple", "Honey", "Mushroom", "Tea", "Pepper", "Milk"}
	fmt.Println(ShoppingListSort(slice))
}

func ShoppingListSort(slice []string) []string {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if len(slice[i]) > len(slice[j]) {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return slice
}
*/

/* Q10-16
func main() {
	for i := 0; i < len(os.Args); i++ {
		if os.Args[i] == "01" || os.Args[i] == "galaxy" || os.Args[i] == "galaxy 01" {
			fmt.Println("Alert!!!")
		}
	}
}
*/

/* Q10-17
func main() {
	fmt.Println(ReverseMenuIndex([]string{"desserts", "mains", "drinks", "starters"}))
}

func ReverseMenuIndex(menu []string) []string {
	res := []string{}
	for i := len(menu) - 1; i >= 0; i-- {
		res = append(res, menu[i])
	}
	return res
}
*/

/* Q10-18
func main() {

	p4 := []string{"4th Place"}
	p3 := []string{"3rd Place"}
	p2 := []string{"2nd Place"}
	p1 := []string{"1st Place"}

	position := [][]string{p4, p3, p2, p1}
	fmt.Println(PodiumPosition(position))
}

func PodiumPosition(podium [][]string) [][]string {
	for i := 0; i < len(podium); i++ {
		for j := i + 1; j < len(podium); j++ {
			if podium[i][0] > podium[j][0] {
				podium[i], podium[j] = podium[j], podium[i]
			}
		}
	}
	return podium
}
*/

/* Q10-19
func main() {
	fmt.Println(ActiveBits(7))
}

func ActiveBits(n int) int {
	count := 0
	for n > 0 {
		if n%2 == 1 {
			count++
		}
		n /= 2
	}
	return count
}
*/

/* Q10-20
type Pilot struct {
	Name     string
	Life     float64
	Age      int
	Aircraft int
}

func main() {
	var donnie Pilot
	donnie.Name = "Donnie"
	donnie.Life = 100.0
	donnie.Age = 24
	donnie.Aircraft = AIRCRAFT1

	fmt.Println(donnie)
}

const AIRCRAFT1 = 1
*/

/* Q10-21
type Door struct {
	state string
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...")
	ptrDoor.state = "CLOSE"
}

func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...")
	ptrDoor.state = "OPEN"
}

func IsDoorOpen(Door Door) bool {
	PrintStr("is the Door opened ?")
	z01.PrintRune('\n')
	return Door.state == "OPEN"
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	z01.PrintRune('\n')
	return ptrDoor.state == "CLOSE"
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(*door) {
		CloseDoor(door)
	}
	if door.state == "OPEN" {
		CloseDoor(door)
	}
}
*/

/* Q10-22
func main() {
	a := []int{23, 123, 1, 11, 55, 93}
	max := Max(a)
	fmt.Println(max)
}

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a[len(a)-1]
}
*/

/* Q10-23
func main() {
	fmt.Println(RockAndRoll(4))
	fmt.Println(RockAndRoll(9))
	fmt.Println(RockAndRoll(6))
}

func RockAndRoll(n int) string {
	if n < 0 {
		return "error: number is negative\n"
	} else if n%2 == 0 && n%3 == 0 {
		return "rock and roll\n"
	} else if n%2 == 0 {
		return "rock\n"
	} else if n%3 == 0 {
		return "roll\n"
	}
	return "error: non divisible\n"
}
*/

/* Q10-24
func main() {
	fmt.Print(LoafOfBread("deliciousbread"))
	fmt.Print(LoafOfBread("This is a loaf of bread"))
	fmt.Print(LoafOfBread("loaf"))
}

func LoafOfBread(str string) string {
	res := ""
	resFinal := ""
	for _, c := range str {
		if c != ' ' {
			res += string(c)
		}
	}
	if len(res) < 5 {
		return "Invalid Output\n"
	}
	count := 0
	for _, c := range res {
		if count == 5 {
			resFinal += " "
			count = 0
			continue
		} else {
			resFinal += string(c)
		}
		count++
	}
	resFinal += "\n"
	return resFinal
}
*/

//---------  Quest XI  ---------//

/* Q11-1
func main() {

	link := &List{}

	ListPushBack(link, "Hello")
	ListPushBack(link, "man")
	ListPushBack(link, "how are you")

	for link.Head != nil {
		fmt.Println(link.Head.Data)
		link.Head = link.Head.Next
	}
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	Node := &NodeL{Data: data, Next: nil}
	if l.Head == nil {
		l.Head = Node
		l.Tail = Node
	} else {
		l.Tail.Next = Node
		l.Tail = Node
	}
}
*/

/* Q11-2
func main() {

	link := &List{}

	ListPushFront(link, "Hello")
	ListPushFront(link, "man")
	ListPushFront(link, "how are you")

	it := link.Head
	for it != nil {
		fmt.Print(it.Data, " ")
		it = it.Next
	}
	fmt.Println()
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	Node := &NodeL{Data: data, Next: nil}
	if l.Head == nil {
		l.Head = Node
		l.Tail = Node
	} else {
		Node.Next = l.Head
		l.Head = Node
	}
}
*/

/* Q11-3
func main() {
	link := &List{}

	ListPushFront(link, "Hello")
	ListPushFront(link, "2")
	ListPushFront(link, "you")
	ListPushFront(link, "man")

	fmt.Println(ListSize(link))
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}
type List struct {
	Head *NodeL
	Tail *NodeL
}
func ListSize(l *List) int {
	count := 0
	for l.Head != nil {
		count++
		l.Head = l.Head.Next
	}
	return count
}
*/

/* Q11-4
func main() {
	link := &List{}
	link2 := &List{}

	ListPushBack(link, "three")
	ListPushBack(link, 3)
	ListPushBack(link, "1")

	fmt.Println(ListLast(link))
	fmt.Println(ListLast(link2))
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	for l.Head.Next != nil {
		l.Head = l.Head.Next
	}
	return l.Head.Data
}
*/

/* Q11-5
type List = List
type Node = NodeL

func PrintList(l *List) {
	link := l.Head
	for link != nil {
		fmt.Print(link.Data, " -> ")
		link = link.Next
	}
	fmt.Println(nil)
}

func main() {
	link := &List{}

	ListPushBack(link, "I")
	ListPushBack(link, 1)
	ListPushBack(link, "something")
	ListPushBack(link, 2)

	fmt.Println("------list------")
	PrintList(link)
	ListClear(link)
	fmt.Println("------updated list------")
	PrintList(link)
}

func ListClear(l *List) {
	l.Head = nil
}
*/

/* Q11-6
func main() {
	link := &piscine.List{}

	piscine.ListPushBack(link, "hello")
	piscine.ListPushBack(link, "how are")
	piscine.ListPushBack(link, "you")
	piscine.ListPushBack(link, 1)

	fmt.Println(piscine.ListAt(link.Head, 3).Data)
	fmt.Println(piscine.ListAt(link.Head, 1).Data)
	fmt.Println(piscine.ListAt(link.Head, 7))
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil {
		return nil
	}
	for i := 0; i < pos; i++ {
		if l.Next == nil {
			return nil
		}
		l = l.Next
	}
	return l
}
*/

/* Q11-7
func main() {
	link := &piscine.List{}

	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, 2)
	piscine.ListPushBack(link, 3)
	piscine.ListPushBack(link, 4)

	piscine.ListReverse(link)

	it := link.Head

	for it != nil {
		fmt.Println(it.Data)
		it = it.Next
	}

	fmt.Println("Tail", link.Tail)
	fmt.Println("Head", link.Head)
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListReverse(l *List) {
	var prev, current *NodeL
	current = l.Head
	l.Tail = l.Head
	for current != nil {
		prev, current, current.Next = current, current.Next, prev
	}
	l.Head = prev
}
*/

/* Q11-8
func main() {
	link := &piscine.List{}

	piscine.ListPushBack(link, "1")
	piscine.ListPushBack(link, "2")
	piscine.ListPushBack(link, "3")
	piscine.ListPushBack(link, "5")

	piscine.ListForEach(link, piscine.Add2_node)

	it := link.Head
	for it != nil {
		fmt.Println(it.Data)
		it = it.Next
	}
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListForEach(l *List, f func(*NodeL)) {
	current := l.Head
	for current != nil {
		f(current)
		current = current.Next
	}
}

func Add2_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) + 2
	case string:
		node.Data = node.Data.(string) + "2"
	}
}

func Subtract3_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) - 3
	case string:
		node.Data = node.Data.(string) + "-3"
	}
}
*/

/* Q11-9
func PrintElem(node *piscine.NodeL) {
	fmt.Println(node.Data)
}

func StringToInt(node *piscine.NodeL) {
	node.Data = 2
}

func PrintList(l *piscine.List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, "->")
		it = it.Next
	}
	fmt.Print("nil","\n")
}

func main() {
	link := &piscine.List{}

	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, "hello")
	piscine.ListPushBack(link, 3)
	piscine.ListPushBack(link, "there")
	piscine.ListPushBack(link, 23)
	piscine.ListPushBack(link, "!")
	piscine.ListPushBack(link, 54)

	PrintList(link)

	fmt.Println("--------function applied--------")
	piscine.ListForEachIf(link, PrintElem, piscine.IsPositiveNode)

	piscine.ListForEachIf(link, StringToInt, piscine.IsAlNode)

	fmt.Println("--------function applied--------")
	PrintList(link)

	fmt.Println()
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func IsPositiveNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return node.Data.(int) > 0
	default:
		return false
	}
}

func IsAlNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return false
	default:
		return true
	}
}

func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
	it := l.Head
	for it != nil {
		if cond(it) {
			f(it)
		}
		it = it.Next
	}
}
*/

/* Q4-9 / Q8-5 / Q11-10 / Q11-11 / Q11-12 / Q11-13 / Q11-14 / Q11-15 */

/* MAP / STRUCT / MAKE / OS.FUNCTION */
