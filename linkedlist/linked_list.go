package list

import (
	"github.com/kulics/gollection/option"
	"github.com/kulics/gollection/ref"
	"github.com/kulics/gollection/seq"
)

func Of[T any](elements ...T) *List[T] {
	var list = &List[T]{0, nil, nil}
	for _, v := range elements {
		list.AddLast(v)
	}
	return list
}

func From[T any](collection seq.Collection[T]) *List[T] {
	var list = Of[T]()
	seq.ForEach[T](list.AddLast, collection)
	return list
}

type List[T any] struct {
	length int
	first  *LinkedListNode[T]
	last   *LinkedListNode[T]
}

// Returns the element at the start.
// Return None when the list is empty.
func (a *List[T]) First() ref.Ref[T] {
	if first := a.first; first != nil {
		return ref.Of(&first.Value)
	}
	return ref.Of[T](nil)
}

// Add element at the start.
func (a *List[T]) AddFirst(element T) {
	a.linkFirst(element)
}

// Remove element at the start.
// Return None when the list is empty.
func (a *List[T]) RemoveFirst() option.Option[T] {
	var first = a.first
	if first == nil {
		return option.None[T]()
	}
	return option.Some(a.unlinkFirst(first))
}

// Returns the element at the end.
// Return None when the list is empty.
func (a *List[T]) Last() ref.Ref[T] {
	if last := a.last; last != nil {
		return ref.Of(&last.Value)
	}
	return ref.Of[T](nil)
}

// Add element at the end.
func (a *List[T]) AddLast(element T) {
	a.linkLast(element)
}

// Remove element at the end.
// Return None when the list is empty.
func (a *List[T]) RemoveLast() option.Option[T] {
	var last = a.last
	if last == nil {
		return option.None[T]()
	}
	return option.Some(a.unlinkLast(last))
}

// Add element at the index.
func (a *List[T]) Add(index int, element T) {
	if index < 0 || index > a.length {
		panic(seq.OutOfBounds)
	}
	if index == 0 {
		a.linkLast(element)
	} else {
		a.linkBefore(element, a.at(index))
	}
}

// Add elements at the index.
func (a *List[T]) AddAll(index int, elements seq.Collection[T]) {
	if index < 0 || index > a.length {
		panic(seq.OutOfBounds)
	}
	var length = elements.Count()
	if length == 0 {
		return
	}
	var pred, succ *LinkedListNode[T]
	if index == a.length {
		succ = nil
		pred = a.last
	} else {
		succ = a.at(index)
		pred = succ.prev
	}
	for it := elements.Iterator(); true; {
		if v, ok := it.Next().Val(); ok {
			var newNode = &LinkedListNode[T]{Value: v, prev: pred, next: nil}
			if pred == nil {
				a.first = newNode
			} else {
				pred.next = newNode
			}
			pred = newNode
		} else {
			break
		}
	}
	if succ == nil {
		a.last = pred
	} else {
		pred.next = succ
		succ.prev = pred
	}
	a.length += length
}

// Remove element at the index.
func (a *List[T]) RemoveAt(index int) T {
	if a.isOutOfBounds(index) {
		panic(seq.OutOfBounds)
	}
	return a.unlink(a.at(index))
}

// Remove elements between begin and end.
func (a *List[T]) RemoveRange(begin, end int) {
	if a.isOutOfBounds(begin) || a.isOutOfBounds(end) {
		panic(seq.OutOfBounds)
	}
	if end == begin {
		return
	}
	var first = a.at(begin - 1)
	var last = first.next
	var index = begin
	for index < end {
		last = last.next
		index++
	}
	first.next = last
	last.prev = first
	a.length -= end - begin
}

// Clears all elements.
func (a *List[T]) Clear() {
	for x := a.first; x != nil; {
		var next = x.next
		var empty T
		x.Value = empty
		x.next = nil
		x.prev = nil
		x = next
	}
	a.first = nil
	a.last = nil
	a.length = 0
}

// Return the number of elements of list.
func (a *List[T]) Count() int {
	return a.length
}

// Return the Iterator of list.
func (a *List[T]) Iterator() seq.Iterator[T] {
	return &linkedListIterator[T]{a.first}
}

// Return a new list that copies all elements.
func (a *List[T]) Clone() *List[T] {
	return From[T](a)
}

func (a *List[T]) isOutOfBounds(index int) bool {
	if index < 0 || index >= a.length {
		return true
	}
	return false
}

func (a *List[T]) at(index int) *LinkedListNode[T] {
	if index < (a.length >> 1) {
		var x = a.first
		for i := 0; i < index; i++ {
			x = x.next
		}
		return x
	} else {
		var x = a.last
		for i := a.length - 1; i > index; i-- {
			x = x.prev
		}
		return x
	}
}

