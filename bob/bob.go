// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
	"unicode"
)

var allNonLetters = `\W`
var nonLettersAndDigits = `\W|\d`
var fine = "Fine. Be that way!"

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.Trim(remark, " ")
	if remark == "" {
		return fine
	}
	yelling := isYelling(remark)
	askingQuestion := isAskingQuestion(remark)
	if yelling && askingQuestion {
		return "Calm down, I know what I'm doing!"
	}
	if yelling {
		return "Whoa, chill out!"
	}
	if askingQuestion {
		return "Sure."
	}
	if isSayingNothing(remark) {
		return fine
	}

	return "Whatever."
}

func isAskingQuestion(sentence string) bool {
	return (len(sentence) - 1) == strings.LastIndex(sentence, "?")
}

func isYelling(sentence string) bool {
	re := regexp.MustCompile(nonLettersAndDigits)
	updatedWord := re.ReplaceAllString(sentence, "")
	if len(updatedWord) == 0 {
		return false
	}
	for _, c := range sentence {
		if unicode.IsLetter(c) && unicode.IsLower(c) {
			return false
		}
	}
	return true
}

func isSayingNothing(sentence string) bool {
	re := regexp.MustCompile(allNonLetters)
	updatedWord := re.ReplaceAllString(sentence, "")
	if updatedWord == "" {
		return true
	}
	return false
}
