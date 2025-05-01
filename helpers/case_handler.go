package helpers

import (
	"fmt"
	"strconv"
	"strings"
)

// UP LOW CAP
func ConvertCases(words []string) []string {
	for i := 0; i < len(words); i++ {

		switch {
		case words[i] == "(up)":
			if i != 0 {
				words[i-1] = strings.ToUpper(words[i-1])
			}
			removeWord(&words, i)
			i--

		case words[i] == "(low)":
			if i != 0 {
				words[i-1] = strings.ToLower(words[i-1])
			}
			removeWord(&words, i)
			i--

		case words[i] == "(cap)":
			if i != 0 {
				words[i-1] = capitalize(words[i-1])
			}
			removeWord(&words, i)
			i--

		case words[i] == "(up,":
			if i == len(words)-1 {
				continue
			}
			if words[i+1][len(words[i+1])-1] != ')' {
				continue
			}
			nbr, err := strconv.Atoi(words[i+1][:len(words[i+1])-1])
			if err != nil {
				fmt.Println("[-] There is an Error: you give invalid syntax")
				continue
			}
			if nbr > len(words[:i]) {
				fmt.Println("[-] WARNING: the number you give is higher than the number of words")
				nbr = len(words[:i])
			}
			for j := 1; j <= nbr; j++ {
				words[i-j] = strings.ToUpper(words[i-j])
			}
			removeWord(&words, i)
			removeWord(&words, i)
			i--

		case words[i] == "(low,":
			if i == len(words)-1 {
				continue
			}
			if words[i+1][len(words[i+1])-1] != ')' {
				continue
			}
			nbr, err := strconv.Atoi(words[i+1][:len(words[i+1])-1])
			if err != nil {
				fmt.Println("[-] There is an Error: you give invalid syntax")
				continue
			}
			if nbr > len(words[:i]) {
				fmt.Println("[-] WARNING: the number you give is higher than the number of words")
				nbr = len(words[:i])
			}
			for j := 1; j <= nbr; j++ {
				words[i-j] = strings.ToLower(words[i-j])
			}
			removeWord(&words, i)
			removeWord(&words, i)
			i--

		case words[i] == "(cap,":
			if i == len(words)-1 {
				continue
			}
			if words[i+1][len(words[i+1])-1] != ')' {
				continue
			}
			nbr, err := strconv.Atoi(words[i+1][:len(words[i+1])-1])
			if err != nil {
				fmt.Println("[-] There is an Error: you give invalid syntax")
				continue
			}
			if nbr > len(words[:i]) {
				fmt.Println("[-] WARNING: the number you give is higher than the number of words")
				nbr = len(words[:i])
			}
			for j := 1; j <= nbr; j++ {
				words[i-j] = capitalize(words[i-j])
			}

			removeWord(&words, i)
			removeWord(&words, i)
			i--
		}
	}
	return words
}
