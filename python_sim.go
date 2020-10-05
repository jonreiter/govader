package govader

import (
	"regexp"
	"strings"
)

// PythonesqueRegex holds all the regex needed to replicate
// python string manipulaton behaviors
type PythonesqueRegex struct {
	LowerRegex        *regexp.Regexp
	UpperRegex        *regexp.Regexp
	PunctuationRegex  *regexp.Regexp
	PunctuationString string
}

// NewPythonesqueRegex builds a new set of regex
func NewPythonesqueRegex() *PythonesqueRegex {
	var pr PythonesqueRegex
	pr.LowerRegex, _ = regexp.Compile("[a-z]+")
	pr.UpperRegex, _ = regexp.Compile("[A-Z]+")
	pr.PunctuationRegex, _ = regexp.Compile("[^a-zA-Z0-9]+")
	pr.PunctuationString = `"!#$%&\'()*+,-./:;<=>?@[\\]^_` + "`{|}~"
	return &pr
}

// this needs to implement pythons toupper
// only the presence of a lowercase character flips it to false
func (pr *PythonesqueRegex) stringIsUpper(s string) bool {
	hasLower := pr.LowerRegex.MatchString(s)
	if hasLower {
		return false
	}
	return pr.UpperRegex.MatchString(s)
}

// strips only leading and trailing punctuation
func (pr *PythonesqueRegex) stripPunctuationIfWord(text string) string {
	strippedText := strings.Trim(text, pr.PunctuationString)

	if len(strippedText) < 3 {
		return text
	}
	return strippedText
}

// Check whether just some words in the input are ALL CAPS
// words: The words to inspect
// returns: true if some but not all items in words are ALL CAPS
func (pr *PythonesqueRegex) allcapDifferential(words []string) bool {
	allcapWords := 0
	for _, word := range words {
		if pr.stringIsUpper(word) {
			allcapWords++
		}
	}
	capDifferential := len(words) - allcapWords
	if 0 < capDifferential && capDifferential < len(words) {
		return true
	}
	return false
}

// eof
