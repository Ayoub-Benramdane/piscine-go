package piscine

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	answer := ""
	count := 0
	for _, value := range str {
		if value != ' ' && count != 5 {
			answer += string(value)
			count++
		} else if count == 5 {
			answer += " "
			count = 0
		}
	}
	answer += "\n"
	return answer
}
