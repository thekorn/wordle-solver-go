package guesser

type GuessResult struct {
	Word string
	Mask []Correctness
}

func (gr GuessResult) Matches(word string) bool {
	a := ComputeCorrectness(word, gr.Word)
	b := gr.Mask

	if len(a) != 5 || len(b) != 5 {
		panic("correctness results must both be of size 5")
	}
	return a[0] == b[0] && a[1] == b[1] && a[2] == b[2] && a[3] == b[3] && a[4] == b[4]
}

type Guesser interface {
	Guess(history []GuessResult) string
}
