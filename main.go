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

func emptyArrayList[T any]() *arraylist[T] {
	return &arraylist[T]{make([]T, 0)}
}

func arrayListOfCap[T any](capacity int) *arraylist[T] {
	return &arraylist[T]{make([]T, 0, capacity)}
}

func arrayListOf[T any](elements ...T) *arraylist[T] {
	return &arraylist[T]{elements}
}

type arraylist[T any] struct {
	data []T
}

func (a *arraylist[T]) append(element T)  {
	a.data = append(a.data, element)
}

func (a *arraylist[T]) remove(index int) T {
	var removed = a.data[index]
	a.data = append(a.data[:index], a.data[index+1:]...) 
	return removed
}

func (a *arraylist[T]) get(index int) T {
	return a.data[index]
}

func (a *arraylist[T]) set(index int, newElement T) {
	a.data[index] = newElement
}

func (a *arraylist[T]) clean() {
	a.data = make([]T, 0)
}

func (a *arraylist[T]) size() int {
	return len(a.data)
}

func (a *arraylist[T]) isEmpty() bool {
	return len(a.data)== 0
}

func (a *arraylist[T]) iter() iterator[T] {
	return &arraylistIterator[T]{-1, *a}
}

type arraylistIterator[T any] struct {
	index int
	source arraylist[T]
}

func (a *arraylistIterator[T]) next() option[T] {
	if a.index < a.source.size()-1 {
		a.index++
		return some(a.source.data[a.index])
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
