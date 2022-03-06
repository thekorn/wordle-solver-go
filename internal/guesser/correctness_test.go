package guesser

import (
	"testing"
)

func TestMakeCorrectnessPatterns(t *testing.T) {
	c := 0
	for range MakeCorrectnessPatterns() {
		c += 1
	}
	if c != 243 {
		t.Errorf("correctness patterns must be the cartesian product")
	}
}

func assertCorrectness(t *testing.T, a []Correctness, b []Correctness) {
	if len(a) != 5 || len(b) != 5 {
		t.Error("correctness results must both be of size 5")
	}
	if !(a[0] == b[0] && a[1] == b[1] && a[2] == b[2] && a[3] == b[3] && a[4] == b[4]) {
		t.Error("not the same correctness")
	}
}

func TestComputeCorrectness(t *testing.T) {
	assertCorrectness(t, ComputeCorrectness("abcde", "abcde"), []Correctness{C, C, C, C, C})
	assertCorrectness(t, ComputeCorrectness("abcde", "fghij"), []Correctness{W, W, W, W, W})
	assertCorrectness(t, ComputeCorrectness("abcde", "eabcd"), []Correctness{M, M, M, M, M})
	assertCorrectness(t, ComputeCorrectness("aabbb", "aaccc"), []Correctness{C, C, W, W, W})
	assertCorrectness(t, ComputeCorrectness("aabbb", "ccaac"), []Correctness{W, W, M, M, W})
	assertCorrectness(t, ComputeCorrectness("aabbb", "caacc"), []Correctness{W, C, M, W, W})
	assertCorrectness(t, ComputeCorrectness("azzaz", "aaabb"), []Correctness{C, M, W, W, W})
	assertCorrectness(t, ComputeCorrectness("baccc", "aaddd"), []Correctness{W, C, W, W, W})
	assertCorrectness(t, ComputeCorrectness("abcde", "aacde"), []Correctness{C, W, C, C, C})

}
