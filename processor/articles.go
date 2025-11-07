package processor

import "strings"

func ApplyArticles(tokens []string) []string {
	vowels := "aeiouhAEIOUH"
	for i := 0; i < len(tokens)-1; i++ {
		if strings.ToLower(tokens[i]) == "a" && len(tokens[i+1]) > 0 {
			firstChar := tokens[i+1][0]
			if strings.ContainsRune(vowels, rune(firstChar)) {
				tokens[i] = "an"
			}
		}
	}
	return tokens
}
