package stack

import (
	. "github.com/kulics/gollection"
	. "github.com/kulics/gollection/union"
)

func LinkedStackOf[T any](elements ...T) LinkedStack[T] {
	var inner = &linkedStack[T]{0, nil}
	var stack = LinkedStack[T]{inner}
	for _, v := range elements {
		stack.Push(v)
	}
	return stack
}

func LinkedStackFrom[T any, I Collection[T]](collection I) LinkedStack[T] {
	var stack = LinkedStackOf[T]()
	ForEach(stack.Push, collection)
	return stack
}

type LinkedStack[T any] struct {
	inner *linkedStack[T]
}

type linkedStack[T any] struct {
	size  int
	first *node[T]
}

type node[T any] struct {
	value T
	next  *node[T]
}

func (a LinkedStack[T]) Size() int {
	return a.inner.size
}

func (a LinkedStack[T]) IsEmpty() bool {
	return a.inner.size == 0
}

func (a LinkedStack[T]) Push(element T) {
	if a.inner.first == nil {
		a.inner.first = &node[T]{element, nil}
	} else {
		a.inner.first = &node[T]{element, a.inner.first}
	}
	a.inner.size++
}

func (a LinkedStack[T]) Pop() T {
	if v, ok := a.TryPop().Get(); ok {
		return v
	}
	panic(OutOfBounds)
}

func (a LinkedStack[T]) Peek() T {
	if v, ok := a.TryPeek().Get(); ok {
		return v
	}
	panic(OutOfBounds)
}

func (a LinkedStack[T]) TryPop() Option[T] {
	if a.inner.first == nil {
		return None[T]()
	}
	a.inner.size--
	var item = a.inner.first.value
	a.inner.first = a.inner.first.next
	return Some(item)
}

func (a LinkedStack[T]) TryPeek() Option[T] {
	if a.inner.first == nil {
		return None[T]()
	}
	return Some(a.inner.first.value)
}

func (a LinkedStack[T]) Iter() Iterator[T] {
	return &linkedStackIterator[T]{a.inner.first}
}

func (a LinkedStack[T]) ToSlice() []T {
	var arr = make([]T, a.Size())
	ForEach(func(t T) {
		arr = append(arr, t)
	}, a)
	return arr
}

type linkedStackIterator[T any] struct {
	current *node[T]
}

func (a *linkedStackIterator[T]) Next() Option[T] {
	if a.current != nil {
		var item = a.current.value
		a.current = a.current.next
		return Some(item)
	}
	return None[T]()
}

func (a *linkedStackIterator[T]) Iter() Iterator[T] {
	return a
}
