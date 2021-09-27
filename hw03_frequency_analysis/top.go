package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var (
	regexSymbols = regexp.MustCompile("[^А-Яа-яA-Za-z-]+")
	regexWord    = regexp.MustCompile(`([А-Яа-яA-Za-z]+)|\\[-]`)
)

type wordOccStruct struct {
	word   string
	amount int
}

func Top10(str string) []string {
	wordToAmountMap := mapWordToAmount(str)
	array := convertMapToArray(wordToAmountMap)
	result := getTop10Words(array)
	return result
}

func mapWordToAmount(str string) map[string]int {
	result := make(map[string]int)

	for _, row := range strings.Fields(str) {
		for _, word := range strings.Split(row, " ") {
			word = strings.ToLower(word)
			word, ok := extractWord(word)
			if !ok {
				continue
			}
			result[word]++
		}
	}

	return result
}

func extractWord(word string) (string, bool) {
	for _, w := range regexSymbols.Split(word, -1) {
		if regexWord.MatchString(w) {
			return w, true
		}
	}
	return "", false
}

func convertMapToArray(wordToAmountMap map[string]int) []wordOccStruct {
	array := make([]wordOccStruct, len(wordToAmountMap))
	count := 0
	for k, v := range wordToAmountMap {
		array[count] = wordOccStruct{
			word:   k,
			amount: v,
		}
		count++
	}
	return array
}

func getTop10Words(array []wordOccStruct) []string {
	arrToSort := make([]wordOccStruct, len(array))
	copy(arrToSort, array)
	sort.Slice(arrToSort, func(i, j int) bool {
		if arrToSort[i].amount == arrToSort[j].amount {
			return arrToSort[i].word < arrToSort[j].word
		}
		return arrToSort[i].amount > arrToSort[j].amount
	})

	var result []string
	for i := 0; i < 10 && i < len(arrToSort); i++ {
		result = append(result, arrToSort[i].word)
	}
	return result
}
