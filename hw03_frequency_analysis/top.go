package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	var words = strings.Fields(text)
	var counterWords = map[string]int{}
	for _, word := range words {
		counterWords[word]++
	}
	type counterWordEntries struct {
		word  string
		count int
	}
	var counterWordsSlice []counterWordEntries
	for word, count := range counterWords {
		counterWordsSlice = append(counterWordsSlice, counterWordEntries{word: word, count: count})
	}

	sort.Slice(counterWordsSlice, func(i, j int) bool {
		if counterWordsSlice[i].count < counterWordsSlice[j].count {
			return false
		}
		if counterWordsSlice[i].count == counterWordsSlice[j].count {
			return strings.Compare(counterWordsSlice[i].word, counterWordsSlice[j].word) == -1
		}
		return true
	})

	var result []string
	for i := 0; i < countTopWords && i < len(counterWords); i++ {
		result = append(result, counterWordsSlice[i].word)
	}

	return result
}

const countTopWords = 10
