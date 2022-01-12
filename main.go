package main


func main() {
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

func (a *arrayStack[T]) pop() (value T, ok bool) {
	var size = len(a.source)
	if size == 0 {
		return
	}
	var item = a.source[size-1]
	a.source = a.source[:size-1]
	return item, true
}

func (a *arrayStack[T]) peek() (value T, ok bool) {
	var size = len(a.source)
	if size == 0 {
		return
	}
	return a.source[size - 1], true
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

func (a *linkedStack[T]) pop() (value T, ok bool) {
	if a.head == nil {
		return
	}
	a.currentSize--
	var item = a.head.value
	a.head = a.head.next
	return item, true
}

func (a *linkedStack[T]) peek() (value T, ok bool) {
	if a.head == nil {
		return
	}
	return a.head.value, true
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

func (a *arrayList[T]) remove(index int) (value T, ok bool) {
	if index < 0 || index >= a.size() {
		return
	}
	var removed = a.data[index]
	a.data = append(a.data[:index], a.data[index+1:]...) 
	return removed, true
}

func (a *arrayList[T]) get(index int) (value T, ok bool) {
	if index < 0 || index >= a.size() {
		return
	}
	return a.data[index], true
}

func (a *arrayList[T]) set(index int, newElement T) (value T, ok bool) {
	if index < 0 || index >= a.size() {
		return
	}
	var oldElement = a.data[index]
	a.data[index] = newElement
	return oldElement, true
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

func (a *arrayListIterator[T]) next() (value T, ok bool) {
	if a.index < a.source.size()-1 {
		a.index++
		return a.source.data[a.index], true
	}
	return
}

func (a *arrayListIterator[T]) iter() iterator[T] {
	return a
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

func (a *linkedList[T]) remove(index int) (value T, ok bool) {
	if index < 0 || index >= a.size() {
		return
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
	return item, true
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

func (a *linkedList[T]) get(index int) (value T, ok bool) {
	if index < 0 || index >= a.size() {
		return
	}
	return getNode(a.head, index), true
}

func getNode[T any](n *node[T], i int) T {
	if i == 0 {
		return n.value
	} else {
		return getNode(n.next, i-1)
	}
}

func (a *linkedList[T]) set(index int, newElement T) (value T, ok bool) {
	if index < 0 || index >= a.size() {
		return
	}
	return setNode(a.head, index, newElement), true
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

func (a *linkedListIterator[T]) next() (value T, ok bool) {
	if a.current != nil {
		var item = a.current.value
		a.current = a.current.next
		return item, true
	}
	return
}

func (a *linkedListIterator[T]) iter() iterator[T] {
	return a
}

type iterator[T any] interface {
	iter() iterator[T]
	next() (T, bool)
}

type iterable[T any] interface {
	iter() iterator[T]
}

type sliceIterator[T any] struct {
	index int
	source []T
}

func (s *sliceIterator[T]) next() (value T, ok bool) {
	if s.index < len(s.source) - 1 {
		s.index++
		return s.source[s.index], true
	}
	return
}

func (s *sliceIterator[T]) iter() iterator[T] {
	return s
}

func sliceIter[T any](source []T) iterator[T] {
	return &sliceIterator[T]{-1, source}
}

type indexStream[T any] struct {
	index int
	iterator iterator[T]
}

type pair[T any, R any] struct {
	first T
	second R
}

func (i *indexStream[T]) next() (value pair[int, T], ok bool) {
	if v, ok := i.iterator.next(); ok {
		i.index++
		return pair[int, T]{i.index, v}, true
	}
	return
}

func (i *indexStream[T]) iter() iterator[pair[int, T]] {
	return i
}

func withIndex[T any](it iterable[T]) iterator[pair[int, T]] {
	return &indexStream[T]{-1, it.iter()}
}

type mapStream[T any, R any] struct {
	transform func(T) R
	iterator iterator[T]
}

func (m mapStream[T, R]) next() (value R, ok bool) {
	if v, ok := m.iterator.next(); ok {
		return m.transform(v), true
	}
	return
}

func (m mapStream[T, R]) iter() iterator[R] {
	return m
}

func mapto[T any, R any](transform func(T) R, it iterable[T]) iterator[R] {
	return mapStream[T, R]{transform, it.iter()}
}

type filterStream[T any] struct {
	predecate func(T) bool
	iterator iterator[T]
}

func (f filterStream[T]) next() (value T, ok bool) {
	for v, ok := f.iterator.next(); ok; v, ok = f.iterator.next() {
		if f.predecate(v) {
			return v, true
		}
	}
	return
}

func (f filterStream[T]) iter() iterator[T] {
	return f
}

func filter[T any](predecate func(T) bool, it iterable[T]) iterator[T] {
	return filterStream[T]{predecate, it.iter()}
}

type limitStream[T any] struct {
	limit int
	index int
	iterator iterator[T]
}

func (l *limitStream[T]) next() (value T, ok bool) {
	if v, ok := l.iterator.next(); ok && l.index < l.limit {
		l.index++
		return v, true
	}
	return
}

func (l *limitStream[T]) iter() iterator[T] {
	return l
}

func limit[T any](count int, it iterable[T]) iterator[T] {
	return &limitStream[T]{count, 0, it.iter()}
}

type skipStream[T any] struct {
	skip int
	index int
	iterator iterator[T]
}

func (l *skipStream[T]) next() (value T, ok bool) {
	for v, ok := l.iterator.next(); ok; v, ok = l.iterator.next() {
		if l.index < l.skip {
			l.index++
			continue
		}
		return v, true
	}
	return
}

func (l *skipStream[T]) iter() iterator[T] {
	return l
}

func skip[T any](count int, it iterable[T]) iterator[T] {
	return &skipStream[T]{count, 0, it.iter()}
}

type stepStream[T any] struct {
	step int
	index int
	iterator iterator[T]
}

func (l *stepStream[T]) next() (value T, ok bool) {
	for v, ok := l.iterator.next(); ok; v, ok = l.iterator.next() {
		if l.index < l.step {
			l.index++
			continue
		}
		l.index = 1
		return v, true
	}
	return
}

func (l *stepStream[T]) iter() iterator[T] {
	return l
}

func step[T any](count int, it iterable[T]) iterator[T] {
	return &stepStream[T]{count, count, it.iter()}
}

type concatStream[T any] struct {
	firstok bool
	first iterator[T]
	last iterator[T]
}

func (l *concatStream[T]) next() (value T, ok bool) {
	if l.firstok {
		if v, ok := l.first.next(); ok {
			return v, true
		}
		l.firstok = false
		return l.next()
	}
	return l.last.next()
}

func (l *concatStream[T]) iter() iterator[T] {
	return l
}

func concat[T any](left iterable[T], right iterable[T]) iterator[T] {
	return &concatStream[T]{false, left.iter(), right.iter()}
}

type number interface {
	type int, int64, int32, int16, int8, uint64, uint32, uint16, uint8, float64, float32
}

func sum[T number](it iterable[T]) T {
	var result T
	var iter = it.iter()
	for v, ok := iter.next(); ok; v, ok = iter.next() {
		result += v
	}
	return result
}

func count[T any](it iterable[T]) int {
	var result int
	var iter = it.iter()
	for _, ok := iter.next(); ok; _, ok = iter.next() {
		result++
	}
	return result
}

func max[T number](it iterable[T]) T {
	var result T
	var iter = it.iter()
	for v, ok := iter.next(); ok; v, ok = iter.next() {
		if result < v {
			result = v
		}
	}
	return result
}

func min[T number](it iterable[T]) T {
	var result T
	var iter = it.iter()
	for v, ok := iter.next(); ok; v, ok = iter.next() {
		if result > v {
			result = v
		}
	}
	return result
}

func foreach[T any](action func(T), it iterable[T]) {
	var iter = it.iter()
	for v, ok := iter.next(); ok; v, ok = iter.next() {
		action(v)
	}
}

func allMatch[T any](predicate func(T) bool, it iterable[T]) bool {
	var iter = it.iter()
	for v, ok := iter.next(); ok; v, ok = iter.next() {
		if !predicate(v) {
			return false
		}
	}
	return true
}

func noneMatch[T any](predicate func(T) bool, it iterable[T]) bool {
	var iter = it.iter()
	for v, ok := iter.next(); ok; v, ok = iter.next() {
		if predicate(v) {
			return false
		}
	}
	return true
}

func anyMatch[T any](predicate func(T) bool, it iterable[T]) bool {
	var iter = it.iter()
	for v, ok := iter.next(); ok; v, ok = iter.next() {
		if predicate(v) {
			return true
		}
	}
	return false
}

func first[T any](it iterable[T]) (value T, ok bool) {
	return it.iter().next()
}

func last[T any](it iterable[T]) (value T, ok bool) {
	var iter = it.iter()
	v, ok := iter.next()
	for ok {
		v, ok = iter.next()
	}
	return v, ok
}

func at[T any](index int, it iterable[T]) (value T, ok bool) {
	var iter = it.iter()
	v, ok := iter.next()
	var i = 0
	for i < index && ok {
		v, ok = iter.next()
		i++
	}
	return v, ok
}

func reduce[T any, R any](initial R, operation func(R, T) R, it iterable[T]) R {
	var iter = it.iter()
	var result = initial
	for v, ok := iter.next(); ok; v, ok = iter.next() {
		result = operation(result, v)
	}
	return result
}

func fold[T any, R any](initial R, operation func(T, R) R, it iterable[T]) R {
	var reverse = make([]T, 0)
	var iter = it.iter()
	for v, ok := iter.next(); ok; v, ok = iter.next() {
		reverse = append(reverse, v)
	}
	var result = initial
	for i := len(reverse)-1; i > 0; i-- {
		result = operation(reverse[i], result)
	}
	return result
}