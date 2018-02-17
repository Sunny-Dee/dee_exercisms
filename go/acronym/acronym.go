package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate creates an acronym based on s
func Abbreviate(s string) string {
	r := regexp.MustCompile("[[:space:][:punct:]]+")
	words := r.Split(s, -1)
	acronym := ""

	for _, word := range words {
		acronym += strings.ToUpper(string(word[0]))
	}

	return acronym
}
