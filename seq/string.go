package seq

import "github.com/kulics/gollection/option"

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

func (a *stringIterator) Next() option.Option[rune] {
	if a.index < len(a.source)-1 {
		a.index++
		return option.Some(a.source[a.index])
	}
	return option.None[rune]()
}
