package iter

import (
	"testing"
)

func TestTransform(t *testing.T) {
	show := func(i int) {
		println(i)
	}
	even := func(i int) bool {
		return i%2 == 0
	}
	square := func(i int) int {
		return i * i
	}
	ForEach(show, Map(square, Filter(even, Slice[int]([]int{1, 2, 3, 4, 5, 6, 7}).Iterator())))
}
