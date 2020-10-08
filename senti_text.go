package govader

import (
	"strings"
)

// SentiText holds sentiment-relevant string-level properties of input text.
type SentiText struct {
	Text                   string
	WordsAndEmoticons      []string
	WordsAndEmoticonsLower []string
	IsCapDiff              bool
}

// NewSentiText initliazes and returns a fleshed-out SentiText.
func NewSentiText(text string, pr *PythonesqueRegex) *SentiText {
	var sit SentiText
	sit.Text = text
	tokenList := strings.Split(text, " ")
	sit.WordsAndEmoticons = make([]string, 0)
	for _, token := range tokenList {
		strippedToken := pr.stripPunctuationIfWord(token)
		sit.WordsAndEmoticons = append(sit.WordsAndEmoticons, strippedToken)
	}
	sit.WordsAndEmoticonsLower = stringSliceToLower(sit.WordsAndEmoticons)
	sit.IsCapDiff = pr.allcapDifferential(sit.WordsAndEmoticons)
	return &sit
}

// eof
