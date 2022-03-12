package tuple

func TripleOf[T1 any, T2 any, T3 any](f T1, s T2, t T3) Triple[T1, T2, T3] {
	return Triple[T1, T2, T3]{f, s, t}
}

type Triple[T1 any, T2 any, T3 any] struct {
	First  T1
	Second T2
	Third  T3
}

func (a Triple[T1, T2, T3]) Get() (T1, T2, T3) {
	return a.First, a.Second, a.Third
}
