package gollection

import (
	"testing"
)

func TestArrayList(t *testing.T) {
	var list = ArrayListOf[int]()
	if list.Count() != 0 {
		t.Fatal("list length not eq 0")
	}
	if list.Capacity() != defaultElementsLength {
		t.Fatal("list capacity not eq defaultElementsLength")
	}
	list.Append(1)
	if list.Count() != 1 {
		t.Fatal("list length not eq 1")
	}
	if list.Capacity() != defaultElementsLength {
		t.Fatal("list capacity not eq defaultElementsLength")
	}
	if list.Get(0) != 1 {
		t.Fatal("element of index 0 is not 1")
	}
	if list.TryGet(2).IsSome() {
		t.Fatal("out of bounds check fail")
	}
	if list.TryGet(-1).IsSome() {
		t.Fatal("out of bounds check fail")
	}
	if v := list.Set(0, 2); v != 1 {
		t.Fatal("element of index 0 is not 1")
	}
	if list.Get(0) != 2 {
		t.Fatal("element of index 0 is not 2")
	}
	for i := 0; i < 10; i++ {
		list.Append(i)
	}
	if list.Count() != 11 {
		t.Fatal("list length not eq 11")
	}
	if list.Capacity() != 15 {
		t.Fatal("list capacity not grow *1.5")
	}
	list = list.Clone()
	if list.Count() != 11 {
		t.Fatal("list length not eq 11")
	}
	if list.Capacity() != 15 {
		t.Fatal("list capacity not grow *1.5")
	}
	list.Clear()
	if list.Count() != 0 {
		t.Fatal("list length not eq 0")
	}
	if list.Capacity() != 15 {
		t.Fatal("list capacity not grow *1.5")
	}
	list.Reserve(10)
	if list.Count() != 0 {
		t.Fatal("list length not eq 0")
	}
	if list.Capacity() != 15 {
		t.Fatal("list capacity not grow *1.5")
	}
	list.Reserve(30)
	if list.Count() != 0 {
		t.Fatal("list length not eq 0")
	}
	if list.Capacity() != 30 {
		t.Fatal("list capacity not grow to 30")
	}
	var slice = list.ToSlice()
	if len(slice) != 0 {
		t.Fatal("ToSlice length not eq to 0")
	}
	var listB = ArrayListFrom[int](ArrayListOf(1, 2, 3))
	if listB.Count() != 3 {
		t.Fatal("list length not eq 3")
	}
	if listB.Capacity() != 3 {
		t.Fatal("list capacity not eq 3")
	}
	list.PrependAll(listB)
	if list.Count() != 3 {
		t.Fatal("list length not eq 3")
	}
	if list.Capacity() != 30 {
		t.Fatal("list capacity not eq 30")
	}
	var it = list.Iter()
	for i := 1; i <= 3; i++ {
		var item = it.Next()
		if i != item.OrPanic() {
			t.Fatal("element error")
		}
	}
	list.PrependAll(ArrayListOf(1, 2, 3))
	if list.Count() != 6 {
		t.Fatal("list length not eq 6")
	}
	list.RemoveRange(RangeOf(1, 5))
	if list.Count() != 2 {
		t.Fatal("list length not eq 2")
	}
	if !EqualsList[int](ArrayListOf(1, 3), list) {
		t.Fatal("list elements not expect")
	}
	var _ AnyList[int] = list
	var _ AnyMutableList[int] = list
}
