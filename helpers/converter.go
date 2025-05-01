package helpers

import (
	"fmt"
	"strconv"
	"strings"
)

// Bin Hex
func ConvertToDecemal(words []string) []string {
	for i := 0; i < len(words); i++ {
		switch {
		case strings.ToLower(words[i]) == "(hex)":
			if i != 0 {
				num, err := strconv.ParseInt(words[i-1], 16, 32)
				if err != nil {
					fmt.Println("[-] There is an Error:", err)
				}
				if num != 0 {
					words[i-1] = fmt.Sprint(num)
				}
			}
			removeWord(&words, i)
			i--

		case strings.ToLower(words[i]) == "(bin)":
			if i != 0 {
				num, err := strconv.ParseInt(words[i-1], 2, 32)
				if err != nil {
					fmt.Println("[-] There is an Error:", err)
				}
				if num != 0 {
					words[i-1] = fmt.Sprint(num)
				}
			}

			removeWord(&words, i)
			i--
		}
	}
	return words
}
