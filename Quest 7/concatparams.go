package piscine

func ConcatParams(args []string) string {
	result := ""
	for i, arg := range args {
		result += arg
		if i < len(args)-1 {
			result += "\n"
		}
	}
	return result
}
