package gollection

import "testing"

func TestLinkedList(t *testing.T) {
	var list = LinkedListOf[int]()
	if list.Size() != 0 {
		t.Fatal("list size not eq 0")
	}
	list.Append(1)
	if list.Size() != 1 {
		t.Fatal("list size not eq 1")
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
		t.Fatal("list size not eq 11")
	}
	list = list.Clone()
	if list.Size() != 11 {
		t.Fatal("list size not eq 11")
	}
	list.Clear()
	if list.Size() != 0 {
		t.Fatal("list size not eq 0")
	}
	var slice = list.ToSlice()
	if len(slice) != 0 {
		t.Fatal("ToSlice size not eq to 0")
	}
	var listB = LinkedListFrom[int](ArrayListOf(1, 2, 3))
	if listB.Size() != 3 {
		t.Fatal("list size not eq 3")
	}
	list.PrependAll(listB)
	if list.Size() != 3 {
		t.Fatal("list size not eq 3")
	}
	var iter = list.Iter()
	for i := 1; i <= 3; i++ {
		var item = iter.Next()
		if i != item.value {
			t.Fatal("element error")
		}
	}
	list.PrependAll(ArrayListOf(1, 2, 3))
	if list.Size() != 6 {
		t.Fatal("list size not eq 6")
	}
	if list.GetFirst() != 1 {
		t.Fatal("first elements of list is not 1")
	}
	if list.RemoveFirst() != 1 {
		t.Fatal("remove list first is not 1")
	}
	if list.GetFirst() != 2 {
		t.Fatal("first elements of list is not 2")
	}
	if list.GetLast() != 3 {
		t.Fatal("last elements of list is not 3")
	}
	if list.RemoveLast() != 3 {
		t.Fatal("remove list last is not 3")
	}
	if list.GetLast() != 2 {
		t.Fatal("last elements of list is not 2")
	}
}
