package main

import (
	"converter/helpers"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("[-] Error: only use 2 Arguments")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	lines, err := helpers.SplitFileIntoLines(inputFile)
	if err != nil {
		fmt.Println("error reading file")
		return
	}

	result := []string{}
	for i := 0; i < len(lines); i++ {
		words := strings.Fields(string(lines[i]))
		words = helpers.ConvertCases(words)
		words = helpers.ConvertToDecemal(words)
		words = helpers.CorrectArticles(words)
		line := strings.Join(words, " ")
		line = helpers.CorrectPunctuations(line)
		line = helpers.FormatQuotes(line)
		line = helpers.CleanSpaces(line)
		result = append(result, line)
	}

	res := strings.Join(result, "\n")
	os.WriteFile(outputFile, []byte(res), 0o644)
}
