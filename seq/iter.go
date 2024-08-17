package seq

import "github.com/kulics/gollection/option"

// By implementing Next you can perform iterations and end them when the return value is None.
type Iterator[T any] interface {
	Next() option.Option[T]
}

const OutOfBounds = "out of bounds"
