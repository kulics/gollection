package stack

import (
	"github.com/kulics/gollection/iter"
	"github.com/kulics/gollection/util"
)

const defaultElementsLength = 10

func arrayGrow(length int) int {
	var newLength = length + (length >> 1)
	if newLength < defaultElementsLength {
		newLength = defaultElementsLength
	}
	return newLength
}

// Constructing an ArrayStack with variable-length parameters
func ArrayStackOf[T any](elements ...T) *ArrayStack[T] {
	var length = len(elements)
	var stack = MakeArrayStack[T](length)
	copy(stack.elements, elements)
	stack.length = length
	return stack
}

// Constructing an empty ArrayStack with capacity.
func MakeArrayStack[T any](capacity int) *ArrayStack[T] {
	if capacity < defaultElementsLength {
		capacity = defaultElementsLength
	}
	return &ArrayStack[T]{make([]T, capacity), 0}
}

// Constructing an ArrayStack from other Collection.
func ArrayStackFrom[T any](collection iter.Collection[T]) *ArrayStack[T] {
	return &ArrayStack[T]{iter.ToSlice(collection), collection.Count()}
}

// Stack implemented using Array.
type ArrayStack[T any] struct {
	elements []T
	length   int
}

// Return the number of elements of stack.
func (a *ArrayStack[T]) Count() int {
	return a.length
}

// Add an element to the top of the stack.
func (a *ArrayStack[T]) Push(element T) {
	if growLength := a.length + 1; len(a.elements) < growLength {
		a.grow(growLength)
	}
	a.elements[a.length] = element
	a.length++
}

// Add multiple elements to the top of the stack.
func (a *ArrayStack[T]) PushAll(elements iter.Collection[T]) {
	var additional = elements.Count()
	if growLength := a.length + additional; len(a.elements) < growLength {
		a.grow(growLength)
	}
	var i = a.length
	for it := elements.Iterator(); true; {
		if item, ok := it.Next().Val(); ok {
			a.elements[i] = item
			a.length++
			i++
		} else {
			break
		}
	}
}

// Remove an element from the top of the stack.
// Return None when the stack is empty.
func (a *ArrayStack[T]) Pop() util.Opt[T] {
	if iter.IsEmpty[T](a) {
		return util.None[T]()
	}
	var index = a.length - 1
	var item = a.elements[index]
	var empty T
	a.elements[index] = empty
	a.length--
	return util.Some(item)
}

// Return an element at the top of the stack, but does not remove it.
// Return None when the stack is empty.
func (a *ArrayStack[T]) Peek() util.Ref[T] {
	if iter.IsEmpty[T](a) {
		return util.RefOf[T](nil)
	}
	return util.RefOf(&a.elements[a.length-1])
}

// Return the Iterator of stack.
func (a *ArrayStack[T]) Iterator() iter.Iterator[T] {
	return &arrayStackIterator[T]{a.Count(), a}
}

// Return a new stack that copies all elements.
func (a *ArrayStack[T]) Clone() *ArrayStack[T] {
	var elements = make([]T, len(a.elements))
	copy(elements, a.elements)
	return &ArrayStack[T]{
		elements: elements,
		length:   a.length,
	}
}

// Ensure that stack have enough space before expansion.
func (a *ArrayStack[T]) Reserve(additional int) {
	if addable := len(a.elements) - a.length; addable < additional {
		a.grow(a.length + additional)
	}
}

// Return the capacity of stack.
func (a *ArrayStack[T]) Capacity() int {
	return len(a.elements)
}

// Clears all elements, but does not reset the space.
func (a *ArrayStack[T]) Clear() {
	var emptyValue T
	for i := 0; i < a.length; i++ {
		a.elements[i] = emptyValue
	}
	a.length = 0
}

func (a *ArrayStack[T]) grow(minCapacity int) {
	var newLength = arrayGrow(len(a.elements))
	if newLength < minCapacity {
		newLength = minCapacity
	}
	var newSource = make([]T, newLength)
	copy(newSource, a.elements)
	a.elements = newSource
}

type arrayStackIterator[T any] struct {
	index  int
	source *ArrayStack[T]
}

func (a *arrayStackIterator[T]) Next() util.Opt[T] {
	if a.index > 0 {
		a.index--
		return util.Some(a.source.elements[a.index])
	}
	return util.None[T]()
}

func ArrayStackCollector[T any]() iter.Collector[*ArrayStack[T], T, *ArrayStack[T]] {
	return arrayStackCollector[T]{}
}

type arrayStackCollector[T any] struct{}

func (a arrayStackCollector[T]) Builder() *ArrayStack[T] {
	return MakeArrayStack[T](10)
}

func (a arrayStackCollector[T]) Append(supplier *ArrayStack[T], element T) {
	supplier.Push(element)
}

func (a arrayStackCollector[T]) Finish(supplier *ArrayStack[T]) *ArrayStack[T] {
	return supplier
}
