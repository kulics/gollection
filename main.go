package main

import "fmt"

func main() {
	fmt.Println((1))
	var list = emptyArrayList[int]()
	for i := 0; i < 10; i ++ {
		list.append(i)
	}
	fmt.Println(list.data)
	list = arrayListOf(7,7,7,7,7,7,7)
	fmt.Println(list.data)
	list.remove(0)
	fmt.Println(list.data)
	fmt.Println(list.size())
	fmt.Println(list.isEmpty())
	fmt.Println(list.get(1))
	list.set(1, 5)
	fmt.Println(list.data)

	list = emptyArrayList[int]()
	for i := 0; i < 10; i ++ {
		list.append(i)
	}
	fmt.Println(list.data)

	fmt.Println(count[int](list.iter()))

	count := sum(mapto(func(it int) int { 
		return it * 2 
	})(filter(func(it int) bool { 
		if it % 2 == 0 { 
			return true
		} else { 
			return false 
		}
	})(list.iter())))
	fmt.Println(count)
}

func emptyArrayStack[T any]() *arrayStack[T] {
	return &arrayStack[T]{make([]T, 0)}
}

type arrayStack[T any] struct {
	source []T
}

func (a *arrayStack[T]) size() int  {
	return len(a.source)
}

func (a *arrayStack[T]) isEmpty() bool  {
	return a.size() == 0
}

func (a *arrayStack[T]) push(element T) {
	a.source = append(a.source, element)
}

func (a *arrayStack[T]) pop() option[T] {
	var size = len(a.source)
	if size == 0 {
		return none[T]()
	}
	var item = a.source[size-1]
	a.source = a.source[:size-1]
	return some(item)
}

func (a *arrayStack[T]) peek() option[T] {
	var size = len(a.source)
	if size == 0 {
		return none[T]()
	}
	return some(a.source[size-1])
}

type node[T any] struct {
	value T
	next *node[T]
}

func emptyLinkedStack[T any]() *linkedStack[T] {
	return &linkedStack[T]{0, nil}
}

type linkedStack[T any] struct {
	currentSize int
	head *node[T]
}

func (a *linkedStack[T]) size() int  {
	return a.currentSize
}

func (a *linkedStack[T]) isEmpty() bool  {
	return a.size() == 0
}

func (a *linkedStack[T]) push(element T) {
	a.currentSize++
	if a.head == nil {
		a.head = &node[T]{element, nil}
	} else {
		a.head = &node[T]{element, a.head}
	}
}

func (a *linkedStack[T]) pop() option[T] {
	if a.head == nil {
		return none[T]()
	}
	a.currentSize--
	var item = a.head.value
	a.head = a.head.next
	return some(item)
}

func (a *linkedStack[T]) peek() option[T] {
	if a.head == nil {
		return none[T]()
	}
	return some(a.head.value)
}

func emptyArrayList[T any]() *arrayList[T] {
	return &arrayList[T]{make([]T, 0)}
}

func newArrayListWithCap[T any](capacity int) *arrayList[T] {
	return &arrayList[T]{make([]T, 0, capacity)}
}

func arrayListOf[T any](elements ...T) *arrayList[T] {
	return &arrayList[T]{elements}
}

type arrayList[T any] struct {
	data []T
}

func (a *arrayList[T]) prepend(element T) {
	a.insert(0, element)
}

func (a *arrayList[T]) append(element T) {
	a.data = append(a.data, element)
}

func (a *arrayList[T]) insert(index int, element T) bool {
	if index < 0 || index >= a.size() {
		return false
	}
	a.data = append(a.data, element)
	copy(a.data[index+1:], a.data[index:])
	a.data[index] = element
	return true
}

func (a *arrayList[T]) remove(index int) option[T] {
	if index < 0 || index >= a.size() {
		return none[T]()
	}
	var removed = a.data[index]
	a.data = append(a.data[:index], a.data[index+1:]...) 
	return some(removed)
}

func (a *arrayList[T]) get(index int) option[T] {
	if index < 0 || index >= a.size() {
		return none[T]()
	}
	return some(a.data[index])
}

func (a *arrayList[T]) set(index int, newElement T) option[T] {
	if index < 0 || index >= a.size() {
		return none[T]()
	}
	var oldElement = a.data[index]
	a.data[index] = newElement
	return some(oldElement)
}

func (a *arrayList[T]) clean() {
	a.data = make([]T, 0)
}

func (a *arrayList[T]) size() int {
	return len(a.data)
}

func (a *arrayList[T]) isEmpty() bool {
	return a.size()== 0
}

func (a *arrayList[T]) iter() iterator[T] {
	return &arrayListIterator[T]{-1, *a}
}

type arrayListIterator[T any] struct {
	index int
	source arrayList[T]
}

func (a *arrayListIterator[T]) next() option[T] {
	if a.index < a.source.size()-1 {
		a.index++
		return some(a.source.data[a.index])
	}
	return none[T]()
}

