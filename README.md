# GoVader
GoVader: Vader sentiment analysis in Go

This is a port of https://github.com/cjhutto/vaderSentiment from
Python to Go.

There are tests which check it gives the same answers as the original package. And there are docs at https://godoc.org/github.com/jonreiter/govader.

Usage:
```go
import (
    "fmt"
    "github.com/jonreiter/govader"
)

sentimentAnalyzer := govader.NewSentimentIntensityAnalyzer()
sentiment := sentimentAnalyzer.PolarityScores("Usage is similar to all the other ports.")

fmt.Println("Compound score:",sentiment.Compound)
fmt.Println("Positive score:",sentiment.Positive)
fmt.Println("Neutral score:",sentiment.Neutral)
fmt.Println("Negative score:",sentiment.Negative)

```