func (a *List[T]) linkFirst(element T) *LinkedListNode[T] {
	var first = a.first
	var newNode = &LinkedListNode[T]{Value: element, prev: nil, next: first}
	a.first = newNode
	if first == nil {
		a.last = newNode
	} else {
		first.prev = newNode
	}
	a.length++
	return newNode
}

func (a *List[T]) linkLast(element T) *LinkedListNode[T] {
	var last = a.last
	var newNode = &LinkedListNode[T]{Value: element, next: nil, prev: last}
	a.last = newNode
	if last == nil {
		a.first = newNode
	} else {
		last.next = newNode
	}
	a.length++
	return newNode
}

func (a *List[T]) linkBefore(element T, succ *LinkedListNode[T]) *LinkedListNode[T] {
	var pred = succ.prev
	var newNode = &LinkedListNode[T]{Value: element, prev: pred, next: succ}
	succ.prev = newNode
	if pred == nil {
		a.first = newNode
	} else {
		pred.next = newNode
	}
	a.length++
	return newNode
}

func (a *List[T]) linkAfter(element T, pred *LinkedListNode[T]) *LinkedListNode[T] {
	var succ = pred.next
	var newNode = &LinkedListNode[T]{Value: element, prev: pred, next: succ}
	pred.next = newNode
	if succ == nil {
		a.last = newNode
	} else {
		succ.prev = newNode
	}
	a.length++
	return newNode
}

func (a *List[T]) unlink(x *LinkedListNode[T]) T {
	var element = x.Value
	var next = x.next
	var prev = x.prev
	if prev == nil {
		a.first = next
	} else {
		prev.next = next
		x.prev = nil
	}

	if next == nil {
		a.last = prev
	} else {
		next.prev = prev
		x.next = nil
	}
	var empty T
	x.Value = empty
	a.length--
	return element
}

func (a *List[T]) unlinkFirst(x *LinkedListNode[T]) T {
	var element = x.Value
	var next = x.next
	var empty T
	x.Value = empty
	x.next = nil
	a.first = next
	if next == nil {
		a.last = nil
	} else {
		next.prev = nil
	}
	a.length--
	return element
}

func (a *List[T]) unlinkLast(x *LinkedListNode[T]) T {
	var element = x.Value
	var prev = x.prev
	var empty T
	x.Value = empty
	x.prev = nil
	a.last = prev
	if prev == nil {
		a.first = nil
	} else {
		prev.next = nil
	}
	a.length--
	return element
}

func (a *List[T]) Front() *LinkedListNode[T] {
	return a.first
}

func (a *List[T]) Back() *LinkedListNode[T] {
	return a.last
}

func (a *List[T]) Remove(mark *LinkedListNode[T]) T {
	return a.unlink(mark)
}

func (a *List[T]) InsertAfter(mark *LinkedListNode[T], newElement T) *LinkedListNode[T] {
	return a.linkAfter(newElement, mark)
}

func (a *List[T]) InsertBefore(mark *LinkedListNode[T], newElement T) *LinkedListNode[T] {
	return a.linkBefore(newElement, mark)
}

func (a *List[T]) InsertBack(newElement T) *LinkedListNode[T] {
	return a.linkFirst(newElement)
}

func (a *List[T]) InsertFront(newElement T) *LinkedListNode[T] {
	return a.linkLast(newElement)
}

type LinkedListNode[T any] struct {
	Value T
	next  *LinkedListNode[T]
	prev  *LinkedListNode[T]
}

func (a *LinkedListNode[T]) Next() *LinkedListNode[T] {
	return a.next
}

func (a *LinkedListNode[T]) Prev() *LinkedListNode[T] {
	return a.prev
}

type linkedListIterator[T any] struct {
	current *LinkedListNode[T]
}

func (a *linkedListIterator[T]) Next() option.Option[T] {
	if a.current != nil {
		var current = a.current.Value
		a.current = a.current.next
		return option.Some(current)
	}
	return option.None[T]()
}

func LinkedListCollector[T any]() seq.Collector[*List[T], T, *List[T]] {
	return linkedListCollector[T]{}
}

type linkedListCollector[T any] struct{}

func (a linkedListCollector[T]) Builder() *List[T] {
	return Of[T]()
}

func (a linkedListCollector[T]) Append(supplier *List[T], element T) {
	supplier.AddLast(element)
}

func (a linkedListCollector[T]) Finish(supplier *List[T]) *List[T] {
	return supplier
}
