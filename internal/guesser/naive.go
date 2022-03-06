package guesser

type NaiveGuesser struct{}

func (n NaiveGuesser) Guess(history []GuessResult) string {
	return ""
}

func MakeNaiveGuesser() NaiveGuesser {
	return NaiveGuesser{}
}
