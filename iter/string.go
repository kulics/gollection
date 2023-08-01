package iter

import "github.com/kulics/gollection/util"

// Collection is implemented via String, which is isomorphic to the built-in string.
type String string

func (a String) Iterator() Iterator[rune] {
	return &stringIterator{-1, []rune(a)}
}

func (a String) Count() int {
	return len(a)
}

type stringIterator struct {
	index  int
	source []rune
}

func (a *stringIterator) Next() util.Opt[rune] {
	if a.index < len(a.source)-1 {
		a.index++
		return util.Some(a.source[a.index])
	}
	return util.None[rune]()
}
