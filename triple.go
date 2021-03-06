package gollection

// Constructing a Triple with three parameters.
func TripleOf[T1 any, T2 any, T3 any](f T1, s T2, t T3) Triple[T1, T2, T3] {
	return Triple[T1, T2, T3]{f, s, t}
}

// Triple is a combination type containing three elements.
type Triple[T1 any, T2 any, T3 any] struct {
	First  T1
	Second T2
	Third  T3
}

// Get can use go's customary deconstructed Triple,
// which is used like the built-in map.
func (a Triple[T1, T2, T3]) Get() (T1, T2, T3) {
	return a.First, a.Second, a.Third
}
