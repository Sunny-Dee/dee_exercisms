package isogram

import (
	"log"
	"regexp"
	"strings"
)

// IsIsogram determines is a word is an isogram
func IsIsogram(word string) bool {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	word = strings.ToLower(reg.ReplaceAllString(word, ""))

	tally := make(map[byte]struct{})
	for i := 0; i < len(word); i++ {
		letter := word[i]

		if _, ok := tally[letter]; ok {
			return false
		}
		tally[letter] = struct{}{}

	}
	return true
}
