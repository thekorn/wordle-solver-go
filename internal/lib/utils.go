package lib

import "fmt"

type StringTuple struct {
	a, b rune
}

func (st StringTuple) IsEqual() bool {
	return st.a == st.b
}

func StringZip(a string, b string) []StringTuple {
	la := len(a)
	if la != len(b) {
		panic(fmt.Sprintf("string a and b must have the same length, got %d != %d", len(a), len(b)))
	}
	r := make([]StringTuple, la, la)
	for i, e := range a {
		r[i] = StringTuple{e, rune(b[i])}
	}
	return r
}
