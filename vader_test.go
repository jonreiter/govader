package govader_test

import (
	"testing"

	"github.com/jonreiter/govader"
	"gonum.org/v1/gonum/mat"
)

type PolarityTestCase struct {
	Text   string
	Scores govader.Sentiment
}

// the python reference implementation rounds scores to 3 or 4 decimal
// places.  so we test to that tolerance.
const matchEpsilon = 0.5e-3

func scoresMatch(expectedScore, realizedScore govader.Sentiment) bool {
	eVec := mat.NewVecDense(4, []float64{expectedScore.Compound, expectedScore.Negative, expectedScore.Neutral, expectedScore.Positive})
	rVec := mat.NewVecDense(4, []float64{realizedScore.Compound, realizedScore.Negative, realizedScore.Neutral, realizedScore.Positive})
	return mat.EqualApprox(eVec, rVec, matchEpsilon)
}

// GetExamples returns the examples, with scores, from the reference
// python implementation
func GetExamples() []PolarityTestCase {
	examples := []PolarityTestCase{
		{Text: "VADER is smart, handsome, and funny.", Scores: govader.Sentiment{Negative: 0, Neutral: 0.254, Positive: 0.746, Compound: 0.8316}},
		{Text: "VADER is smart, handsome, and funny!", Scores: govader.Sentiment{Negative: 0.0, Neutral: 0.248, Positive: 0.752, Compound: 0.8439}},
		{Text: "VADER is very smart, handsome, and funny.", Scores: govader.Sentiment{Negative: 0.0, Neutral: 0.299, Positive: 0.701, Compound: 0.8545}},
		{Text: "VADER is VERY SMART, handsome, and FUNNY.", Scores: govader.Sentiment{Negative: 0.0, Neutral: 0.246, Positive: 0.754, Compound: 0.9227}},
		{Text: "VADER is VERY SMART, handsome, and FUNNY!!!", Scores: govader.Sentiment{Negative: 0.0, Neutral: 0.233, Positive: 0.767, Compound: 0.9342}},
		{Text: "VADER is VERY SMART, uber handsome, and FRIGGIN FUNNY!!!", Scores: govader.Sentiment{Negative: 0.0, Neutral: 0.294, Positive: 0.706, Compound: 0.9469}},
		{Text: "VADER is not smart, handsome, nor funny.", Scores: govader.Sentiment{Negative: 0.646, Neutral: 0.354, Positive: 0.0, Compound: -0.7424}},
		{Text: "The book was good.", Scores: govader.Sentiment{Negative: 0.0, Neutral: 0.508, Positive: 0.492, Compound: 0.4404}},
		{Text: "At least it isn't a horrible book.", Scores: govader.Sentiment{Negative: 0.0, Neutral: 0.678, Positive: 0.322, Compound: 0.431}},
		{Text: "The book was only kind of good.", Scores: govader.Sentiment{Negative: 0.0, Neutral: 0.697, Positive: 0.303, Compound: 0.3832}},
		{Text: "The plot was good, but the characters are uncompelling and the dialog is not great.", Scores: govader.Sentiment{Negative: 0.327, Neutral: 0.579, Positive: 0.094, Compound: -0.7042}},
		{Text: "Today SUX!", Scores: govader.Sentiment{Negative: 0.779, Neutral: 0.221, Positive: 0.0, Compound: -0.5461}},
		{Text: "Today only kinda sux! But I'll get by, lol", Scores: govader.Sentiment{Negative: 0.127, Neutral: 0.556, Positive: 0.317, Compound: 0.5249}},
		{Text: "Make sure you :) or :D today!", Scores: govader.Sentiment{Negative: 0.0, Neutral: 0.294, Positive: 0.706, Compound: 0.8633}},
		{Text: "Catch utf-8 emoji such as such as 💘 and 💋 and 😁", Scores: govader.Sentiment{Negative: 0.0, Neutral: 0.746, Positive: 0.254, Compound: 0.7003}},
		{Text: "Not bad at all", Scores: govader.Sentiment{Negative: 0.0, Neutral: 0.513, Positive: 0.487, Compound: 0.431}},
		// from here all are additional tests to increase coverage. scores are checked against original values here not against the original implementation
		{Text: "Not GREATLY bad at all!!!!!", Scores: govader.Sentiment{Negative: 0.0, Neutral: 0.4557, Positive: 0.5442, Compound: 0.6982}},
		{Text: "Not GREATLY bad at all??", Scores: govader.Sentiment{Negative: 0.0, Neutral: 0.5019, Positive: 0.4980, Compound: 0.6084}},
		{Text: "Not GREATLY bad at all????", Scores: govader.Sentiment{Negative: 0.0, Neutral: 0.4667, Positive: 0.5332, Compound: 0.6777}},
		{Text: "Not least GREATLY bad at all????", Scores: govader.Sentiment{Negative: 0.0, Neutral: 0.5225, Positive: 0.4775, Compound: 0.6777}},
		{Text: "Not GREATLY least bad at all????", Scores: govader.Sentiment{Negative: 0.4358, Neutral: 0.5642, Positive: 0.0, Compound: -0.5944}},
		{Text: "Catch utf-8 emoji such as such as💘 and 💋 and 😁", Scores: govader.Sentiment{Negative: 0.0, Neutral: 0.746, Positive: 0.254, Compound: 0.7003}},
		{Text: "least bad at all????", Scores: govader.Sentiment{Negative: 0.0, Neutral: 0.4405, Positive: 0.5595, Compound: 0.5873}},
		{Text: "No not GREATLY least bad at all????", Scores: govader.Sentiment{Negative: 0.5480, Neutral: 0.4520, Positive: 0.0, Compound: -0.7238}},
		{Text: "The book was only kind of no good.", Scores: govader.Sentiment{Negative: 0.3813, Neutral: 0.6186, Positive: 0.0, Compound: -0.4017}},
		{Text: "The book was only kind of bad ass good.", Scores: govader.Sentiment{Negative: 0.4618, Neutral: 0.3800, Positive: 0.1583, Compound: -0.6997}},
		{Text: "The book was only kind of never so bad ass good.", Scores: govader.Sentiment{Negative: 0.2568, Neutral: 0.4574, Positive: 0.2859, Compound: -0.1258}},
		{Text: "The book was only kind of without doubt bad ass good.", Scores: govader.Sentiment{Negative: 0.1943, Neutral: 0.3887, Positive: 0.4170, Compound: 0.4606}},
		{Text: "The book was only kind of badn't ass good.", Scores: govader.Sentiment{Negative: 0.1963, Neutral: 0.5711, Positive: 0.2325, Compound: 0.1139}},
	}
	return examples
}

func TestPolarityScores(t *testing.T) {
	sia := govader.NewSentimentIntensityAnalyzer()
	for _, testCase := range GetExamples() {
		realizedScore := sia.PolarityScores(testCase.Text)
		if !scoresMatch(testCase.Scores, realizedScore) {
			t.Error("score mismatch on:", testCase, "vs", realizedScore)
		}
	}
}

func BenchmarkPolarityScores(b *testing.B) {
	examples := GetExamples()
	sia := govader.NewSentimentIntensityAnalyzer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, e := range examples {
			_ = sia.PolarityScores(e.Text)
		}
	}
}

func BenchmarkPolarityScoresLarge(b *testing.B) {
	examples := GetExamples()
	bigText := ""
	for i := 0; i < 10; i++ {
		for _, example := range examples {
			bigText = bigText + example.Text
		}
	}
	sia := govader.NewSentimentIntensityAnalyzer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sia.PolarityScores(bigText)
	}
}

// eof
