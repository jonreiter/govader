package govader

import (
	"bufio"
	"bytes"
	"log"
	"strconv"
	"strings"

	"github.com/jonreiter/govader/data"
)

const lexiconAssetName = "rawdata/vader_lexicon.txt"
const emojiAssetName = "rawdata/emoji_utf8_lexicon.txt"

// SentimentIntensityAnalyzer ...
type SentimentIntensityAnalyzer struct {
	Lexicon   map[string]float64
	EmojiDict map[string]string
	Constants *TermConstants
}

// Sentiment encapsulates a single sentiment measure for a statement
type Sentiment struct {
	Negative float64
	Neutral  float64
	Positive float64
	Compound float64
}

func (sia *SentimentIntensityAnalyzer) makeLexDict() {
	sia.Lexicon = make(map[string]float64)
	asset, err := data.Asset(lexiconAssetName)
	if err != nil {
		log.Panic("could not open lexicon data")
	}
	file := bytes.NewReader(asset)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		thisRawLine := scanner.Text()
		thisSplitLine := strings.Split(thisRawLine, "\t")
		word := thisSplitLine[0]
		measure, _ := strconv.ParseFloat(thisSplitLine[1], 64)
		sia.Lexicon[word] = measure
	}
}

func (sia *SentimentIntensityAnalyzer) makeEmojiDict() {
	sia.EmojiDict = make(map[string]string)
	asset, err := data.Asset(emojiAssetName)
	if err != nil {
		log.Panic("could not open emoji data")
	}
	file := bytes.NewReader(asset)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		thisRawLine := scanner.Text()
		thisSplitLine := strings.Split(thisRawLine, "\t")
		word := thisSplitLine[0]
		descr := thisSplitLine[1]
		sia.EmojiDict[word] = descr
	}
}

// PolarityScores ...
// Return a float for sentiment strength based on the input text.
// Positive values are positive valence, negative value are negative
// valence.
func (sia *SentimentIntensityAnalyzer) PolarityScores(text string) Sentiment {
	textNoEmoji := ""
	prevSpace := true
	for _, rune := range text {
		chr := string(rune)
		if inStringStringMap(sia.EmojiDict, chr) {
			description := sia.EmojiDict[chr]
			if !prevSpace {
				textNoEmoji = textNoEmoji + " "
			}
			textNoEmoji = textNoEmoji + description
			prevSpace = false
		} else {
			textNoEmoji = textNoEmoji + chr
			prevSpace = false
			if chr == " " {
				prevSpace = true
			}
		}
	}
	trimmedText := strings.TrimSpace(textNoEmoji)

	sentitext := NewSentiText(trimmedText, sia.Constants.Regex)

	sentiments := make([]float64, 0)
	wordsAndEmoticons := sentitext.WordsAndEmoticons
	wordsAndEmoticonsLower := sentitext.WordsAndEmoticonsLower
	for i, item := range wordsAndEmoticons {
		valence := 0.0
		itemLower := wordsAndEmoticonsLower[i]

		// check for vader_lexicon words that may be used as modifiers or negations
		if inStringMap(sia.Constants.BoosterDict, itemLower) {
			sentiments = append(sentiments, valence)
		} else if i < (len(wordsAndEmoticons)-1) && itemLower == "kind" &&
			wordsAndEmoticonsLower[i+1] == "of" {
			sentiments = append(sentiments, valence)
		} else {
			sentiments = sia.sentimentValence(valence, sentitext, item, i, sentiments)
		}
	}
	sentiments = butCheck(wordsAndEmoticonsLower, sentiments)
	valenceDict := scoreValence(sentiments, trimmedText)

	return valenceDict
}

// sentimentValence ...
func (sia *SentimentIntensityAnalyzer) sentimentValence(valence float64, sit *SentiText, item string, i int, sentiments []float64) []float64 {
	isCapDiff := sit.IsCapDiff
	wordsAndEmoticons := sit.WordsAndEmoticons
	wordsAndEmoticonsLower := sit.WordsAndEmoticonsLower
	itemLower := strings.ToLower(item)

	newValence := valence

	if inStringMap(sia.Lexicon, itemLower) {
		newValence = sia.Lexicon[itemLower]
		if itemLower == "no" && i+1 < len(wordsAndEmoticonsLower) {
			if inStringMap(sia.Lexicon, wordsAndEmoticonsLower[i+1]) {
				newValence = 0
			}
		}
		if (i > 0 && wordsAndEmoticonsLower[i-1] == "no") ||
			(i > 1 && wordsAndEmoticonsLower[i-2] == "no") ||
			(i > 2 && wordsAndEmoticonsLower[i-3] == "no" && inStringSlice([]string{"or", "nor"}, wordsAndEmoticonsLower[i-1])) {
			newValence = sia.Lexicon[itemLower] * nSCALAR
		}

		if sia.Constants.Regex.stringIsUpper(item) && isCapDiff {
			if newValence > 0 {
				newValence += cINCR
			} else {
				newValence -= cINCR
			}
		}

		for startI := range []int{0, 1, 2} {
			if i > startI &&
				!inStringMap(sia.Lexicon, wordsAndEmoticons[i-(startI+1)]) {
				s := sia.Constants.scalarIncDec(wordsAndEmoticons[i-(startI+1)], wordsAndEmoticonsLower[i-(startI+1)], newValence, isCapDiff)
				if startI == 1 && s != 0 {
					s = s * valenceScalarScale1
				}
				if startI == 2 && s != 0 {
					s = s * valenceScalarScale2
				}
				newValence = newValence + s
				newValence = negationCheck(newValence, wordsAndEmoticonsLower, startI, i, sia.Constants.NegateList)
				if startI == 2 {
					newValence = sia.Constants.specialIdiomsCheck(newValence, wordsAndEmoticonsLower, i, sia.Constants.BoosterDict)
				}
			}
		}
		newValence = sia.leastCheck(newValence, wordsAndEmoticons, i)
	}
	sentiments = append(sentiments, newValence)
	return sentiments
}

func (sia *SentimentIntensityAnalyzer) leastCheck(valence float64, wordsAndEmoticonsLower []string, i int) float64 {
	// check for negation case using "least"
	newValence := valence
	if i > 0 &&
		!inStringMap(sia.Lexicon, wordsAndEmoticonsLower[i-1]) &&
		wordsAndEmoticonsLower[i-1] == "least" {
		if wordsAndEmoticonsLower[i-2] != "at" &&
			wordsAndEmoticonsLower[i-2] != "very" {
			newValence = newValence * nSCALAR
		}
	} else if i > 0 &&
		!inStringMap(sia.Lexicon, wordsAndEmoticonsLower[i-1]) &&
		wordsAndEmoticonsLower[i-1] == "least" {
		newValence = newValence * nSCALAR
	}
	return newValence
}

// NewSentimentIntensityAnalyzer ...
func NewSentimentIntensityAnalyzer() *SentimentIntensityAnalyzer {
	var sia SentimentIntensityAnalyzer
	sia.makeLexDict()
	sia.makeEmojiDict()
	sia.Constants = NewTermConstants()
	return &sia
}

// eof
