# gollection

A generic generic collection library based on go's generic implementation.

## Core Interfaces

```go
type Iterable[T any] interface {
	Iter() Iterator[T]
}

type Iterator[T any] interface {
	Next() util.Opt[T]
}
```

gollection applies the iterator pattern design, the core interface consists of Iterator and Iterable.

Iterator is responsible for providing iterative functionality, Iterator is unidirectional and lazy, each call to next will only return one result.

Iterable is responsible for providing Iterator, the implementation type determines whether the provided Iterator is reusable.

The inert traversal feature allows the combination of higher-order functions without significant overhead and can provide a richer combination of functions.

Here is a simple example of direct traversal:

```go
func printAll[T any](it Iterable[T]) {
	for iter = it.Iter(); true; {
		if v, ok := iter.Next().Get(); ok {
			println(v)
		} else {
			break
		}
	}
}
```

## Streaming operations

gollection provides a rich set of stream manipulation functions that can be used in combination with any Iterable.

Here is an example of a simple combination used:

```go
func foo() {
	show := func(i int) {
		println(i)
	}
	even := func(i int) bool {
		return i%2 == 0
	}
	square := func(i int) int {
		return i * i
	}
	ForEach(show, Map(square, Filter(even, Slice[int]([]int{1, 2, 3, 4, 5, 6, 7}).Iterator())))
    // Result:
    // 4
    // 16
    // 36
}
```

### Transform Iterable

A series of conversion functions are provided to process one Iterable conversion to another Iterable. these conversions are not executed immediately and only act one at a time when iterating.

```go
func Enumerate[T any](it Iterator[T]) Iterator[Pair[int, T]]
func Map[T any, R any](transform func(T) R, it Iterator[T]) Iterator[R]
func Filter[T any](predicate func(T) bool, it Iterator[T]) Iterator[T]
func Limit[T any](count int, it Iterator[T]) Iterator[T]
func Skip[T any](count int, it Iterator[T]) Iterator[T]
func Step[T any](count int, it Iterator[T]) Iterator[T]
func Concat[T any](left Iterator[T], right Iterator[T]) Iterator[T]
func Flatten[T Iterable[U], U any](it Iterator[T]) Iterator[U]
func Zip[T any, U any](left Iterator[T], right Iterator[U]) Iterator[util.Pair[T, U]]
```

### Terminal Iterable

A set of functions that evaluate the Iterable and are executed immediately.

```go
func Contains[T comparable](target T, it Iterator[T]) bool
func Sum[T Integer | Float](it Iterator[T]) T
func Product[T Integer | Float](it Iterator[T]) T
func Average[T Integer | Float](it Iterator[T]) float64
func Count[T any](it Iterator[T]) int
func Max[T Ordered](it Iterator[T]) util.Opt[T]
func Min[T Ordered](it Iterator[T]) util.Opt[T]
func MaxBy[T any](greater func(T, T) bool, it Iterator[T]) util.Opt[T]
func MinBy[T any](less func(T, T) bool, it Iterator[T]) util.Opt[T]
func ForEach[T any](action func(T), it Iterator[T])
func AllMatch[T any](predicate func(T) bool, it Iterator[T]) bool
func NoneMatch[T any](predicate func(T) bool, it Iterator[T]) bool
func AnyMatch[T any](predicate func(T) bool, it Iterator[T]) bool
func First[T any](it Iterator[T]) util.Opt[T]
func Last[T any](it Iterator[T]) util.Opt[T]
func At[T any](index int, it Iterator[T]) util.Opt[T]
func Reduce[T any](operation func(T, T) T, it Iterator[T]) util.Opt[T]
func Fold[T any, R any](initial R, operation func(T, R) R, it Iterator[T]) R
func Unzip[A any, B any](it Iterator[util.Pair[A, B]]) util.Pair[[]A, []B]
func Collect[T any, S any, R any](collector Collector[S, T, R], it Iterator[T]) R
```

## ToString and ToSlice

In order to make go's native string and slice also iterable, we have introduced `ToSlice` and `ToString` to make these two types implement the interface.

```go
var str = "Hello, world!"
var sli = []int{1, 2, 3}
Count(ToString(str).Iter()) // 13
Count(ToSlice(sli).Iter()) // 3
```

We also provide the version that gets the iterator directly.

```go
Count(ToStringIter(str)) // 13
Count(ToSliceIter(sli)) // 3
```

## Collection

We define a unified collection type interface to describe more information than iterators to facilitate performance optimization.

```go
type Collection[T any] interface {
	Iterable[T]

	Count() int
}

func IsEmpty[T any](c Collection[T]) bool
func IsNotEmpty[T any](c Collection[T]) bool
func ToSlice[T any](c Collection[T]) Slice[T]
```

### List

We provide the `List` interface, `ForwardList` interface, `BackwardList` interface and `IndexList` interface to describe the ordered sequences, and provide `ArrayList` and `LinkedList` as its implementation types.

