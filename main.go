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
	list.foreach(func (it int)  {
		fmt.Println(it)
	})

	list = emptyArrayList[int]()
	for i := 0; i < 10; i ++ {
		list.append(i)
	}
	fmt.Println(list.data)

	fmt.Println(count[int](list))

	count := sum(mapto(func(it int) int { 
		return it * 2 
	})(filter(func(it int) bool { 
		if it % 2 == 0 { 
			return true
		} else { 
			return false 
		}
	})(list)))
	fmt.Println(count)
}

func emptyArrayList[T any]() *arraylist[T] {
	return &arraylist[T]{make([]T, 0)}
}

func arrayListOfCap[T any](capacity int) *arraylist[T] {
	return &arraylist[T]{make([]T, capacity)}
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

func (a *arraylist[T]) foreach(action func(T)) {
	for _, v := range a.data {
		action(v)
	}
}

type iterator[T any] interface {
	foreach(action func(T))
}

type mapStream[T any, R any] struct {
	transform func(T) R
	iter iterator[T]
}

func (m mapStream[T, R]) foreach(action func(R)) {
	m.iter.foreach(func(it T) {
		action(m.transform(it))
	})
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

func (f filterStream[T]) foreach(action func(T)) {
	f.iter.foreach(func(it T) {
		if f.predecate(it) {
			action(it)
		}
	})
}

func filter[T any](predecate func(T) bool) func(iterator[T]) iterator[T] {
	return func(it iterator[T]) iterator[T] {
		return filterStream[T]{predecate, it}
	}
}

type number interface {
	type int, int64, int32, int16, int8, uint64, uint32, uint16, uint8, float64, float32
}

func sum[T number](iter iterator[T]) T {
	var value T
	iter.foreach(func(it T) {
		value += it
	})
	return value
}

func max[T number](iter iterator[T]) T {
	var value T
	iter.foreach(func(it T) {
		if value < it {
			value = it
		}
	})
	return value
}

func min[T number](iter iterator[T]) T {
	var value T
	iter.foreach(func(it T) {
		if value > it {
			value = it
		}
	})
	return value
}

func count[T any](iter iterator[T]) int {
	var value int
	iter.foreach(func(it T) {
		value++
	})
	return value
}

func fold[T any, R any](initial R, operation func(R, T) R) func(iterator[T]) R {
	return func(iter iterator[T]) R {
		iter.foreach(func(it T) {
			initial = operation(initial, it)
		})
		return initial
	}
}