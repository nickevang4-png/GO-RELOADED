package processor

import (
	"strconv"
	"strings"
)

func ApplyCasing(tokens []string) []string {
	result := []string{}
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		// Simple one-word markers
		switch {
		case token == "(up)":
			if len(result) > 0 {
				result[len(result)-1] = strings.ToUpper(result[len(result)-1])
			}
			continue
		case token == "(low)":
			if len(result) > 0 {
				result[len(result)-1] = strings.ToLower(result[len(result)-1])
			}
			continue
		case token == "(cap)":
			if len(result) > 0 {
				result[len(result)-1] = strings.Title(result[len(result)-1])
			}
			continue
		// Counted casing e.g., (up, 3)
		case strings.HasPrefix(token, "(up,"):
			n := parseCount(token)
			applyToLastN(result, n, strings.ToUpper)
			continue
		case strings.HasPrefix(token, "(low,"):
			n := parseCount(token)
			applyToLastN(result, n, strings.ToLower)
			continue
		case strings.HasPrefix(token, "(cap,"):
			n := parseCount(token)
			applyToLastN(result, n, strings.Title)
			continue
		}

		result = append(result, token)
	}

	return result
}

func parseCount(token string) int {
	token = strings.TrimSuffix(strings.TrimPrefix(token, "("), ")")
	parts := strings.Split(token, ",")
	n, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
	return n
}

func applyToLastN(tokens []string, n int, fn func(string) string) {
	start := len(tokens) - n
	if start < 0 {
		start = 0
	}
	for i := start; i < len(tokens); i++ {
		tokens[i] = fn(tokens[i])
	}
}