```go
package list

type List[T any] interface {
	iter.Collection[T]

	Peek() util.Ref[T]
	Push(element T)
	Pop() util.Opt[T]
	Clear()
}

type ForwardList[T any] interface {
	List[T]

	PeekFront() util.Ref[T]
	PushFront(element T)
	PopFront() util.Opt[T]
}

type BackwardList[T any] interface {
	List[T]

	PeekBack() util.Ref[T]
	PopBack() util.Opt[T]
	PushBack(element T)
}

type IndexList[T any] interface {
	List[T]

	At(index int) util.Ref[T]
	Insert(index int, element T)
	Remove(index int) T
	RemoveRange(at iter.Range[int])
}

func ArrayListOf[T any](elements ...T) *ArrayList[T]
func MakeArrayList[T any](capacity int) *ArrayList[T]
func ArrayListFrom[T any](collection iter.Collection[T]) *ArrayList[T]

type ArrayList[T any] struct {}
func (a *ArrayList[T]) LastIndex() int
func (a *ArrayList[T]) Peek() util.Ref[T]
func (a *ArrayList[T]) Push(element T)
func (a *ArrayList[T]) PushAll(elements iter.Collection[T])
func (a *ArrayList[T]) Pop() util.Opt[T]
func (a *ArrayList[T]) PeekBack() util.Ref[T]
func (a *ArrayList[T]) PushBack(element T)
func (a *ArrayList[T]) PushBackAll(elements iter.Collection[T])
func (a *ArrayList[T]) PopBack() util.Opt[T]
func (a *ArrayList[T]) PeekFront() util.Ref[T]
func (a *ArrayList[T]) At(index int) util.Ref[T]
func (a *ArrayList[T]) Insert(index int, element T)
func (a *ArrayList[T]) InsertAll(index int, elements iter.Collection[T])
func (a *ArrayList[T]) Remove(index int) T
func (a *ArrayList[T]) RemoveRange(at iter.Range[int])
func (a *ArrayList[T]) Reserve(additional int)
func (a *ArrayList[T]) Count() int
func (a *ArrayList[T]) Capacity() int
func (a *ArrayList[T]) Iterator() iter.Iterator[T]
func (a *ArrayList[T]) Clone() *ArrayList[T]
func (a *ArrayList[T]) Clear()

func LinkedListOf[T any](elements ...T) *LinkedList[T]
func LinkedListFrom[T any](collection iter.Collection[T]) *LinkedList[T]

type LinkedList[T any] struct {}
func (a *LinkedList[T]) Peek() util.Ref[T]
func (a *LinkedList[T]) Push(element T)
func (a *LinkedList[T]) Pop() util.Opt[T]
func (a *LinkedList[T]) PeekFront() util.Ref[T]
func (a *LinkedList[T]) PushFront(element T)
func (a *LinkedList[T]) PopFront() util.Opt[T]
func (a *LinkedList[T]) PeekBack() util.Ref[T]
func (a *LinkedList[T]) PushBack(element T)
func (a *LinkedList[T]) PopBack() util.Opt[T]
func (a *LinkedList[T]) Count() int
func (a *LinkedList[T]) Iterator() iter.Iterator[T]
func (a *LinkedList[T]) Clone() *LinkedList[T]
func (a *LinkedList[T]) Clear()
func (a *LinkedList[T]) Front() *LinkedListNode[T]
func (a *LinkedList[T]) Back() *LinkedListNode[T]
func (a *LinkedList[T]) Remove(mark *LinkedListNode[T]) T
func (a *LinkedList[T]) InsertFront(newElement T) *LinkedListNode[T]
func (a *LinkedList[T]) InsertBack(newElement T) *LinkedListNode[T]
func (a *LinkedList[T]) InsertAfter(mark *LinkedListNode[T], newElement T) *LinkedListNode[T]
func (a *LinkedList[T]) InsertBefore(mark *LinkedListNode[T], newElement T) *LinkedListNode[T]

type LinkedListNode[T any] struct {
	Value T
}
func (a *LinkedListNode[T]) Next() *LinkedListNode[T]
func (a *LinkedListNode[T]) Prev() *LinkedListNode[T]
```

### Dict

We provide the `Dict` interface to describe the mapping type, and provide `HashMap` as its implementation type.

```go
package dict

type Dict[K any, V any] interface {
	iter.Collection[util.Pair[K, V]]

	Contains(key K) bool
	At(key K) util.Ref[V]
	Put(key K, value V) util.Opt[V]
	Remove(key K) util.Opt[V]
	Clear()
}

func Equals[K any, V comparable](l Dict[K, V], r Dict[K, V]) bool

func HashDictOf[K comparable, V any](elements ...util.Pair[K, V]) *HashDict[K, V]
func MakeHashDict[K comparable, V any](capacity int) *HashDict[K, V]
func MakeHashDictWithHasher[K comparable, V any](hasher func(K) uint64, capacity int) *HashDict[K, V]
func HashDictFrom[K comparable, V any](collection iter.Collection[util.Pair[K, V]]) *HashDict[K, V]

type HashDict[K comparable, V any] struct
func (a *HashDict[K, V]) Count() int
func (a *HashDict[K, V]) Contains(key K) bool
func (a *HashDict[K, V]) At(key K) util.Ref[V]
func (a *HashDict[K, V]) Put(key K, value V) util.Opt[V]
func (a *HashDict[K, V]) Remove(key K) util.Opt[V]
func (a *HashDict[K, V]) Iterator() iter.Iterator[util.Pair[K, V]]
func (a *HashDict[K, V]) Clone() *HashDict[K, V]
func (a *HashDict[K, V]) Clear()
```

