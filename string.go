package gollection

func ToString(a string) String {
	return String(a)
}

func ToStringIter(a string) Iterator[rune] {
	return &stringIterator{-1, []rune(a)}
}

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

func (a *stringIterator) Iter() Iterator[rune] {
	return a
}
