package game

import (
	"fmt"
	"strings"

	"github.com/thekorn/wordle-solver-go/internal/data"
	"github.com/thekorn/wordle-solver-go/internal/guesser"
	"github.com/thekorn/wordle-solver-go/internal/lib"
)

const COUNT = 32

func InitDictionarySet() lib.StringSet {
	s := strings.Split(string(data.DICTIONARY), "\n")
	x := make([]string, len(s))
	for _, e := range s {
		if e == "" {
			continue
		}
		w := strings.Fields(e)
		x = append(x, w[0])
	}
	d := lib.MakeStringSetFromSlice(x, true)
	return d
}

type Game[G guesser.Guesser] interface {
	Play(answer string, guesser G) int
}

type Wordle struct {
	dictionary lib.StringSet
}

func (w Wordle) Play(answer string, g guesser.Guesser) int {
	history := make([]guesser.GuessResult, COUNT)
	for i := 1; i <= COUNT; i++ {
		guess := g.Guess(history)
		if guess == answer {
			return i
		}
		if !w.dictionary.Has(guess) {
			panic(fmt.Sprintf("guess '%s' is not in the dictionary", guess))
		}
		correctness := guesser.ComputeCorrectness(answer, guess)
		history = append(history, guesser.GuessResult{
			Word: guess,
			Mask: correctness,
		})
	}
	panic(fmt.Sprintf("unable to solve the wordle after %d iterations", COUNT))
}

func MakeWordle() Wordle {
	dictionary := InitDictionarySet()
	return Wordle{dictionary: dictionary}
}
