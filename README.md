# GoVader

[![GoDoc](https://godoc.org/github.com/jonreiter/govader?status.svg)](https://godoc.org/github.com/jonreiter/govader)  [![Go Report](https://goreportcard.com/badge/github.com/jonreiter/govader)](https://goreportcard.com/badge/github.com/jonreiter/govader)

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

analyzer := govader.NewSentimentIntensityAnalyzer()
sentiment := analyzer.PolarityScores("Usage is similar to all the other ports.")

fmt.Println("Compound score:", sentiment.Compound)
fmt.Println("Positive score:", sentiment.Positive)
fmt.Println("Neutral score:", sentiment.Neutral)
fmt.Println("Negative score:", sentiment.Negative)

```

