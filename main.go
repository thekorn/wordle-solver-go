package main

import (
	_ "embed"
	"strings"
)

//go:embed data/answers.txt
var GAMES []byte

func play[G any](answer string, guesser G) {
	// play six round where it invokes the guesser each round
}

type Correctness int

const (
	// Green
	Correct = iota
	// Yellow
	Misplaces
	// Gray
	Wrong
)

type Guess struct {
	Word string
	Mask [5]Correctness
}

type Guesser interface {
	Guess(history []Guess) string
}

type NaiveGuesser struct{}

func (n *NaiveGuesser) Guess(history []Guess) string {
	return ""
}

func MakeNaiveGuesser() NaiveGuesser {
	return NaiveGuesser{}
}

func main() {
	guesser := MakeNaiveGuesser()
	for _, answer := range strings.Split(string(GAMES), "\n") {
		if answer == "" {
			continue
		}
		play(answer, guesser)
	}
}
