package list

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

// Constructing an List with variable-length parameters
func Of[T any](elements ...T) *List[T] {
	var length = len(elements)
	var list = Make[T](length)
	copy(list.elements, elements)
	list.length = length
	return list
}

// Constructing an empty List with capacity.
func Make[T any](capacity int) *List[T] {
	if capacity < defaultElementsLength {
		capacity = defaultElementsLength
	}
	return &List[T]{make([]T, capacity), 0}
}

// Constructing an List from other Collection.
func From[T any](collection seq.Collection[T]) *List[T] {
	return &List[T]{seq.ToSlice(collection), collection.Count()}
}

// List implemented using Array.
// It has easier in-place modification than the built-in slice.
type List[T any] struct {
	elements []T
	length   int
}

// Returns the index at the end.
func (a *List[T]) LastIndex() int {
	return a.length - 1
}

// Returns the element at the end.
// Return None when the list is empty.
func (a *List[T]) Last() ref.Ref[T] {
	return a.At(a.LastIndex())
}

// Add element at the end.
func (a *List[T]) AddLast(element T) {
	if growLength := a.length + 1; len(a.elements) < growLength {
		a.grow(growLength)
	}
	a.elements[a.length] = element
	a.length++
}

// Remove element at the end.
func (a *List[T]) RemoveLast() option.Option[T] {
	if a.length == 0 {
		return option.None[T]()
	}
	var removed = a.elements[a.length-1]
	var emptyValue T
	a.elements[a.length-1] = emptyValue
	a.length--
	return option.Some(removed)
}

// Returns the element at the begin.
// Return None when the list is empty.
func (a *List[T]) First() ref.Ref[T] {
	return a.At(0)
}

// Return the element at the index.
// Return None when a subscript is out of bounds.
func (a *List[T]) At(index int) ref.Ref[T] {
	if a.isOutOfBounds(index) {
		return ref.Of[T](nil)
	}
	return ref.Of(&a.elements[index])
}

// Add element at the index.
func (a *List[T]) Add(index int, element T) {
	if index < 0 || index > a.length {
		panic(seq.OutOfBounds)
	}
	if growLength := a.length + 1; len(a.elements) < growLength {
		a.grow(growLength)
	}
	copy(a.elements[index+1:], a.elements[index:])
	a.elements[index] = element
}

// Add multiple elements at the index.
func (a *List[T]) AddAll(index int, elements seq.Collection[T]) {
	if index < 0 || index > a.length {
		panic(seq.OutOfBounds)
	}
	var additional = elements.Count()
	if growLength := a.length + additional; len(a.elements) < growLength {
		a.grow(growLength)
	}
	copy(a.elements[index+additional:], a.elements[index:])
	var i = index
	seq.ForEach[T](func(item T) {
		a.elements[i] = item
		a.length++
		i++
	}, elements)
}

// Remove element at the index.
func (a *List[T]) Remove(index int) T {
	if a.isOutOfBounds(index) {
		panic(seq.OutOfBounds)
	}
	var removed = a.elements[index]
	copy(a.elements[:index], a.elements[index+1:])
	var emptyValue T
	a.elements[a.length-1] = emptyValue
	a.length--
	return removed
}

// Remove elements between begin and end.
func (a *List[T]) RemoveRange(begin, end int) {
	if a.isOutOfBounds(begin) || a.isOutOfBounds(end-1) {
		panic(seq.OutOfBounds)
	}
	if end == begin {
		return
	}
	copy(a.elements[begin:], a.elements[end:])
	var emptyValue T
	for i := len(a.elements) - (end - begin); i < len(a.elements); i++ {
		a.elements[i] = emptyValue
		a.length--
	}
}

// Ensure that list have enough space before expansion.
func (a *List[T]) Reserve(additional int) {
	if addable := len(a.elements) - a.length; addable < additional {
		a.grow(a.length + additional)
	}
}

// Clears all elements, but does not reset the space.
func (a *List[T]) Clear() {
	var emptyValue T
	for i := 0; i < a.length; i++ {
		a.elements[i] = emptyValue
	}
	a.length = 0
}

// Return the number of elements of list.
func (a *List[T]) Count() int {
	return a.length
}

// Return the Iterator of list.
func (a *List[T]) Iterator() seq.Iterator[T] {
	return &arrayListIterator[T]{-1, a}
}

// Return a new list that copies all elements.
func (a *List[T]) Clone() *List[T] {
	var elements = make([]T, len(a.elements))
	copy(elements, a.elements)
	return &List[T]{
		elements: elements,
		length:   a.length,
	}
}

// Return the capacity of list.
func (a *List[T]) Capacity() int {
	return len(a.elements)
}

func (a *List[T]) isOutOfBounds(index int) bool {
	if index < 0 || index >= a.length {
		return true
	}
	return false
}

func (a *List[T]) grow(minCapacity int) {
	var newLength = arrayGrow(a.length)
	if newLength < minCapacity {
		newLength = minCapacity
	}
	var newSource = make([]T, newLength)
	copy(newSource, a.elements)
	a.elements = newSource
}

type arrayListIterator[T any] struct {
	index  int
	source *List[T]
}

func (a *arrayListIterator[T]) Next() option.Option[T] {
	if a.index < a.source.Count()-1 {
		a.index++
		return option.Some(a.source.elements[a.index])
	}
	return option.None[T]()
}

func Collector[T any]() seq.Collector[*List[T], T, *List[T]] {
	return collector[T]{}
}

type collector[T any] struct{}

func (a collector[T]) Builder() *List[T] {
	return Make[T](10)
}

func (a collector[T]) Append(supplier *List[T], element T) {
	supplier.AddLast(element)
}

func (a collector[T]) Finish(supplier *List[T]) *List[T] {
	return supplier
}
