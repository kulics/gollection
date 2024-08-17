package seq

// Sequence's extended interfaces, can provide more information to optimize performance.
type Collection[T any] interface {
	Sequence[T]

	Count() int
}

func Equals[T comparable](l Collection[T], r Collection[T]) bool {
	if l.Count() != r.Count() {
		return false
	}
	var lIter = l.Iterator()
	var rIter = r.Iterator()
	for {
		if v1, ok1 := lIter.Next().Val(); ok1 {
			if v2, _ := rIter.Next().Val(); v1 != v2 {
				return false
			}
		} else {
			break
		}
	}
	return true
}
