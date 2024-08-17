package dict

type Entry[T1 any, T2 any] struct {
	Key   T1
	Value T2
}
