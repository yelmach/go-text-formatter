package helpers

import (
	"regexp"
	"strings"
)

// articles
func CorrectArticles(words []string) []string {
	for i := 0; i < len(words); i++ {
		if i != len(words)-1 && (strings.ToLower(words[i]) == "a" || strings.ToLower(words[i]) == "'a") {
			if isVowel(words[i+1]) {
				words[i] += "n"
			}
		}
	}
	return words
}

// correct punctuations
func CorrectPunctuations(text string) string {
	punc := regexp.MustCompile(`\s*([.,!?:;])`)
	corrected := punc.ReplaceAllString(text, "$1")
	punc = regexp.MustCompile(`\s*([.,!?:;]+)\s*`)
	corrected = punc.ReplaceAllString(corrected, "${1} ")
	return corrected
}
