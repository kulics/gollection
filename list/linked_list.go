package list

import (
	"github.com/kulics/gollection/iter"
	"github.com/kulics/gollection/util"
)

func LinkedListOf[T any](elements ...T) *LinkedList[T] {
	var list = &LinkedList[T]{0, nil, nil}
	for _, v := range elements {
		list.PushBack(v)
	}
	return list
}

func LinkedListFrom[T any](collection iter.Collection[T]) *LinkedList[T] {
	var list = LinkedListOf[T]()
	iter.ForEach(list.PushBack, collection.Iterator())
	return list
}

type LinkedList[T any] struct {
	length int
	first  *LinkedListNode[T]
	last   *LinkedListNode[T]
}

// Returns the element at the end.
// Return None when the list is empty.
func (a *LinkedList[T]) Peek() util.Ref[T] {
	return a.PeekBack()
}

// Add element at the end.
func (a *LinkedList[T]) Push(element T) {
	a.PushBack(element)
}

// Remove element at the end.
// Return None when the list is empty.
func (a *LinkedList[T]) Pop() util.Opt[T] {
	return a.PopBack()
}

func (a *LinkedList[T]) PeekFront() util.Ref[T] {
	if first := a.first; first != nil {
		return util.RefOf(&first.Value)
	}
	return util.RefOf[T](nil)
}

func (a *LinkedList[T]) PushFront(element T) {
	a.linkFirst(element)
}

func (a *LinkedList[T]) PopFront() util.Opt[T] {
	var first = a.first
	if first == nil {
		util.None[T]()
	}
	return util.Some(a.unlinkFirst(first))
}

func (a *LinkedList[T]) PeekBack() util.Ref[T] {
	if last := a.last; last != nil {
		return util.RefOf(&last.Value)
	}
	return util.RefOf[T](nil)
}

func (a *LinkedList[T]) PushBack(element T) {
	a.linkLast(element)
}

func (a *LinkedList[T]) PopBack() util.Opt[T] {
	var last = a.last
	if last == nil {
		util.None[T]()
	}
	return util.Some(a.unlinkLast(last))
}

func (a *LinkedList[T]) insert(index int, element T) {
	if index < 0 || index > a.length {
		panic(iter.OutOfBounds)
	}
	if index == 0 {
		a.linkLast(element)
	} else {
		a.linkBefore(element, a.at(index))
	}
}

func (a *LinkedList[T]) insertAll(index int, elements iter.Collection[T]) {
	if index < 0 || index > a.length {
		panic(iter.OutOfBounds)
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

func (a *LinkedList[T]) removeAt(index int) T {
	if a.isOutOfBounds(index) {
		panic(iter.OutOfBounds)
	}
	return a.unlink(a.at(index))
}

func (a *LinkedList[T]) removeRange(at iter.Range[int]) {
	var begin, end = at.Get()
	if a.isOutOfBounds(begin) || a.isOutOfBounds(end) {
		panic(iter.OutOfBounds)
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

func (a *LinkedList[T]) Clear() {
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

func (a *LinkedList[T]) Count() int {
	return a.length
}

func (a *LinkedList[T]) Iterator() iter.Iterator[T] {
	return &linkedListIterator[T]{a.first}
}

func (a *LinkedList[T]) Clone() *LinkedList[T] {
	return LinkedListFrom[T](a)
}

func (a *LinkedList[T]) isOutOfBounds(index int) bool {
	if index < 0 || index >= a.length {
		return true
	}
	return false
}

func (a *LinkedList[T]) at(index int) *LinkedListNode[T] {
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

func (a *LinkedList[T]) linkFirst(element T) *LinkedListNode[T] {
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

func (a *LinkedList[T]) linkLast(element T) *LinkedListNode[T] {
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

func (a *LinkedList[T]) linkBefore(element T, succ *LinkedListNode[T]) *LinkedListNode[T] {
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

func (a *LinkedList[T]) linkAfter(element T, pred *LinkedListNode[T]) *LinkedListNode[T] {
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

func (a *LinkedList[T]) unlink(x *LinkedListNode[T]) T {
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

func (a *LinkedList[T]) unlinkFirst(x *LinkedListNode[T]) T {
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

func (a *LinkedList[T]) unlinkLast(x *LinkedListNode[T]) T {
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

func (a *LinkedList[T]) Front() *LinkedListNode[T] {
	return a.first
}

func (a *LinkedList[T]) Back() *LinkedListNode[T] {
	return a.last
}

func (a *LinkedList[T]) Remove(mark *LinkedListNode[T]) T {
	return a.unlink(mark)
}

func (a *LinkedList[T]) InsertAfter(mark *LinkedListNode[T], newElement T) *LinkedListNode[T] {
	return a.linkAfter(newElement, mark)
}

func (a *LinkedList[T]) InsertBefore(mark *LinkedListNode[T], newElement T) *LinkedListNode[T] {
	return a.linkBefore(newElement, mark)
}

func (a *LinkedList[T]) InsertBack(newElement T) *LinkedListNode[T] {
	return a.linkFirst(newElement)
}

func (a *LinkedList[T]) InsertFront(newElement T) *LinkedListNode[T] {
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

func (a *linkedListIterator[T]) Next() util.Opt[T] {
	if a.current != nil {
		var current = a.current.Value
		a.current = a.current.next
		return util.Some(current)
	}
	return util.None[T]()
}

func LinkedListCollector[T any]() iter.Collector[*LinkedList[T], T, *LinkedList[T]] {
	return linkedListCollector[T]{}
}

type linkedListCollector[T any] struct{}

func (a linkedListCollector[T]) Builder() *LinkedList[T] {
	return LinkedListOf[T]()
}

func (a linkedListCollector[T]) Append(supplier *LinkedList[T], element T) {
	supplier.PushBack(element)
}

func (a linkedListCollector[T]) Finish(supplier *LinkedList[T]) *LinkedList[T] {
	return supplier
}
