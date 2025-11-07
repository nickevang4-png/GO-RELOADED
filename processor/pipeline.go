package processor

func Process(tokens []string) []string {
	tokens = ApplyConversions(tokens)
	tokens = ApplyCasing(tokens)
	tokens = ApplyQuotes(tokens) // You’ll implement quote hugging here
	tokens = ApplyPunctuation(tokens)
	tokens = ApplyArticles(tokens)
	return tokens
}
