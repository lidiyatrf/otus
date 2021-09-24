package hw03frequencyanalysis

import (
	"math"
	"regexp"
	"sort"
	"strings"
)

var (
	regexSymbols = regexp.MustCompile("[^А-Яа-яA-Za-z-]+")
	regexWord    = regexp.MustCompile(`([А-Яа-яA-Za-z]+)|\\[-]`)
)

func Top10(str string) []string {
	wordOccurrence := make(map[string]int)

	for _, row := range strings.Fields(str) {
		for _, word := range strings.Split(row, " ") {
			word = strings.ToLower(word)
			var found bool
			for _, w := range regexSymbols.Split(word, -1) {
				if regexWord.MatchString(w) {
					word = w
					found = true
					break
				}
			}
			if !found {
				continue
			}
			if count, ok := wordOccurrence[word]; ok {
				wordOccurrence[word] = count + 1
			} else {
				wordOccurrence[word] = 1
			}
		}
	}

	occurrenceWords := make(map[int][]string)
	maxOccurrence := math.MinInt32
	for k, v := range wordOccurrence {
		if v > maxOccurrence {
			maxOccurrence = v
		}
		occurrenceWords[v] = append(occurrenceWords[v], k)
	}

	result := make([]string, 0, 10)
	for i := maxOccurrence; i >= 0; i-- {
		words, hasValues := occurrenceWords[i]
		if !hasValues {
			continue
		}
		sort.Strings(words)
		for _, word := range words {
			if len(result) == 10 {
				return result
			}
			result = append(result, word)
		}
	}

	return result
}