### Set

We provide the `Set` interface to describe the element-unique collection type, and we provide `HashSet` as its implementation type.

```go
package set

type Set[T any] interface {
	iter.Collection[T]

	Contains(element T) bool
	Put(element T) bool
	Remove(element T) util.Opt[T]
	Clear()
}

func HashSetOf[T comparable](elements ...T) *HashSet[T]
func MakeHashSet[T comparable](capacity int) *HashSet[T]
func MakeHashSetWithHasher[T comparable](hasher func(data T) uint64, capacity int) *HashSet[T]
func HashSetFrom[T comparable](collection iter.Collection[T]) *HashSet[T]

type HashSet[T comparable] dict.HashDict[T, util.Void]
func (a *HashSet[T]) Count() int
func (a *HashSet[T]) Put(element T) bool
func (a *HashSet[T]) Remove(element T) util.Opt[T]
func (a *HashSet[T]) Contains(element T) bool
func (a *HashSet[T]) Iterator() iter.Iterator[T]
func (a *HashSet[T]) Clone() *HashSet[T]
func (a *HashSet[T]) Clear()
```

### Stack

We provide the `Stack` interface to describe the stack data structure and provide `ArrayStack` and `LinkedStack` as its implementation types.

```go
package stack

type Stack[T any] interface {
	iter.Collection[T]

	Push(element T)
	Pop() util.Opt[T]
	Peek() util.Ref[T]
	Clear()
}

func ArrayStackOf[T any](elements ...T) *ArrayStack[T]
func MakeArrayStack[T any](capacity int) *ArrayStack[T]
func ArrayStackFrom[T any](collection iter.Collection[T]) *ArrayStack[T]

type ArrayStack[T any] struct {}
func (a *ArrayStack[T]) Count() int
func (a *ArrayStack[T]) Push(element T)
func (a *ArrayStack[T]) PushAll(elements iter.Collection[T])
func (a *ArrayStack[T]) Pop() util.Opt[T]
func (a *ArrayStack[T]) Peek() util.Ref[T]
func (a *ArrayStack[T]) Iterator() iter.Iterator[T]
func (a *ArrayStack[T]) Clone() *ArrayStack[T]
func (a *ArrayStack[T]) Reserve(additional int)
func (a *ArrayStack[T]) Capacity() int
func (a *ArrayStack[T]) Clear()

func LinkedStackOf[T any](elements ...T) *LinkedStack[T]
func LinkedStackFrom[T any](collection iter.Collection[T]) *LinkedStack[T]

type LinkedStack[T any] struct {}
func (a *LinkedStack[T]) Count() int
func (a *LinkedStack[T]) Push(element T)
func (a *LinkedStack[T]) Pop() util.Opt[T]
func (a *LinkedStack[T]) Peek() util.Ref[T]
func (a *LinkedStack[T]) Iterator() iter.Iterator[T]
func (a *LinkedStack[T]) Clone() *ArrayStack[T]
func (a *LinkedStack[T]) Clear()
```

### Util

We have also introduced several convenient util types for use, and indeed gollection uses them as well. Including `Void`, `Pair`, `Ref`, `Opt`, `Result`.

```go
package util

type Void struct{}

func PairOf[T1 any, T2 any](f T1, s T2) Pair[T1, T2]

type Pair[T1 any, T2 any] struct {
	First  T1
	Second T2
}
func (a Pair[T1, T2]) Val() (T1, T2)

func RefOf[T any](v *T) Ref[T]

type Ref[T any] struct {}
func (a Ref[T]) Val() (v T, ok bool)
func (a Ref[T]) Get() T 
func (a Ref[T]) Set(v T) T
func (a Ref[T]) IsNil() bool
func (a Ref[T]) IsNotNil() bool

func Some[T any](a T) Opt[T]
func None[T any]() Opt[T]

type Opt[T any] struct {}
func (a Opt[T]) Val() (value T, ok bool)
func (a Opt[T]) Or(value T) T
func (a Opt[T]) OrDefault() T
func (a Opt[T]) OrPanic() T
func (a Opt[T]) IsSome() bool
func (a Opt[T]) IsNone() bool
func (a Opt[T]) IfSome(action func(value T))
func (a Opt[T]) IfNone(action func())
func (a Opt[T]) Next() Opt[T]

func Ok[T any](a T) Result[T]
func Err[T any](a error) Result[T]

type Result[T any] struct {}
func (a Result[T]) Val() (value T, err error)
func (a Result[T]) Or(value T) T
func (a Result[T]) OrDefault() (v T)
func (a Result[T]) OrPanic() T
func (a Result[T]) IsOk() bool
func (a Result[T]) IsErr() bool
func (a Result[T]) IfOk(action func(value T))
func (a Result[T]) IfErr(action func(err error))
```