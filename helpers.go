package govader

import (
	"strings"
)

func inStringSlice(slice []string, theString string) bool {
	for _, w := range slice {
		if w == theString {
			return true
		}
	}
	return false
}

func inStringMap(theMap map[string]float64, theString string) bool {
	if _, ok := theMap[theString]; ok {
		return true
	}
	return false
}

func inStringStringMap(theMap map[string]string, theString string) bool {
	if _, ok := theMap[theString]; ok {
		return true
	}
	return false
}

// strips only leading and trailing punctuation
func stripPunctuationIfWord(text string) string {
	cutset := `"!#$%&\'()*+,-./:;<=>?@[\\]^_` + "`{|}~"
	strippedText := strings.Trim(text, cutset)

	if len(strippedText) < 3 {
		return text
	}
	return strippedText
}

func firstIndexOfStringInSlice(slice []string, toFind string) int {
	for i, v := range slice {
		if v == toFind {
			return i
		}
	}
	return -1
}

func firstIndexOfFloatInSlice(slice []float64, toFind float64) int {
	for i, v := range slice {
		if v == toFind {
			return i
		}
	}
	return -1
}

func (pr *PythonesqueRegex) allcapDifferential(words []string) bool {
	//    Check whether just some words in the input are ALL CAPS
	//    :param list words: The words to inspect
	//    :returns: `True` if some but not all items in `words` are ALL CAPS
	isDifferent := false
	allcapWords := 0
	for _, word := range words {
		if pr.stringIsUpper(word) {
			allcapWords++
		}
	}
	capDifferential := len(words) - allcapWords
	if 0 < capDifferential && capDifferential < len(words) {
		isDifferent = true
	}
	return isDifferent
}

// eof
