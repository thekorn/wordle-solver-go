package guesser

import (
	"fmt"

	"github.com/schwarmco/go-cartesian-product"
	"github.com/thekorn/wordle-solver-go/internal/lib"
)

type Correctness int

const (
	// Green
	Correct Correctness = iota
	// Yellow
	Misplaced
	// Gray
	Wrong
)

// short form
const C = Correct
const W = Wrong
const M = Misplaced

func (c Correctness) String() string {
	switch c {
	case Correct:
		return "Correct"
	case Misplaced:
		return "Misplaced"
	case Wrong:
		return "Wrong"
	}
	panic(fmt.Sprintf("Unknown Correctness '%d'", c))
}

func ComputeCorrectness(answer string, guess string) []Correctness {
	if len(answer) != 5 {
		panic(fmt.Sprintf("answer '%s' is not of length 5", answer))
	}
	if len(guess) != 5 {
		panic(fmt.Sprintf("guess '%s' is not of length 5", guess))
	}

	r := []Correctness{Wrong, Wrong, Wrong, Wrong, Wrong}
	used := []bool{false, false, false, false, false}
	for i, st := range lib.StringZip(answer, guess) {
		if st.IsEqual() {
			r[i] = Correct
			used[i] = true
		}
	}

	for i, g := range guess {
		if r[i] == Correct {
			// already marked as green
			continue
		}
		isCool := false
		for n, v := range answer {
			if g == v && !used[n] {
				used[n] = true
				isCool = true
				break
			}
		}
		if isCool {
			r[i] = Misplaced
		}
	}
	return r
}

func MakeCorrectnessPatterns() chan []interface{} {
	return cartesian.Iter(
		[]interface{}{Correct, Misplaced, Wrong},
		[]interface{}{Correct, Misplaced, Wrong},
		[]interface{}{Correct, Misplaced, Wrong},
		[]interface{}{Correct, Misplaced, Wrong},
		[]interface{}{Correct, Misplaced, Wrong},
	)
}
