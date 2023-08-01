package stack

import (
	"github.com/kulics/gollection/iter"
	"github.com/kulics/gollection/util"
)

func LinkedStackOf[T any](elements ...T) *LinkedStack[T] {
	var stack = &LinkedStack[T]{0, nil}
	for _, v := range elements {
		stack.Push(v)
	}
	return stack
}

func LinkedStackFrom[T any](collection iter.Collection[T]) *LinkedStack[T] {
	var stack = LinkedStackOf[T]()
	iter.ForEach(stack.Push, collection.Iterator())
	return stack
}

type LinkedStack[T any] struct {
	length int
	first  *linkedStackNode[T]
}

type linkedStackNode[T any] struct {
	value T
	next  *linkedStackNode[T]
}

func (a *LinkedStack[T]) Count() int {
	return a.length
}

func (a *LinkedStack[T]) Push(element T) {
	if a.first == nil {
		a.first = &linkedStackNode[T]{element, nil}
	} else {
		a.first = &linkedStackNode[T]{element, a.first}
	}
	a.length++
}

func (a *LinkedStack[T]) Pop() util.Opt[T] {
	if a.first == nil {
		return util.None[T]()
	}
	a.length--
	var item = a.first.value
	a.first = a.first.next
	return util.Some(item)
}

func (a *LinkedStack[T]) Peek() util.Ref[T] {
	if a.first == nil {
		return util.RefOf[T](nil)
	}
	return util.RefOf(&a.first.value)
}

func (a *LinkedStack[T]) Iterator() iter.Iterator[T] {
	return &linkedStackIterator[T]{a.first}
}

func (a *LinkedStack[T]) Clone() *LinkedStack[T] {
	return LinkedStackFrom[T](a)
}

func (a *LinkedStack[T]) Clear() {
	a.length = 0
	a.first = nil
}

type linkedStackIterator[T any] struct {
	current *linkedStackNode[T]
}

func (a *linkedStackIterator[T]) Next() util.Opt[T] {
	if a.current != nil {
		var value = a.current.value
		a.current = a.current.next
		return util.Some(value)
	}
	return util.None[T]()
}

func LinkedStackCollector[T any]() iter.Collector[*LinkedStack[T], T, *LinkedStack[T]] {
	return linkedStackCollector[T]{}
}

type linkedStackCollector[T any] struct{}

func (a linkedStackCollector[T]) Builder() *LinkedStack[T] {
	return LinkedStackOf[T]()
}

func (a linkedStackCollector[T]) Append(supplier *LinkedStack[T], element T) {
	supplier.Push(element)
}

func (a linkedStackCollector[T]) Finish(supplier *LinkedStack[T]) *LinkedStack[T] {
	return supplier
}
