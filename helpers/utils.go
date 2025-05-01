package helpers

import (
	"os"
	"strings"
)

// removeWord function, to remove the cases
func removeWord(words *[]string, index int) {
	*words = append((*words)[:index], (*words)[index+1:]...)
}

func SplitFileIntoLines(file string) ([]string, error) {
	textFile, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(textFile), "\n"), nil
}

func capitalize(words string) string {
	result := strings.ToUpper(string(words[0])) + strings.ToLower(string(words[1:]))
	return result
}

func isVowel(words string) bool {
	if len(words) == 0 {
		return false
	}
	switch words[0] {
	case 'a', 'e', 'i', 'o', 'u', 'h', 'A', 'E', 'I', 'O', 'U', 'H':
		return true
	default:
		return false
	}
}
