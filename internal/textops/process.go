package textops

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// Process runs all transformation functions in order.
func Process(text string) string {
	text = handleHex(text)
	text = handleBin(text)
	text = handleCase(text)
	text = handlePunctuation(text)
	text = handleArticles(text)
	text = handleQuotes(text)
	return text
}

// ---------------------
// 1️⃣ Hexadecimal handler
func handleHex(text string) string {
	re := regexp.MustCompile(`\b([0-9a-fA-F]+) \(hex\)`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		parts := strings.Split(match, " ")
		val, err := strconv.ParseInt(parts[0], 16, 64)
		if err != nil {
			return match
		}
		return strconv.FormatInt(val, 10)
	})
}

// ---------------------
// 2️⃣ Binary handler
func handleBin(text string) string {
	re := regexp.MustCompile(`\b([01]+) \(bin\)`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		parts := strings.Split(match, " ")
		val, err := strconv.ParseInt(parts[0], 2, 64)
		if err != nil {
			return match
		}
		return strconv.FormatInt(val, 10)
	})
}

// ---------------------
// 3️⃣ Casing handler
func handleCase(text string) string {
	words := strings.Fields(text)
	markerRe := regexp.MustCompile(`^\s*([a-zA-Z]+)(?:\s*(?:,|\s)\s*(\d+))?\s*$`)

	for i := 0; i < len(words); i++ {
		word := words[i]

		// Trim trailing punctuation so markers like "(cap, 6)," are detected
		trimmedRight := strings.TrimRight(word, ".,!?:;\"'")

		// If marker spans multiple tokens (e.g. "(cap, 6)") merge tokens until closing ')'
		merged := trimmedRight
		lastOrig := word
		endIdx := i
		if strings.HasPrefix(trimmedRight, "(") && !strings.HasSuffix(trimmedRight, ")") {
			j := i + 1
			for j < len(words) {
				lastOrig = words[j]
				part := strings.TrimRight(words[j], ".,!?:;\"'")
				merged += " " + part
				endIdx = j
				if strings.HasSuffix(part, ")") {
					break
				}
				j++
			}
		}

		// Now check merged (which may equal trimmedRight) for a marker
		if strings.HasPrefix(merged, "(") && strings.HasSuffix(merged, ")") {
			markerRaw := strings.Trim(merged, "()")
			m := markerRe.FindStringSubmatch(markerRaw)
			if m == nil {
				continue
			}
			markerName := strings.ToLower(m[1])
			count := 1
			if m[2] != "" {
				if cnt, err := strconv.Atoi(m[2]); err == nil {
					count = cnt
				}
			}

			// Apply to previous 'count' words
			for j := 0; j < count && i-1-j >= 0; j++ {
				switch markerName {
				case "up":
					words[i-1-j] = strings.ToUpper(words[i-1-j])
				case "low":
					words[i-1-j] = strings.ToLower(words[i-1-j])
				case "cap":
					words[i-1-j] = capitalize(words[i-1-j])
				}
			}

			// Preserve any trailing punctuation that was after the marker (e.g., comma)
			rOrig := []rune(lastOrig)
			rTrim := []rune(strings.TrimRight(lastOrig, ".,!?:;\"'"))
			trailing := ""
			if len(rOrig) > len(rTrim) {
				trailing = string(rOrig[len(rTrim):])
			}

			// Replace the first marker token with its trailing punctuation (or empty)
			words[i] = trailing
			// Clear any tokens that were part of the merged marker (except the first)
			for k := i + 1; k <= endIdx && k < len(words); k++ {
				words[k] = ""
			}

			// Move index past consumed tokens
			i = endIdx
		}
	}

	// Remove empty markers
	return strings.Join(filterEmpty(words), " ")
}

// ---------------------
// 4️⃣ Punctuation handler
func handlePunctuation(text string) string {
	// Keep groups like ... and !? tight, single space after
	text = regexp.MustCompile(`\s*\.\.\.\s*`).ReplaceAllString(text, "...")
	text = regexp.MustCompile(`\s*!\?\s*`).ReplaceAllString(text, "!?")

	// Single punctuation hugs previous word and leave one space after
	puncts := []string{".", ",", "!", "?", ":", ";"}
	for _, p := range puncts {
		re := regexp.MustCompile(`\s*` + regexp.QuoteMeta(p) + `\s*`)
		text = re.ReplaceAllString(text, p+" ")
	}

	// Remove extra space before punctuation (but keep single space after)
	text = regexp.MustCompile(`\s+([.,!?:;])`).ReplaceAllString(text, "$1")
	return strings.TrimSpace(text)
}

// ---------------------
// 5️⃣ Quotes handler
func handleQuotes(text string) string {
	re := regexp.MustCompile(`'\s*(.*?)\s*'`)
	return re.ReplaceAllString(text, "'$1'")
}

// ---------------------
// 6️⃣ Articles handler
func handleArticles(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words)-1; i++ {
		word := words[i]
		if strings.ToLower(word) != "a" {
			continue
		}

		// Look ahead for next "real" word
		nextWord := ""
		for j := i + 1; j < len(words); j++ {
			candidate := words[j]
			if len(candidate) == 0 {
				continue
			}
			r := rune(candidate[0])
			if unicode.IsLetter(r) {
				nextWord = candidate
				break
			}
		}

		if nextWord == "" {
			continue
		}

		firstRune := unicode.ToLower(rune(nextWord[0]))
		if strings.ContainsRune("aeiouh", firstRune) {
			words[i] = "an"
		}
	}

	return strings.Join(words, " ")
}

func capitalize(s string) string {
	if s == "" {
		return s
	}
	r := []rune(strings.ToLower(s))
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

// ---------------------
// Helper: remove empty strings
func filterEmpty(words []string) []string {
	var out []string
	for _, w := range words {
		if w != "" {
			out = append(out, w)
		}
	}
	return out
}
