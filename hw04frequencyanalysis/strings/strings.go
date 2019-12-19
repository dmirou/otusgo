package strings

import (
	"strings"
)

// WordsCount calculates a count of the each word in the string
func WordsCount(s string) map[string]int {
	var sWithSpace = s + " "
	var wordToCount = make(map[string]int, 100)
	var builder strings.Builder
	var isBreak bool
	var countAdded bool
	for _, r := range sWithSpace {
		isBreak = strings.ContainsAny(string(r), ".,- ")
		if !isBreak {
			builder.WriteRune(r)
			continue
		}
		if builder.Len() == 0 {
			continue
		}
		word := builder.String()
		builder.Reset()
		countAdded = false
		for key := range wordToCount {
			if strings.EqualFold(word, key) {
				wordToCount[key]++
				countAdded = true
				break
			}
		}
		if !countAdded {
			wordToCount[word]++
		}
	}
	return wordToCount
}

// Top10 returns ten or less common words per line
func Top10(s string) []string {
	var topWords []string
	var wordToCount = WordsCount(s)
	if len(wordToCount) == 0 {
		return topWords
	}

	var maxCount int
	for _, val := range wordToCount {
		if maxCount < val {
			maxCount = val
		}
	}

	const wordLimit = 10
	for word, count := range wordToCount {
		if len(topWords) >= wordLimit {
			break
		}
		if count == maxCount {
			topWords = append(topWords, word)
		}
	}
	return topWords
}
