package helpers

import (
	"regexp"
	"strings"
)

// correct quotations
func FormatQuotes(text string) string {
	offon := false
	slice := []string{}
	id := 0
	for i := 0; i < len(text); i++ {
		if text[i] == '\'' {
			if i != 0 && i != len(text)-1 {
				if text[i-1] == ' ' || text[i+1] == ' ' {
					if !offon {
						slice = append(slice, text[id:i])
						id = i
						offon = true
					} else {
						slice = append(slice, text[id:i+1])
						id = i + 1
						offon = false
					}
				}
			} else {
				if !offon {
					slice = append(slice, text[id:i])
					id = i
					offon = true
				} else {
					slice = append(slice, text[id:i+1])
					id = i + 1
					offon = false
				}
			}
		}
	}
	slice = append(slice, text[id:])
	str := ""
	for i := 0; i < len(slice); i++ {
		reQuotes := regexp.MustCompile(`' *(.*[\S]) *'`)
		slice[i] = reQuotes.ReplaceAllString(slice[i], "'$1'")
		str += slice[i]
	}
	return str
}

// clean spaces
func CleanSpaces(text string) string {
	space := regexp.MustCompile(`\s+`)
	corrected := space.ReplaceAllString(text, " ")
	corrected = strings.TrimSpace(corrected)
	return corrected
}