func emptyLinkedList[T any]() *linkedList[T] {
	return &linkedList[T]{0, nil}
}

type linkedList[T any] struct {
	currentSize int
	head *node[T]
}

func (a *linkedList[T]) prepend(element T) {
	if a.head == nil {
		a.head = &node[T]{element, nil}
	} else {
		a.head = &node[T]{element, a.head}
	}
	a.currentSize++
}

func (a *linkedList[T]) append(element T) {
	addNode(a.head, element)
	a.currentSize++
}

func addNode[T any](n *node[T], v T) {
	if n == nil {
		*n = node[T]{v, nil}
	} else {
		addNode(n.next, v)
	}
}

func (a *linkedList[T]) insert(index int, element T) bool {
	if index < 0 || index >= a.size() {
		return false
	}
	if index == 0 {
		a.prepend(element)
	} else if index == a.size() {
		a.append(element)
	} else {
		insertNode(a.head.next, a.head, index-1, element)
	}
	return true
}

func insertNode[T any](n *node[T], pre *node[T], i int, e T) {
	if i == 0 {
		pre.next = &node[T]{e, n}
	} else {
		insertNode(n.next, n, i-1, e)
	}
}

func (a *linkedList[T]) remove(index int) option[T] {
	if index < 0 || index >= a.size() {
		return none[T]()
	}
	var item T
	if index == 0 {
		var temp = a.head
		a.head = a.head.next
		temp.next = nil
		item = temp.value
	} else {
		item = removeNode(a.head.next, a.head, index-1)
	}
	a.currentSize--
	return some(item)
}

func removeNode[T any](n *node[T], pre *node[T], i int) T {
	if i == 0 {
		var item = n.value
		pre.next = n.next
		n.next = nil
		return item
	} else {
		return removeNode(n.next, n, i-1)
	}
}

func (a *linkedList[T]) get(index int) option[T] {
	if index < 0 || index >= a.size() {
		return none[T]()
	}
	return some(getNode(a.head, index))
}

func getNode[T any](n *node[T], i int) T {
	if i == 0 {
		return n.value
	} else {
		return getNode(n.next, i-1)
	}
}

func (a *linkedList[T]) set(index int, newElement T) option[T] {
	if index < 0 || index >= a.size() {
		return none[T]()
	}
	return some(setNode(a.head, index, newElement))
}

func setNode[T any](n *node[T], i int, v T) T {
	if i == 0 {
		var oldValue = n.value
		n.value = v
		return oldValue
	} else {
		return setNode(n.next, i-1, v)
	}
}

func (a *linkedList[T]) clean() {
	a.head = nil
	a.currentSize = 0
}

func (a *linkedList[T]) size() int {
	return a.currentSize
}

func (a *linkedList[T]) isEmpty() bool {
	return a.size()== 0
}

func (a *linkedList[T]) iter() iterator[T] {
	return &linkedListIterator[T]{a.head}
}

type linkedListIterator[T any] struct {
	current *node[T]
}

func (a *linkedListIterator[T]) next() option[T] {
	if a.current != nil {
		var item = a.current.value
		a.current = a.current.next
		return some(item)
	}
	return none[T]()
}

type iterator[T any] interface {
	next() option[T]
}

type iterable[T any] interface {
	iter() iterator[T]
}

type option[T any] struct {
	val T
	ok bool
}

func (o option[T]) orElse(v T) T {
	if o.ok {
		return o.val
	}
	return v
}

func some[T any](value T) option[T] {
	return option[T]{value, true}
}

func none[T any]() option[T] {
	var empty T
	return option[T]{empty, false}
}

type sliceIterator[T any] struct {
	index int
	source []T
}

func (s *sliceIterator[T]) next() option[T] {
	if s.index < len(s.source) - 1 {
		s.index++
		return some(s.source[s.index])
	}
	return none[T]()
}

func sliceIter[T any](source []T) iterator[T] {
	return &sliceIterator[T]{0, source}
}

type indexStream[T any] struct {
	index int
	iter iterator[T]
}

type pair[T any, R any] struct {
	first T
	second R
}

func (i *indexStream[T]) next() option[pair[int, T]] {
	if v := i.iter.next(); v.ok {
		i.index++
		return some(pair[int, T]{i.index, v.val})
	}
	return none[pair[int, T]]()
}

func withIndex[T any](iter iterator[T]) iterator[pair[int, T]] {
	return &indexStream[T]{-1, iter}
}

type mapStream[T any, R any] struct {
	transform func(T) R
	iter iterator[T]
}

func (m mapStream[T, R]) next() option[R] {
	if v := m.iter.next(); v.ok {
		return some(m.transform(v.val))
	}
	return none[R]()
}

