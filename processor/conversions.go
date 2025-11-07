package processor

import (
	"strconv"
)

func ApplyConversions(tokens []string) []string {
	result := []string{}

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		if token == "(hex)" && i > 0 {
			val, _ := strconv.ParseInt(result[len(result)-1], 16, 64)
			result[len(result)-1] = strconv.FormatInt(val, 10)
		} else if token == "(bin)" && i > 0 {
			val, _ := strconv.ParseInt(result[len(result)-1], 2, 64)
			result[len(result)-1] = strconv.FormatInt(val, 10)
		} else {
			result = append(result, token)
		}
	}

	return result
}
