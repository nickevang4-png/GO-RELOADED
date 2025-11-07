package processor

import "strings"

var punctuationGroups = []string{"...", "!?"}

func ApplyPunctuation(tokens []string) []string {
	result := []string{}
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		// Skip markers if any
		if token == "" {
			continue
		}

		// Handle punctuation groups
		for _, group := range punctuationGroups {
			if token == group {
				result[len(result)-1] += group
				continue
			}
		}

		// Normal punctuation
		if strings.ContainsAny(token, ".,!?;:") {
			if len(result) > 0 {
				result[len(result)-1] += token
				continue
			}
		}

		result = append(result, token)
	}

	return result
}
