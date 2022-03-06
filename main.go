package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/thekorn/wordle-solver-go/internal/data"
	"github.com/thekorn/wordle-solver-go/internal/game"
	"github.com/thekorn/wordle-solver-go/internal/guesser"
)

func play[G guesser.Guesser](guesser G, max int) {

	w := game.MakeWordle()
	fmt.Printf("game: %s", w)

	for _, answer := range strings.Split(string(data.GAMES), "\n")[0:max] {
		if answer == "" {
			continue
		}
		//w.play(answer, guesser)
	}
}

func initImplementation(s string, max int) guesser.Guesser {
	if s == "naive" {
		return guesser.MakeNaiveGuesser()
	} else {
		panic(fmt.Sprintf("no implementation found for %s", s))
	}
}

func main() {

	implementationPtr := flag.String("implementation", "naive", "implementation to use")
	maxPtr := flag.Int("max", 1, "number of games to run")

	flag.Parse()

	guesser := initImplementation(*implementationPtr, *maxPtr)

	play(guesser, *maxPtr)
}
