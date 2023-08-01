package iter

import "github.com/kulics/gollection/util"

// Iterator can be obtained iteratively through Iter, and each iterator should be independent.
type Iterable[T any] interface {
	Iterator() Iterator[T]
}

// By implementing Next you can perform iterations and end them when the return value is None.
type Iterator[T any] interface {
	Next() util.Opt[T]
}

// Iterable's extended interfaces, can provide more information to optimize performance.
type Collection[T any] interface {
	Iterable[T]

	Count() int
}

const OutOfBounds = "out of bounds"
