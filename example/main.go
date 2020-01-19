package main

import (
	"fmt"

	"github.com/jonreiter/govader"
)

func main() {
	sia := govader.NewSentimentIntensityAnalyzer()
	exampleSentences := []string{"VADER is smart, handsome, and funny.",
		"VADER is smart, handsome, and funny!",
		"VADER is very smart, handsome, and funny.",
		"VADER is VERY SMART, handsome, and FUNNY.",
		"VADER is VERY SMART, handsome, and FUNNY!!!",
		"VADER is VERY SMART, uber handsome, and FRIGGIN FUNNY!!!",
		"VADER is not smart, handsome, nor funny.",
		"The book was good.",
		"At least it isn't a horrible book.",
		"The book was only kind of good.",
		"The plot was good, but the characters are uncompelling and the dialog is not great.",
		"Today SUX!",
		"Today only kinda sux! But I'll get by, lol",
		"Make sure you :) or :D today!",
		"Catch utf-8 emoji such as such as 💘 and 💋 and 😁",
		"Not bad at all",
	}
	for _, text := range exampleSentences[0:] {
		scores := sia.PolarityScores(text)
		fmt.Println("for line:", text, "score:", scores)
	}
}
