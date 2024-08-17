package seq

import (
	"testing"
)

func TestTerminal(t *testing.T) {
	var datas = Slice[int]([]int{1, 2, 3, 4, 5})
	if !Contains[int](3, datas) {
		t.Fatal("Contains error")
	}
	if Sum[int](datas) != 15 {
		t.Fatal("Sum error")
	}
	if Product[int](datas) != 120 {
		t.Fatal("Product error")
	}
	if Average[int](datas) != 3 {
		t.Fatal("Average error")
	}
	if Count[int](datas) != 5 {
		t.Fatal("Count error")
	}
	if Max[int](datas).OrPanic() != 5 {
		t.Fatal("Max error")
	}
	if Min[int](datas).OrPanic() != 1 {
		t.Fatal("Min error")
	}
	var sum int
	testSum := func(i int) {
		sum += i
	}
	ForEach[int](testSum, datas)
	if sum != 15 {
		t.Fatal("ForEach error")
	}
	testMatch := func(i int) bool {
		return i > 0
	}
	if !AllMatch[int](testMatch, datas) {
		t.Fatal("AllMatch error")
	}
	if NoneMatch[int](testMatch, datas) {
		t.Fatal("NoneMatch error")
	}
	if !AnyMatch[int](testMatch, datas) {
		t.Fatal("AnyMatch error")
	}
	if First[int](datas).OrPanic() != 1 {
		t.Fatal("First error")
	}
	if Last[int](datas).OrPanic() != 5 {
		t.Fatal("Last error")
	}
	if At[int](2, datas).OrPanic() != 3 {
		t.Fatal("At error")
	}
	if Reduce[int](func(r int, t int) int {
		return r + t
	}, datas).OrPanic() != 15 {
		t.Fatal("Reduce error")
	}
	if Fold[int](0, func(r int, t int) int {
		return r + t
	}, datas) != 15 {
		t.Fatal("Fold error")
	}
}
