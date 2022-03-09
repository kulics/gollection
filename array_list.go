package gollection

func ArrayListOf[T any](elements ...T) ArrayList[T] {
	var array = make([]T, len(elements))
	copy(array, elements)
	return ArrayList[T]{&struct {
		elements []T
		size     int
	}{array, len(elements)}}
}

func MakeArrayList[T any](capacity int) ArrayList[T] {
	return ArrayList[T]{&struct {
		elements []T
		size     int
	}{make([]T, capacity), 0}}
}

func ArrayListFrom[T any, I Collection[T]](collection I) ArrayList[T] {
	var size = collection.Size()
	var array = make([]T, size)
	ForEach(func(item Pair[int, T]) {
		array[item.First] = item.Second
	}, WithIndex[T](collection))
	return ArrayList[T]{&struct {
		elements []T
		size     int
	}{array, size}}
}

type ArrayList[T any] struct {
	inner *struct {
		elements []T
		size     int
	}
}

func (a ArrayList[T]) Prepend(element T) {
	if len(a.inner.elements) < a.inner.size+1 {
		a.grow(1)
	}
	copy(a.inner.elements[1:], a.inner.elements[0:])
	a.inner.elements[0] = element
}

func (a ArrayList[T]) PrependAll(elements Collection[T]) {
	var additional = elements.Size()
	if len(a.inner.elements) < a.inner.size+additional {
		a.grow(additional)
	}
	copy(a.inner.elements[additional:], a.inner.elements[0:])
	var i = 0
	ForEach(func(item T) {
		a.inner.elements[i] = item
		a.inner.size++
		i++
	}, elements)
}

func (a ArrayList[T]) Append(element T) {
	if len(a.inner.elements) < a.inner.size+1 {
		a.grow(1)
	}
	a.inner.elements[a.inner.size] = element
	a.inner.size++
}

func (a ArrayList[T]) AppendAll(elements Collection[T]) {
	var additional = elements.Size()
	if len(a.inner.elements) < a.inner.size+additional {
		a.grow(additional)
	}
	var i = a.inner.size
	ForEach(func(item T) {
		a.inner.elements[i] = item
		a.inner.size++
		i++
	}, elements)
}

func (a ArrayList[T]) Insert(index int, element T) bool {
	if index < 0 || index > a.inner.size {
		return false
	}
	if len(a.inner.elements) < a.inner.size+1 {
		a.grow(1)
	}
	copy(a.inner.elements[index+1:], a.inner.elements[index:])
	a.inner.elements[index] = element
	return true
}

func (a ArrayList[T]) InsertAll(index int, elements Collection[T]) bool {
	if index < 0 || index > a.inner.size {
		return false
	}
	var additional = elements.Size()
	if len(a.inner.elements) < a.inner.size+additional {
		a.grow(additional)
	}
	copy(a.inner.elements[index+additional:], a.inner.elements[index:])
	var i = index
	ForEach(func(item T) {
		a.inner.elements[i] = item
		a.inner.size++
		i++
	}, elements)
	return true
}

func (a ArrayList[T]) Remove(index int) Option[T] {
	if a.isOutOfBounds(index) {
		return None[T]()
	}
	var removed = a.inner.elements[index]
	copy(a.inner.elements[:index], a.inner.elements[index+1:])
	var emptyValue T
	a.inner.elements[a.inner.size-1] = emptyValue
	a.inner.size--
	return Some(removed)
}

func (a ArrayList[T]) Reserve(additional int) {
	var addable = len(a.inner.elements) - a.inner.size
	if addable < additional {
		a.grow(additional - addable)
	}
}

func (a ArrayList[T]) Get(index int) T {
	if v, ok := a.TryGet(index).Get(); ok {
		return v
	}
	panic(OutOfBounds)
}

func (a ArrayList[T]) Set(index int, newElement T) T {
	if v, ok := a.TrySet(index, newElement).Get(); ok {
		return v
	}
	panic(OutOfBounds)
}

func (a ArrayList[T]) GetAndSet(index int, set func(oldElement T) T) Pair[T, T] {
	if a.isOutOfBounds(index) {
		panic(OutOfBounds)
	}
	var oldElement = a.inner.elements[index]
	var newElement = set(oldElement)
	return PairOf(newElement, oldElement)
}

func (a ArrayList[T]) TryGet(index int) Option[T] {
	if a.isOutOfBounds(index) {
		return None[T]()
	}
	return Some(a.inner.elements[index])
}

func (a ArrayList[T]) TrySet(index int, newElement T) Option[T] {
	if a.isOutOfBounds(index) {
		return None[T]()
	}
	var oldElement = a.inner.elements[index]
	a.inner.elements[index] = newElement
	return Some(oldElement)
}

func (a ArrayList[T]) Clear() {
	var emptyValue T
	for i := 0; i < a.inner.size; i++ {
		a.inner.elements[i] = emptyValue
	}
	a.inner.size = 0
}

func (a ArrayList[T]) Size() int {
	return a.inner.size
}

func (a ArrayList[T]) IsEmpty() bool {
	return a.inner.size == 0
}

func (a ArrayList[T]) Iter() Iterator[T] {
	return &arrayListIterator[T]{-1, a}
}

func (a ArrayList[T]) isOutOfBounds(index int) bool {
	if index < 0 || index >= a.inner.size {
		return true
	}
	return false
}

func (a ArrayList[T]) grow(minCapacity int) {
	var newSize = int(float64(len(a.inner.elements)) * 1.5)
	if newSize < minCapacity {
		newSize = minCapacity
	}
	var newSource = make([]T, newSize)
	copy(newSource, a.inner.elements)
	a.inner.elements = newSource
}

type arrayListIterator[T any] struct {
	index  int
	source ArrayList[T]
}

func (a *arrayListIterator[T]) Next() Option[T] {
	if a.index < a.source.Size()-1 {
		a.index++
		return a.source.TryGet(a.index)
	}
	return None[T]()
}

func (a *arrayListIterator[T]) Iter() Iterator[T] {
	return a
}
