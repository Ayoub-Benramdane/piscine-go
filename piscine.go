package main

import (
	"os"
	"strconv"
)

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
	res := []string{}
	word := ""
	for i := 0; i < len(s); i++ {
		if i+len(sep) >= len(s)-len(sep) && i < len(s) {
			word += string(s[i])
			if i == len(s)-1 {
				res = append(res, word)
			}
		} else if s[i:i+len(sep)] != sep {
			word += string(s[i])
		} else {
			res = append(res, word)
			i += len(sep) -1
			word = ""
		}
	}
	return res
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

//---------  Quest VIII  ---------//

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



/* Q2-8 / Q3-12 / Q4-9 / Q5-20 / Q5-21 / Q6-5 / Q6-6 / Q6-7 / Q7-7 / Q8-4 / Q8-5 / Q9-7 / Q9-8 */
