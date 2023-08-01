package iter

import (
	"testing"
)

func TestTerminal(t *testing.T) {
	var datas = Slice[int]([]int{1, 2, 3, 4, 5})
	if !Contains(3, datas.Iterator()) {
		t.Fatal("Contains error")
	}
	if Sum(datas.Iterator()) != 15 {
		t.Fatal("Sum error")
	}
	if Product(datas.Iterator()) != 120 {
		t.Fatal("Product error")
	}
	if Average(datas.Iterator()) != 3 {
		t.Fatal("Average error")
	}
	if Count(datas.Iterator()) != 5 {
		t.Fatal("Count error")
	}
	if Max(datas.Iterator()).OrPanic() != 5 {
		t.Fatal("Max error")
	}
	if Min(datas.Iterator()).OrPanic() != 1 {
		t.Fatal("Min error")
	}
	var sum int
	testSum := func(i int) {
		sum += i
	}
	ForEach(testSum, datas.Iterator())
	if sum != 15 {
		t.Fatal("ForEach error")
	}
	testMatch := func(i int) bool {
		return i > 0
	}
	if !AllMatch(testMatch, datas.Iterator()) {
		t.Fatal("AllMatch error")
	}
	if NoneMatch(testMatch, datas.Iterator()) {
		t.Fatal("NoneMatch error")
	}
	if !AnyMatch(testMatch, datas.Iterator()) {
		t.Fatal("AnyMatch error")
	}
	if First(datas.Iterator()).OrPanic() != 1 {
		t.Fatal("First error")
	}
	if Last(datas.Iterator()).OrPanic() != 5 {
		t.Fatal("Last error")
	}
	if At(2, datas.Iterator()).OrPanic() != 3 {
		t.Fatal("At error")
	}
	if Reduce(func(r int, t int) int {
		return r + t
	}, datas.Iterator()).OrPanic() != 15 {
		t.Fatal("Reduce error")
	}
	if Fold(0, func(r int, t int) int {
		return r + t
	}, datas.Iterator()) != 15 {
		t.Fatal("Fold error")
	}
}
