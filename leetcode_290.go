func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	charToWord := make(map[rune]string)
	wordToChar := make(map[string]rune)
	for i, char := range pattern {
		word := words[i]
		if mapped, exists := charToWord[char]; exists {
			if mapped != word {
				return false
			}
		} else {
			charToWord[char] = word
		}

		if mapped, exists := wordToChar[word]; exists {
			if mapped != char {
				return false
			}
		} else {
			wordToChar[word] = char
		}
	}

	return true
}
