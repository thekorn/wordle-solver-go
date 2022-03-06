package lib

import "testing"

func shouldPanic(t *testing.T, f func()) {
	defer func() { recover() }()
	f()
	t.Errorf("should have panicked")
}

func TestStringZip(t *testing.T) {
	a := "abc"
	b := "efg"

	v := StringZip(a, b)
	if len(v) != 3 {
		t.Errorf("StringZip result has wrong length")
	}
	if string(v[0].a) != "a" || string(v[0].b) != "e" {
		t.Errorf("first element is wrong")
	}
	if string(v[1].a) != "b" || string(v[1].b) != "f" {
		t.Errorf("second element is wrong")
	}
	if string(v[2].a) != "c" || string(v[2].b) != "g" {
		t.Errorf("third element is wrong")
	}
}

func TestStringZipIsEqual(t *testing.T) {
	a := "abc"
	b := "abc"

	v := StringZip(a, b)
	if len(v) != 3 {
		t.Errorf("StringZip result has wrong length")
	}
	if !v[0].IsEqual() {
		t.Errorf("first elements are not equal")
	}
	if !v[1].IsEqual() {
		t.Errorf("second elements are not equal")
	}
	if !v[2].IsEqual() {
		t.Errorf("third elements are not equal")
	}
}

func TestStringZipSizeMismatch(t *testing.T) {
	a := "abc"
	b := "e"

	sizeMismatchStringZip := func() {
		StringZip(a, b)
	}
	shouldPanic(t, sizeMismatchStringZip)
}
