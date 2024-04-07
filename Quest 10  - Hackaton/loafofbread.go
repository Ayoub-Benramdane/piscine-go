package piscine

func LoafOfBread(str string) string {
	if str == "" {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	answer := ""
	counter := 0
	for i, value := range str {
		if value != ' ' && counter != 5 {
			answer += string(value)
			counter++
		} else if counter == 5 {
			answer += " "
			counter = 0
		}
		if i == len(str)-1 && len(answer) > 0 && answer[len(answer)-1] == ' ' {
			answer = answer[:len(answer)-1]
		}
	}
	answer += "\n"
	return answer
}
