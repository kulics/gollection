package gollection

import (
	"testing"
)

func TestArrayStack(t *testing.T) {
	var stack = ArrayStackOf[int]()
	if stack.Count() != 0 {
		t.Fatal("stack count not eq 0")
	}
	if stack.Capacity() != defaultElementsLength {
		t.Fatal("stack capacity not eq defaultElementsLength")
	}
	stack.Push(1)
	if stack.Count() != 1 {
		t.Fatal("stack count not eq 1")
	}
	if stack.Capacity() != defaultElementsLength {
		t.Fatal("stack capacity not eq defaultElementsLength")
	}
	if stack.Peek() != 1 {
		t.Fatal("element of index 0 is not 1")
	}
	if stack.Count() != 1 {
		t.Fatal("stack count not eq 1")
	}
	if v := stack.Pop(); v != 1 {
		t.Fatal("element of top is not 1")
	}
	if stack.TryPop().IsSome() {
		t.Fatal("stack must has not element")
	}
	for i := 0; i <= 10; i++ {
		stack.Push(i)
	}
	if stack.Count() != 11 {
		t.Fatal("stack count not eq 11")
	}
	if stack.Capacity() != 15 {
		t.Fatal("stack capacity not grow *1.5")
	}
	stack = stack.Clone()
	if stack.Count() != 11 {
		t.Fatal("stack count not eq 11")
	}
	if stack.Capacity() != 15 {
		t.Fatal("stack capacity not grow *1.5")
	}
	stack.Clear()
	if stack.Count() != 0 {
		t.Fatal("stack count not eq 0")
	}
	if stack.Capacity() != 15 {
		t.Fatal("stack capacity not grow *1.5")
	}
	stack.Reserve(10)
	if stack.Count() != 0 {
		t.Fatal("stack count not eq 0")
	}
	if stack.Capacity() != 15 {
		t.Fatal("stack capacity not grow *1.5")
	}
	stack.Reserve(30)
	if stack.Count() != 0 {
		t.Fatal("stack count not eq 0")
	}
	if stack.Capacity() != 30 {
		t.Fatal("stack capacity not grow to 30")
	}
	var slice = stack.ToSlice()
	if len(slice) != 0 {
		t.Fatal("ToSlice count not eq to 0")
	}
	var stackB = ArrayStackFrom[int](ArrayListOf(1, 2, 3))
	if stackB.Count() != 3 {
		t.Fatal("stack count not eq 3")
	}
	if stackB.Capacity() != 3 {
		t.Fatal("stack capacity not eq 3")
	}
	var iter = stackB.Iter()
	for i := 3; i >= 1; i-- {
		var item = iter.Next()
		if i != item.OrPanic() {
			t.Fatal("element error")
		}
	}
	var _ AnyStack[int] = stack
}
