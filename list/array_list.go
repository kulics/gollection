package list

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

// Constructing an ArrayList with variable-length parameters
func ArrayListOf[T any](elements ...T) *ArrayList[T] {
	var length = len(elements)
	var list = MakeArrayList[T](length)
	copy(list.elements, elements)
	list.length = length
	return list
}

// Constructing an empty ArrayList with capacity.
func MakeArrayList[T any](capacity int) *ArrayList[T] {
	if capacity < defaultElementsLength {
		capacity = defaultElementsLength
	}
	return &ArrayList[T]{make([]T, capacity), 0}
}

// Constructing an ArrayList from other Collection.
func ArrayListFrom[T any](collection iter.Collection[T]) *ArrayList[T] {
	return &ArrayList[T]{iter.ToSlice(collection), collection.Count()}
}

// List implemented using Array.
// It has easier in-place modification than the built-in slice.
type ArrayList[T any] struct {
	elements []T
	length   int
}

func (a *ArrayList[T]) LastIndex() int {
	return a.length - 1
}

// Returns the element at the end.
// Return None when the list is empty.
func (a *ArrayList[T]) Peek() util.Ref[T] {
	return a.At(a.LastIndex())
}

// Add element at the end.
func (a *ArrayList[T]) Push(element T) {
	a.PushBack(element)
}

// Add multiple elements at the end.
func (a *ArrayList[T]) PushAll(elements iter.Collection[T]) {
	a.PushBackAll(elements)
}

// Remove element at the end.
// Return None when the list is empty.
func (a *ArrayList[T]) Pop() util.Opt[T] {
	return a.PopBack()
}

// Returns the element at the end.
// Return None when the list is empty.
func (a *ArrayList[T]) PeekBack() util.Ref[T] {
	return a.At(a.LastIndex())
}

// Add element at the end.
func (a *ArrayList[T]) PushBack(element T) {
	if growLength := a.length + 1; len(a.elements) < growLength {
		a.grow(growLength)
	}
	a.elements[a.length] = element
	a.length++
}

// Add multiple elements at the end.
func (a *ArrayList[T]) PushBackAll(elements iter.Collection[T]) {
	var additional = elements.Count()
	if growLength := a.length + additional; len(a.elements) < growLength {
		a.grow(growLength)
	}
	var i = a.length
	iter.ForEach(func(item T) {
		a.elements[i] = item
		a.length++
		i++
	}, elements.Iterator())
}

// Remove element at the end.
func (a *ArrayList[T]) PopBack() util.Opt[T] {
	if a.length == 0 {
		return util.None[T]()
	}
	var removed = a.elements[a.length-1]
	var emptyValue T
	a.elements[a.length-1] = emptyValue
	a.length--
	return util.Some(removed)
}

// Returns the element at the begin.
// Return None when the list is empty.
func (a *ArrayList[T]) PeekFront() util.Ref[T] {
	return a.At(0)
}

// Return the element at the index.
// Return None when a subscript is out of bounds.
func (a *ArrayList[T]) At(index int) util.Ref[T] {
	if a.isOutOfBounds(index) {
		return util.RefOf[T](nil)
	}
	return util.RefOf(&a.elements[index])
}

// Add element at the index.
func (a *ArrayList[T]) Insert(index int, element T) {
	if index < 0 || index > a.length {
		panic(iter.OutOfBounds)
	}
	if growLength := a.length + 1; len(a.elements) < growLength {
		a.grow(growLength)
	}
	copy(a.elements[index+1:], a.elements[index:])
	a.elements[index] = element
}

// Add multiple elements at the index.
func (a *ArrayList[T]) InsertAll(index int, elements iter.Collection[T]) {
	if index < 0 || index > a.length {
		panic(iter.OutOfBounds)
	}
	var additional = elements.Count()
	if growLength := a.length + additional; len(a.elements) < growLength {
		a.grow(growLength)
	}
	copy(a.elements[index+additional:], a.elements[index:])
	var i = index
	iter.ForEach(func(item T) {
		a.elements[i] = item
		a.length++
		i++
	}, elements.Iterator())
}

// Remove element at the index.
func (a *ArrayList[T]) Remove(index int) T {
	if a.isOutOfBounds(index) {
		panic(iter.OutOfBounds)
	}
	var removed = a.elements[index]
	copy(a.elements[:index], a.elements[index+1:])
	var emptyValue T
	a.elements[a.length-1] = emptyValue
	a.length--
	return removed
}

func (a *ArrayList[T]) RemoveRange(at iter.Range[int]) {
	var begin, end = at.Get()
	if a.isOutOfBounds(begin) || a.isOutOfBounds(end-1) {
		panic(iter.OutOfBounds)
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
func (a *ArrayList[T]) Reserve(additional int) {
	if addable := len(a.elements) - a.length; addable < additional {
		a.grow(a.length + additional)
	}
}

// Clears all elements, but does not reset the space.
func (a *ArrayList[T]) Clear() {
	var emptyValue T
	for i := 0; i < a.length; i++ {
		a.elements[i] = emptyValue
	}
	a.length = 0
}

// Return the number of elements of list.
func (a *ArrayList[T]) Count() int {
	return a.length
}

// Return the Iterator of list.
func (a *ArrayList[T]) Iterator() iter.Iterator[T] {
	return &arrayListIterator[T]{-1, a}
}

// Return a new list that copies all elements.
func (a *ArrayList[T]) Clone() *ArrayList[T] {
	var elements = make([]T, len(a.elements))
	copy(elements, a.elements)
	return &ArrayList[T]{
		elements: elements,
		length:   a.length,
	}
}

// Return the capacity of list.
func (a *ArrayList[T]) Capacity() int {
	return len(a.elements)
}

func (a *ArrayList[T]) isOutOfBounds(index int) bool {
	if index < 0 || index >= a.length {
		return true
	}
	return false
}

func (a *ArrayList[T]) grow(minCapacity int) {
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
	source *ArrayList[T]
}

func (a *arrayListIterator[T]) Next() util.Opt[T] {
	if a.index < a.source.Count()-1 {
		a.index++
		return util.Some(a.source.elements[a.index])
	}
	return util.None[T]()
}

func ArrayListCollector[T any]() iter.Collector[*ArrayList[T], T, *ArrayList[T]] {
	return arrayListCollector[T]{}
}

type arrayListCollector[T any] struct{}

func (a arrayListCollector[T]) Builder() *ArrayList[T] {
	return MakeArrayList[T](10)
}

func (a arrayListCollector[T]) Append(supplier *ArrayList[T], element T) {
	supplier.PushBack(element)
}

func (a arrayListCollector[T]) Finish(supplier *ArrayList[T]) *ArrayList[T] {
	return supplier
}
