package list

import (
	"testing"

	"github.com/kulics/gollection/iter"
)

func TestArrayList(t *testing.T) {
	var list = ArrayListOf[int]()
	if list.Count() != 0 {
		t.Fatal("list count not eq 0")
	}
	if list.Capacity() != defaultElementsLength {
		t.Fatal("list capacity not eq defaultElementsLength")
	}
	list.PushBack(1)
	if list.Count() != 1 {
		t.Fatal("list count not eq 1")
	}
	if list.Capacity() != defaultElementsLength {
		t.Fatal("list capacity not eq defaultElementsLength")
	}
	if list.At(0).Get() != 1 {
		t.Fatal("element of index 0 is not 1")
	}
	if list.At(2).IsNotNil() {
		t.Fatal("out of bounds check fail")
	}
	if list.At(-1).IsNotNil() {
		t.Fatal("out of bounds check fail")
	}
	if v := list.At(0).Set(2); v != 1 {
		t.Fatal("element of index 0 is not 1")
	}
	if list.At(0).Get() != 2 {
		t.Fatal("element of index 0 is not 2")
	}
	for i := 0; i < 10; i++ {
		list.PushBack(i)
	}
	if list.Count() != 11 {
		t.Fatal("list count not eq 11")
	}
	if list.Capacity() != 15 {
		t.Fatal("list capacity not grow *1.5")
	}
	list = list.Clone()
	if list.Count() != 11 {
		t.Fatal("list count not eq 11")
	}
	if list.Capacity() != 15 {
		t.Fatal("list capacity not grow *1.5")
	}
	list.Clear()
	if list.Count() != 0 {
		t.Fatal("list count not eq 0")
	}
	if list.Capacity() != 15 {
		t.Fatal("list capacity not grow *1.5")
	}
	list.Reserve(10)
	if list.Count() != 0 {
		t.Fatal("list count not eq 0")
	}
	if list.Capacity() != 15 {
		t.Fatal("list capacity not grow *1.5")
	}
	list.Reserve(30)
	if list.Count() != 0 {
		t.Fatal("list count not eq 0")
	}
	if list.Capacity() != 30 {
		t.Fatal("list capacity not grow to 30")
	}
	var listB = ArrayListFrom[int](ArrayListOf(1, 2, 3))
	if listB.Count() != 3 {
		t.Fatal("list count not eq 3")
	}
	if listB.Capacity() != 3 {
		t.Fatal("list capacity not eq 3")
	}
	list.InsertAll(0, listB)
	if list.Count() != 3 {
		t.Fatal("list count not eq 3")
	}
	if list.Capacity() != 30 {
		t.Fatal("list capacity not eq 30")
	}
	var it = list.Iterator()
	for i := 1; i <= 3; i++ {
		if item, ok := it.Next().Val(); ok && i != item {
			t.Fatal("element error")
		}
	}
	list.InsertAll(0, ArrayListOf(1, 2, 3))
	if list.Count() != 6 {
		t.Fatal("list count not eq 6")
	}
	list.RemoveRange(iter.RangeOf(1, 5))
	if list.Count() != 2 {
		t.Fatal("list count not eq 2")
	}
	if !Equals[int](ArrayListOf(1, 3), list) {
		t.Fatal("list elements not expect")
	}
	var _ IdxList[int] = list
}
