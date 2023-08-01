package util

// Constructing a Pair with two parameters.
func PairOf[T1 any, T2 any](f T1, s T2) Pair[T1, T2] {
	return Pair[T1, T2]{f, s}
}

// Pair is a combination type containing two elements.
type Pair[T1 any, T2 any] struct {
	First  T1
	Second T2
}

// Val can use go's customary deconstructed Pair,
// which is used like the built-in map.
func (a Pair[T1, T2]) Val() (T1, T2) {
	return a.First, a.Second
}
