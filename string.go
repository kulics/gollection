package gollection

// Converts a built-in string to a String, which does not copy elements.
func ToString(a string) String {
	return String(a)
}

// Construct an iterator using the built-in string.
func ToStringIter(a string) Iterator[rune] {
	return &stringIterator{-1, []rune(a)}
}

// Collection is implemented via String, which is isomorphic to the built-in string.
type String string

func (a String) Iter() Iterator[rune] {
	return &stringIterator{-1, []rune(a)}
}

func (a String) Size() int {
	return len(a)
}

func (a String) IsEmpty() bool {
	return len(a) == 0
}

func (a String) ToSlice() []rune {
	return []rune(a)
}

type stringIterator struct {
	index  int
	source []rune
}

func (a *stringIterator) Next() Option[rune] {
	if a.index < len(a.source)-1 {
		a.index++
		return Some(a.source[a.index])
	}
	return None[rune]()
}
