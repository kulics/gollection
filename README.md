# gollection

A generic generic collection library based on go's generic implementation.

## Core Interfaces

```go
type Sequence[T any] interface {
	Iterator() Iterator[T]
}

type Iterator[T any] interface {
	Next() option.Option[T]
}
```

gollection applies the iterator pattern design, the core interface consists of Iterator and Sequence.

Iterator is responsible for providing iterative functionality, Iterator is unidirectional and lazy, each call to next will only return one result.

Sequence is responsible for providing Iterator, the implementation type determines whether the provided Iterator is reusable.

The inert traversal feature allows the combination of higher-order functions without significant overhead and can provide a richer combination of functions.

Here is a simple example of direct traversal:

```go
func printAll[T any](it Sequence[T]) {
	iter := it.Iterator()
	for v, ok := iter.Next().Val(); ok;  v, ok = iter.Next().Val() {
		println(v)
	}
}
```

## Streaming operations

gollection provides a rich set of stream manipulation functions that can be used in combination with any Sequence.

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
	ForEach(show, Map(square, Filter(even, Slice[int]([]int{1, 2, 3, 4, 5, 6, 7}))))
    // Result:
    // 4
    // 16
    // 36
}
```

## ToString and ToSlice

In order to make go's native string and slice also Sequence, we have introduced `ToSlice` and `ToString` to make these two types implement the interface.

```go
var str = "Hello, world!"
var sli = []int{1, 2, 3}
Count(ToString(str)) // 13
Count(ToSlice(sli)) // 3
```

## Collection

We define a unified collection type interface to describe more information than iterators to facilitate performance optimization.

```go
type Collection[T any] interface {
	Sequence[T]

	Count() int
}

func IsEmpty[T any](c Collection[T]) bool
func IsNotEmpty[T any](c Collection[T]) bool
func ToSlice[T any](c Collection[T]) Slice[T]
```

### List and LinkedList

We provide the `List` and `LinkedList` types to describe the ordered sequences.

### Dict

We provide the `Dict` type to describe the mapping type.

### Set

We provide the `Set` type to describe the element-unique collection type.

### Stack

We provide the `Stack` type to describe the stack data structure.

### Others

We have also introduced several convenient util types for use, and indeed gollection uses them as well. Including `Ref`, `Option`, `Result`.