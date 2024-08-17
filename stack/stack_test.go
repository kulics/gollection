package stack

import (
	"testing"
)

func TestArrayStack(t *testing.T) {
	var stack = Of[int]()
	if stack.Count() != 0 {
		t.Fatal("stack count not eq 0")
	}
	if stack.Capacity() != defaultElementsLength {
		t.Fatal("stack capacity not eq defaultElementsLength")
	}
	stack.AddLast(1)
	if stack.Count() != 1 {
		t.Fatal("stack count not eq 1")
	}
	if stack.Capacity() != defaultElementsLength {
		t.Fatal("stack capacity not eq defaultElementsLength")
	}
	if stack.Last().Get() != 1 {
		t.Fatal("element of index 0 is not 1")
	}
	if stack.Count() != 1 {
		t.Fatal("stack count not eq 1")
	}
	if v := stack.RemoveLast().OrPanic(); v != 1 {
		t.Fatal("element of top is not 1")
	}
	if stack.RemoveLast().IsSome() {
		t.Fatal("stack must has not element")
	}
	for i := 0; i <= 10; i++ {
		stack.AddLast(i)
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
	var stackB = From[int](Of(3, 2, 1))
	if stackB.Count() != 3 {
		t.Fatal("stack count not eq 3")
	}
	if stackB.Capacity() != 3 {
		t.Fatalf("stack capacity not eq 3, it is eq %d", stackB.Capacity())
	}
	var iter = stackB.Iterator()
	for i := 3; i >= 1; i-- {
		if item, ok := iter.Next().Val(); ok && i != item {
			t.Fatal("element error")
		}
	}
}
