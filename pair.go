package main

type Pair[T1 any, T2 any] struct {
	First  T1
	Second T2
}

func PairOf[T1 any, T2 any](f T1, s T2) Pair[T1, T2] {
	return Pair[T1, T2]{f, s}
}
