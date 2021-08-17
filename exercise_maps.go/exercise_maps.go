package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func isStringInTheSlice(word string, list []string) bool {
	for _, element := range list {
		if word == element {
			return true
		}
	}

	return false
}

func WordCount(s string) map[string]int {

	// splitting words
	splittedStrings := strings.Split(s, " ")
	countOfWords := make(map[string]int)

	for _, value := range splittedStrings {
		countOfWords[value] = 0
	}

	for _, value := range splittedStrings {
		valueInString := isStringInTheSlice(value, splittedStrings)
		if valueInString {
			wordCount := countOfWords[value]
			wordCount++
			countOfWords[value] = wordCount
		}
	}

	return countOfWords
}

func main() {
	wc.Test(WordCount)
}
