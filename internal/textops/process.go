package textops

// Process runs all transformation functions in order.
func Process(text string) string {
	text = handleHex(text)
	text = handleBin(text)
	text = handleCase(text)
	text = handlePunctuation(text)
	text = handleArticles(text)
	return text
}

// Placeholder functions (implement later)
func handleHex(s string) string         { return s }
func handleBin(s string) string         { return s }
func handleCase(s string) string        { return s }
func handlePunctuation(s string) string { return s }
func handleArticles(s string) string    { return s }
