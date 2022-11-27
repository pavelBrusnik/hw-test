package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var regExp = regexp.MustCompile(`[a-zA-Zа-яА-Я]+(-*[a-zA-Zа-яА-Я]+)*`)

type wordCount struct {
	word  string
	count int
}

func Top10(inputStr string) []string {
	if len(inputStr) == 0 {
		return make([]string, 0)
	}

	words := regExp.FindAllString(inputStr, -1)
	wordsCountMap := make(map[string]int, len(words))

	for _, word := range words {
		wordsCountMap[strings.ToLower(word)]++
	}

	wordsCounts := make([]wordCount, 0)

	for word, count := range wordsCountMap {
		wordsCounts = append(wordsCounts, wordCount{word, count})
	}

	sort.Slice(wordsCounts, func(i, j int) bool {
		if wordsCounts[i].count == wordsCounts[j].count {
			return wordsCounts[j].word > wordsCounts[i].word
		}
		return wordsCounts[i].count > wordsCounts[j].count
	})

	var topWordsCount int

	if len(wordsCounts) < 10 {
		topWordsCount = len(wordsCounts)
	} else {
		topWordsCount = 10
	}

	topWords := make([]string, topWordsCount)

	for i, tWord := range wordsCounts {
		if i == topWordsCount {
			break
		}

		topWords[i] = tWord.word
	}

	return topWords
}
