// Package vader implements the vader sentiment analysis algorithm
// see https://github.com/cjhutto/vaderSentiment
package govader

import (
	"math"
	"strings"
)

func negated(inputWords []string, includeNT bool, negateList []string) bool {
	//    Determine if input contains negation words
	for _, x := range inputWords {
		if inStringSlice(negateList, x) {
			return true
		}
	}
	if includeNT {
		for _, w := range inputWords {
			if strings.Contains(w, "n't") {
				return true
			}
		}
	}
	return false
}

func normalize(score, alpha float64) float64 {
	//    Normalize the score to be between -1 and 1 using an alpha that
	//    approximates the max expected value
	normScore := score / math.Sqrt((score*score)+alpha)
	if normScore < -1.0 {
		return -1.0
	} else if normScore > 1.0 {
		return 1.0
	}
	return normScore
}

func normalizeDefault(score float64) float64 {
	return normalize(score, alphaDefault)
}

func (tc *TermConstants) scalarIncDec(word, wordLower string, valence float64, isCapDiff bool) float64 {
	//    Check if the preceding words increase, decrease, or negate/nullify the
	//    valence
	scalar := 0.0
	if inStringMap(tc.BoosterDict, wordLower) {
		scalar = tc.BoosterDict[wordLower]
		if valence < 0 {
			scalar *= -1
		}
		// check if booster/dampener word is in ALLCAPS (while others aren't)
		if tc.Regex.stringIsUpper(word) && isCapDiff {
			if valence > 0 {
				scalar += cINCR
			} else {
				scalar -= cINCR
			}
		}
	}
	return scalar
}

func siftSentimentScores(sentiments []float64) (float64, float64, int) {
	posSum := 0.0
	negSum := 0.0
	neuCount := 0
	for _, v := range sentiments {
		if v > 0 {
			posSum += v + 1
		}
		if v < 0 {
			negSum += v - 1
		}
		if v == 0 {
			neuCount++
		}
	}

	return posSum, negSum, neuCount
}

func punctuationEmphasis(text string) float64 {
	epAmplifier := amplifyEP(text)
	qmAmplifier := amplifyQM(text)
	return (epAmplifier + qmAmplifier)
}

func amplifyEP(text string) float64 {
	epCount := strings.Count(text, "!")
	if epCount > 4 {
		epCount = 4
	}
	epAmplifier := float64(epCount) * 0.292
	return epAmplifier
}

func amplifyQM(text string) float64 {
	qmCount := strings.Count(text, "?")
	if qmCount > 1 {
		if qmCount <= 3 {
			return float64(qmCount) * 0.18
		}
		return 0.96
	}
	return 0.0
}

func negationCheck(valence float64, wordsAndEmoticonsLower []string, starti, i int, negateList []string) float64 {
	newValence := valence
	if starti == 0 {
		if negated([]string{wordsAndEmoticonsLower[i-(starti+1)]}, true, negateList) {
			newValence = newValence * nSCALAR
		}
	}
	if starti == 1 {
		if wordsAndEmoticonsLower[i-2] == "never" &&
			(wordsAndEmoticonsLower[i-1] == "so" || wordsAndEmoticonsLower[i-1] == "this") {
			newValence = valence * 1.25
		} else if wordsAndEmoticonsLower[i-2] == "without" &&
			wordsAndEmoticonsLower[i-1] == "doubt" {
			newValence = valence
		} else if negated([]string{wordsAndEmoticonsLower[i-(starti+1)]}, true, negateList) {
			newValence = valence * nSCALAR
		}
	}
	if starti == 2 {
		if wordsAndEmoticonsLower[i-3] == "never" &&
			((wordsAndEmoticonsLower[i-2] == "so" || wordsAndEmoticonsLower[i-2] == "this") ||
				(wordsAndEmoticonsLower[i-1] == "so" || wordsAndEmoticonsLower[i-1] == "this")) {
			newValence = valence * 1.25
		} else if wordsAndEmoticonsLower[i-3] == "without" &&
			(wordsAndEmoticonsLower[i-2] == "doubt" || wordsAndEmoticonsLower[i-1] == "doubt") {
			newValence = valence
		} else if negated([]string{wordsAndEmoticonsLower[i-(starti+1)]}, true, negateList) {
			newValence = valence * nSCALAR
		}
	}
	return newValence
}

func butCheck(wordsAndEmoticonsLower []string, sentiments []float64) []float64 {
	// check for modification in sentiment due to contrastive conjunction 'but'
	if inStringSlice(wordsAndEmoticonsLower, "but") {
		bi := firstIndexOfStringInSlice(wordsAndEmoticonsLower, "but")
		for i := range sentiments {
			if i < bi {
				sentiments[i] = 0.5 * sentiments[i]
			}
			if i > bi {
				sentiments[i] = 1.5 * sentiments[i]
			}
		}
	}
	return sentiments
}

// eof
