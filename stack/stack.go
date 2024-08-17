package stack

import (
	"github.com/kulics/gollection/option"
	"github.com/kulics/gollection/ref"
	"github.com/kulics/gollection/seq"
)

const defaultElementsLength = 10

func arrayGrow(length int) int {
	var newLength = length + (length >> 1)
	if newLength < defaultElementsLength {
		newLength = defaultElementsLength
	}
	return newLength
}

// Constructing an Stack with variable-length parameters
func Of[T any](elements ...T) *Stack[T] {
	var length = len(elements)
	var stack = Make[T](length)
	copy(stack.elements, elements)
	stack.length = length
	return stack
}

// Constructing an empty Stack with capacity.
func Make[T any](capacity int) *Stack[T] {
	if capacity < defaultElementsLength {
		capacity = defaultElementsLength
	}
	return &Stack[T]{make([]T, capacity), 0}
}

// Constructing an Stack from other Collection.
func From[T any](collection seq.Collection[T]) *Stack[T] {
	return &Stack[T]{seq.ToSlice(collection), collection.Count()}
}

// Stack implemented using Array.
type Stack[T any] struct {
	elements []T
	length   int
}

// Return the number of elements of stack.
func (a *Stack[T]) Count() int {
	return a.length
}

// Add an element to the top of the stack.
func (a *Stack[T]) AddLast(element T) {
	if growLength := a.length + 1; len(a.elements) < growLength {
		a.grow(growLength)
	}
	a.elements[a.length] = element
	a.length++
}

// Remove an element from the top of the stack.
// Return None when the stack is empty.
func (a *Stack[T]) RemoveLast() option.Option[T] {
	if seq.IsEmpty[T](a) {
		return option.None[T]()
	}
	var index = a.length - 1
	var item = a.elements[index]
	var empty T
	a.elements[index] = empty
	a.length--
	return option.Some(item)
}

// Return an element at the top of the stack, but does not remove it.
// Return None when the stack is empty.
func (a *Stack[T]) Last() ref.Ref[T] {
	if seq.IsEmpty[T](a) {
		return ref.Of[T](nil)
	}
	return ref.Of(&a.elements[a.length-1])
}

// Return the Iterator of stack.
func (a *Stack[T]) Iterator() seq.Iterator[T] {
	return &iterator[T]{a.Count(), a}
}

// Return a new stack that copies all elements.
func (a *Stack[T]) Clone() *Stack[T] {
	var elements = make([]T, len(a.elements))
	copy(elements, a.elements)
	return &Stack[T]{
		elements: elements,
		length:   a.length,
	}
}

// Ensure that stack have enough space before expansion.
func (a *Stack[T]) Reserve(additional int) {
	if addable := len(a.elements) - a.length; addable < additional {
		a.grow(a.length + additional)
	}
}

// Return the capacity of stack.
func (a *Stack[T]) Capacity() int {
	return len(a.elements)
}

// Clears all elements, but does not reset the space.
func (a *Stack[T]) Clear() {
	var emptyValue T
	for i := 0; i < a.length; i++ {
		a.elements[i] = emptyValue
	}
	a.length = 0
}

func (a *Stack[T]) grow(minCapacity int) {
	var newLength = arrayGrow(len(a.elements))
	if newLength < minCapacity {
		newLength = minCapacity
	}
	var newSource = make([]T, newLength)
	copy(newSource, a.elements)
	a.elements = newSource
}

type iterator[T any] struct {
	index  int
	source *Stack[T]
}

func (a *iterator[T]) Next() option.Option[T] {
	if a.index > 0 {
		a.index--
		return option.Some(a.source.elements[a.index])
	}
	return option.None[T]()
}

func Collector[T any]() seq.Collector[*Stack[T], T, *Stack[T]] {
	return collector[T]{}
}

type collector[T any] struct{}

func (a collector[T]) Builder() *Stack[T] {
	return Make[T](10)
}

func (a collector[T]) Append(supplier *Stack[T], element T) {
	supplier.AddLast(element)
}

func (a collector[T]) Finish(supplier *Stack[T]) *Stack[T] {
	return supplier
}
