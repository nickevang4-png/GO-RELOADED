package processor

// ApplyQuotes handles single quote hugging: 'word' or 'multiple words'
func ApplyQuotes(tokens []string) []string {
	result := []string{}
	insideQuote := false
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		if token == "'" {
			if !insideQuote {
				// opening quote
				insideQuote = true
				result = append(result, "'")
			} else {
				// closing quote
				insideQuote = false
				result = append(result, "'")
			}
		} else {
			result = append(result, token)
		}
	}
	return result
}
