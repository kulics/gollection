package gollection

import "testing"

func TestArrayList(t *testing.T) {
	var list = ArrayListOf[int]()
	if list.Size() != 0 {
		t.Fatal("array list size not eq 0")
	}
	if list.Capacity() != defaultElementsSize {
		t.Fatal("array list capacity not eq defaultElementsSize")
	}
	list.Append(1)
	if list.Size() != 1 {
		t.Fatal("array list size not eq 1")
	}
	if list.Capacity() != defaultElementsSize {
		t.Fatal("array list capacity not eq defaultElementsSize")
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
	if list.Size() != 11 {
		t.Fatal("array list size not eq 11")
	}
	if list.Capacity() != 15 {
		t.Fatal("array list capacity not grow *1.5")
	}
	list = list.Clone()
	if list.Size() != 11 {
		t.Fatal("array list size not eq 11")
	}
	if list.Capacity() != 15 {
		t.Fatal("array list capacity not grow *1.5")
	}
	list.Clear()
	if list.Size() != 0 {
		t.Fatal("array list size not eq 0")
	}
	if list.Capacity() != 15 {
		t.Fatal("array list capacity not grow *1.5")
	}
	list.Reserve(10)
	if list.Size() != 0 {
		t.Fatal("array list size not eq 0")
	}
	if list.Capacity() != 15 {
		t.Fatal("array list capacity not grow *1.5")
	}
	list.Reserve(30)
	if list.Size() != 0 {
		t.Fatal("array list size not eq 0")
	}
	if list.Capacity() != 30 {
		t.Fatal("array list capacity not grow to 30")
	}
	var slice = list.ToSlice()
	if len(slice) != 0 {
		t.Fatal("ToSlice size not eq to 0")
	}
	var listB = ArrayListFrom[int](ArrayListOf(1, 2, 3))
	if listB.Size() != 3 {
		t.Fatal("array list size not eq 3")
	}
	if listB.Capacity() != 3 {
		t.Fatal("array list capacity not eq 3")
	}
	list.PrependAll(listB)
	if list.Size() != 3 {
		t.Fatal("array list size not eq 3")
	}
	if list.Capacity() != 30 {
		t.Fatal("array list capacity not eq 30")
	}
	var iter = list.Iter()
	for i := 1; i <= 3; i++ {
		var item = iter.Next()
		if i != item.value {
			t.Fatal("element error")
		}
	}
}
