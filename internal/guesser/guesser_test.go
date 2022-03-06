package guesser

import "testing"

func allows(a string, c []Correctness, b string) bool {
	r := GuessResult{a, c}
	return !r.Matches(b)
}

func disallows(a string, c []Correctness, b string) bool {
	return !allows(a, c, b)
}

func TestGuessResult(t *testing.T) {
	if allows("abcde", []Correctness{C, C, C, C, C}, "abcde") {
		t.Error()
	}
	if disallows("abcdf", []Correctness{C, C, C, C, C}, "abcde") {
		t.Error()
	}

	if allows("abcde", []Correctness{W, W, W, W, W}, "fghij") {
		t.Error()
	}
	if allows("abcde", []Correctness{M, M, M, M, M}, "eabcd") {
		t.Error()
	}
	if allows("abcde", []Correctness{M, M, M, M, M}, "eabcd") {
		t.Error()
	}
	if allows("baaaa", []Correctness{W, C, M, W, W}, "aaccc") {
		t.Error()
	}
	if disallows("baaaa", []Correctness{W, C, M, W, W}, "caacc") {
		t.Error()
	}
	if disallows("tares", []Correctness{W, M, M, W, W}, "brink") {
		t.Error()
	}
	if disallows("aaabb", []Correctness{C, M, W, W, W}, "accaa") {
		t.Error()
	}
	if disallows("abcde", []Correctness{W, W, W, W, W}, "bcdea") {
		t.Error()
	}
}
