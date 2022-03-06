package lib

// https://codetree.dev/golang-implementing-sets/
type StringSet map[string]struct{}

func (d StringSet) Add(s string) {
	d[s] = struct{}{}
}

func (d StringSet) Remove(s string) {
	delete(d, s)
}

func (d StringSet) Has(s string) bool {
	_, ok := d[s]
	return ok
}

func MakeStringSetFromSlice(s []string, skipEmpty bool) StringSet {
	d := StringSet{}
	for _, e := range s {
		if skipEmpty && e == "" {
			continue
		}
		d.Add(e)
	}
	return d
}
