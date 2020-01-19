package govader

import (
	"regexp"
	"strings"
)

// PythonesqueRegex holds all the regex needed to replicate
// python string manipulaton behaviors
type PythonesqueRegex struct {
	LowerRegex       *regexp.Regexp
	UpperRegex       *regexp.Regexp
	PunctuationRegex *regexp.Regexp
}

// NewPythonesqueRegex builds a new set of regex
func NewPythonesqueRegex() *PythonesqueRegex {
	var pr PythonesqueRegex
	pr.LowerRegex, _ = regexp.Compile("[a-z]+")
	pr.UpperRegex, _ = regexp.Compile("[A-Z]+")
	pr.PunctuationRegex, _ = regexp.Compile("[^a-zA-Z0-9]+")
	return &pr
}

// FIXME - this needs to implement pythons toupper
// only the presence of a lowercase character flips it to false
func (pr *PythonesqueRegex) stringIsUpper(s string) bool {
	hasLower := pr.LowerRegex.MatchString(s)
	hasUpper := pr.UpperRegex.MatchString(s)
	if hasUpper && !hasLower {
		return true
	}
	return false
}

func (pr *PythonesqueRegex) stripPunctuation(text string) string {
	reg := pr.PunctuationRegex
	return reg.ReplaceAllString(text, "")
}

func stringSliceToLower(stringsList []string) []string {
	newStrings := make([]string, len(stringsList))
	for i, v := range stringsList {
		newStrings[i] = strings.ToLower(v)
	}
	return newStrings
}

// eof
