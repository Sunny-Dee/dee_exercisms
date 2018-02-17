package bob

import (
	"regexp"
	"strings"
)

// Hey returns what Bob would respond
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	switch {
	case remark == "":
		return "Fine. Be that way!"
	case remark[len(remark)-1] == '?':
		if containsLetters(remark) && allCaps(remark) {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	case remark[len(remark)-1] != '?' && containsLetters(remark) && allCaps(remark):
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}

func containsLetters(remark string) bool {
	match, _ := regexp.MatchString("[A-Za-z]+", remark)
	return match
}

func allCaps(remark string) bool {
	caps := strings.ToUpper(remark)

	return caps == remark
}
