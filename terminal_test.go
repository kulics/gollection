package gollection

import (
	"testing"
)

func TestTerminal(t *testing.T) {
	var datas = Slice[int]([]int{1, 2, 3, 4, 5})
	if !Contains(3, datas.Iter()) {
		t.Fatal("Contains error")
	}
	if Sum(datas.Iter()) != 15 {
		t.Fatal("Sum error")
	}
	if Product(datas.Iter()) != 120 {
		t.Fatal("Product error")
	}
	if Average(datas.Iter()) != 3 {
		t.Fatal("Average error")
	}
	if Count(datas.Iter()) != 5 {
		t.Fatal("Count error")
	}
	if Max(datas.Iter()) != 5 {
		t.Fatal("Max error")
	}
	if Min(datas.Iter()) != 1 {
		t.Fatal("Min error")
	}
	var sum int
	testSum := func(i int) {
		sum += i
	}
	ForEach(testSum, datas.Iter())
	if sum != 15 {
		t.Fatal("ForEach error")
	}
	testMatch := func(i int) bool {
		return i > 0
	}
	if !AllMatch(testMatch, datas.Iter()) {
		t.Fatal("AllMatch error")
	}
	if NoneMatch(testMatch, datas.Iter()) {
		t.Fatal("NoneMatch error")
	}
	if !AnyMatch(testMatch, datas.Iter()) {
		t.Fatal("AnyMatch error")
	}
	if First(datas.Iter()).OrPanic() != 1 {
		t.Fatal("First error")
	}
	if Last(datas.Iter()).OrPanic() != 5 {
		t.Fatal("Last error")
	}
	if At(2, datas.Iter()).OrPanic() != 3 {
		t.Fatal("At error")
	}
	if Reduce(0, func(r int, t int) int {
		return r + t
	}, datas.Iter()) != 15 {
		t.Fatal("Reduce error")
	}
	if Fold(0, func(r int, t int) int {
		return r + t
	}, datas.Iter()) != 15 {
		t.Fatal("Fold error")
	}
}
