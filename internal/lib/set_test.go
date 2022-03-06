package lib

import "testing"

func TestCreateEmptySet(t *testing.T) {
	s := StringSet{}
	if len(s) != 0 {
		t.Errorf("empty StringSet should not contain anything")
	}
}

func TestMakeFromSlice(t *testing.T) {
	s := MakeStringSetFromSlice([]string{"g", "h", "i"}, true)
	if len(s) != 3 || !s.Has("g") || !s.Has("h") || !s.Has("i") {
		t.Errorf("failed creating StringSet from slice")
	}
}

func TestMakeFromSliceSkipEmpty(t *testing.T) {
	s := MakeStringSetFromSlice([]string{"g", "h", "i", ""}, true)
	if len(s) != 3 || !s.Has("g") || !s.Has("h") || !s.Has("i") {
		t.Errorf("failed creating StringSet from slice, skipping empty string")
	}
	s = MakeStringSetFromSlice([]string{"g", "h", "i", ""}, false)
	if len(s) != 4 || !s.Has("g") || !s.Has("h") || !s.Has("i") || !s.Has("") {
		t.Errorf("failed creating StringSet from slice, not skipping empty string")
	}
}

func TestRemoveFromStringSet(t *testing.T) {
	s := MakeStringSetFromSlice([]string{"g", "h", "i"}, true)
	if len(s) != 3 || !s.Has("g") || !s.Has("h") || !s.Has("i") {
		t.Errorf("failed creating StringSet from slice")
	}
	s.Remove("g")
	if len(s) != 2 || s.Has("g") || !s.Has("h") || !s.Has("i") {
		t.Errorf("failed removing from StringSet")
	}

}