func mapto[T any, R any](transform func(T) R) func(iterator[T]) iterator[R] {
	return func(it iterator[T]) iterator[R] {
		return mapStream[T, R]{transform, it}
	}
}

type filterStream[T any] struct {
	predecate func(T) bool
	iter iterator[T]
}

func (f filterStream[T]) next() option[T] {
	for v := f.iter.next(); v.ok; v = f.iter.next() {
		if f.predecate(v.val) {
			return v
		}
	}
	return none[T]()
}

func filter[T any](predecate func(T) bool) func(iterator[T]) iterator[T] {
	return func(it iterator[T]) iterator[T] {
		return filterStream[T]{predecate, it}
	}
}

type limitStream[T any] struct {
	limit int
	index int
	iter iterator[T]
}

func (l *limitStream[T]) next() option[T] {
	if v := l.iter.next(); v.ok && l.index < l.limit {
		l.index++
		return v
	}
	return none[T]()
}

func limit[T any](count int) func(iterator[T]) iterator[T] {
	return func(it iterator[T]) iterator[T] {
		return &limitStream[T]{count, 0, it}
	}
}

type skipStream[T any] struct {
	skip int
	index int
	iter iterator[T]
}

func (l *skipStream[T]) next() option[T] {
	for v := l.iter.next(); v.ok; v = l.iter.next() {
		if l.index < l.skip {
			l.index++
			continue
		}
		return v
	}
	return none[T]()
}

func skip[T any](count int) func(iterator[T]) iterator[T] {
	return func(it iterator[T]) iterator[T] {
		return &skipStream[T]{count, 0, it}
	}
}

type stepStream[T any] struct {
	step int
	index int
	iter iterator[T]
}

func (l *stepStream[T]) next() option[T] {
	for v := l.iter.next(); v.ok; v = l.iter.next() {
		if l.index < l.step {
			l.index++
			continue
		}
		l.index = 0
		return v
	}
	return none[T]()
}

func step[T any](count int) func(iterator[T]) iterator[T] {
	return func(it iterator[T]) iterator[T] {
		return &stepStream[T]{count, 0, it}
	}
}

type concatStream[T any] struct {
	first option[iterator[T]]
	last iterator[T]
}

func (l *concatStream[T]) next() option[T] {
	if l.first.ok {
		if v := l.first.val.next(); v.ok {
			return v
		}
		l.first = none[iterator[T]]()
		return l.next()
	}
	return l.last.next()
}

func concat[T any](left iterator[T]) func (iterator[T]) iterator[T] {
	return func(right iterator[T]) iterator[T] {
		return &concatStream[T]{some(left), right}
	}
}

type number interface {
	type int, int64, int32, int16, int8, uint64, uint32, uint16, uint8, float64, float32
}

func sum[T number](iter iterator[T]) T {
	var result T
	for v := iter.next(); v.ok; v = iter.next() {
		result += v.val
	}
	return result
}

func count[T any](iter iterator[T]) int {
	var result int
	for v := iter.next(); v.ok; v = iter.next() {
		result++
	}
	return result
}

func max[T number](iter iterator[T]) T {
	var result T
	for v := iter.next(); v.ok; v = iter.next() {
		if result < v.val {
			result = v.val
		}
	}
	return result
}

func min[T number](iter iterator[T]) T {
	var result T
	for v := iter.next(); v.ok; v = iter.next() {
		if result > v.val {
			result = v.val
		}
	}
	return result
}

func foreach[T any](action func(T)) func(iterator[T]) {
	return func(iter iterator[T]) {
		for v := iter.next(); v.ok; v = iter.next() {
			action(v.val)
		}
	}
}

func allMatch[T any](predicate func(T) bool) func(iterator[T]) bool {
	return func(iter iterator[T]) bool {
		for v := iter.next(); v.ok; v = iter.next() {
			if !predicate(v.val) {
				return false
			}
		}
		return true
	}
}

func noneMatch[T any](predicate func(T) bool) func(iterator[T]) bool {
	return func(iter iterator[T]) bool {
		for v := iter.next(); v.ok; v = iter.next() {
			if predicate(v.val) {
				return false
			}
		}
		return true
	}
}

func anyMatch[T any](predicate func(T) bool) func(iterator[T]) bool {
	return func(iter iterator[T]) bool {
		for v := iter.next(); v.ok; v = iter.next() {
			if predicate(v.val) {
				return true
			}
		}
		return false
	}
}

func first[T any](iter iterator[T]) option[T] {
	return iter.next()
}

func last[T any](iter iterator[T]) option[T] {
	var result = iter.next()
	for v := iter.next(); v.ok; v = iter.next() {
		result = v
	}
	return result
}

func fold[T any, R any](initial R, operation func(R, T) R) func(iterator[T]) R {
	return func(iter iterator[T]) R {
		var result = initial
		for v := iter.next(); v.ok; v = iter.next() {
			result = operation(result, v.val)
		}
		return result
	}
}
