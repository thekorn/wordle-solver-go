package data

import (
	_ "embed"
)

//go:embed data/answers.txt
var GAMES []byte

//go:embed data/dictionary.txt
var DICTIONARY []byte
