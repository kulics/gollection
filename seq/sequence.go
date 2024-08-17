package seq

// Sequence can be obtained iteratively through Iterator, and each iterator should be independent.
type Sequence[T any] interface {
	Iterator() Iterator[T]
}
