package gollection

func LinkedListOf[T any](elements ...T) LinkedList[T] {
	var inner = &linkedList[T]{0, nil, nil}
	var list = LinkedList[T]{inner}
	for _, v := range elements {
		list.Append(v)
	}
	return list
}

func LinkedListFrom[T any, I Collection[T]](collection I) LinkedList[T] {
	var list = LinkedListOf[T]()
	list.AppendAll(collection)
	return list
}

type LinkedList[T any] struct {
	inner *linkedList[T]
}

type linkedList[T any] struct {
	size  int
	first *twoWayNode[T]
	last  *twoWayNode[T]
}

type twoWayNode[T any] struct {
	value T
	next  *twoWayNode[T]
	prev  *twoWayNode[T]
}

func (a LinkedList[T]) GetFirst() T {
	if v, ok := a.TryGetFirst().Get(); ok {
		return v
	}
	panic(OutOfBounds)
}

func (a LinkedList[T]) TryGetFirst() Option[T] {
	if first := a.inner.first; first != nil {
		return Some(first.value)
	}
	return None[T]()
}

func (a LinkedList[T]) RemoveFirst() T {
	var first = a.inner.first
	if first == nil {
		panic(OutOfBounds)
	}
	return a.unlinkFirst(first)
}

func (a LinkedList[T]) GetLast() T {
	if v, ok := a.TryGetLast().Get(); ok {
		return v
	}
	panic(OutOfBounds)
}

func (a LinkedList[T]) TryGetLast() Option[T] {
	if last := a.inner.last; last != nil {
		return Some(last.value)
	}
	return None[T]()
}

func (a LinkedList[T]) RemoveLast() T {
	var last = a.inner.last
	if last == nil {
		panic(OutOfBounds)
	}
	return a.unlinkLast(last)
}

func (a LinkedList[T]) Prepend(element T) {
	a.linkFirst(element)
}

func (a LinkedList[T]) PrependAll(elements Collection[T]) {
	a.InsertAll(0, elements)
}

func (a LinkedList[T]) Append(element T) {
	a.linkLast(element)
}

func (a LinkedList[T]) AppendAll(elements Collection[T]) {
	a.InsertAll(a.inner.size, elements)
}

func (a LinkedList[T]) Insert(index int, element T) {
	if a.isOutOfBounds(index) {
		panic(OutOfBounds)
	}
	if index == 0 {
		a.linkLast(element)
	} else {
		a.linkBefore(element, a.at(index))
	}
}

func (a LinkedList[T]) InsertAll(index int, elements Collection[T]) {
	if a.isOutOfBounds(index) {
		panic(OutOfBounds)
	}
	var size = elements.Size()
	if size == 0 {
		return
	}
	var pred, succ *twoWayNode[T]
	if index == a.inner.size {
		succ = nil
		pred = a.inner.last
	} else {
		succ = a.at(index)
		pred = succ.prev
	}
	var iter = elements.Iter()
	for v, ok := iter.Next().Get(); ok; v, ok = iter.Next().Get() {
		var newNode = &twoWayNode[T]{value: v, prev: pred, next: nil}
		if pred == nil {
			a.inner.first = newNode
		} else {
			pred.next = newNode
		}
		pred = newNode
	}
	if succ == nil {
		a.inner.last = pred
	} else {
		pred.next = succ
		succ.prev = pred
	}
	a.inner.size += size
}

func (a LinkedList[T]) Remove(index int) T {
	if a.isOutOfBounds(index) {
		panic(OutOfBounds)
	}
	return a.unlink(a.at(index))
}

func (a LinkedList[T]) Get(index int) T {
	if v, ok := a.TryGet(index).Get(); ok {
		return v
	}
	panic(OutOfBounds)
}

func (a LinkedList[T]) Set(index int, newElement T) T {
	if v, ok := a.TrySet(index, newElement).Get(); ok {
		return v
	}
	panic(OutOfBounds)
}

func (a LinkedList[T]) GetAndSet(index int, set func(oldElement T) T) Pair[T, T] {
	if a.isOutOfBounds(index) {
		panic(OutOfBounds)
	}
	var x = a.at(index)
	var oldValue = x.value
	x.value = set(x.value)
	return PairOf(x.value, oldValue)
}

func (a LinkedList[T]) TryGet(index int) Option[T] {
	if a.isOutOfBounds(index) {
		return None[T]()
	}
	return Some(a.at(index).value)
}

func (a LinkedList[T]) TrySet(index int, newElement T) Option[T] {
	if a.isOutOfBounds(index) {
		return None[T]()
	}
	var x = a.at(index)
	var oldValue = x.value
	x.value = newElement
	return Some(oldValue)
}

func (a LinkedList[T]) Clear() {
	for x := a.inner.first; x != nil; {
		var next = x.next
		var empty T
		x.value = empty
		x.next = nil
		x.prev = nil
		x = next
	}
	a.inner.first = nil
	a.inner.last = nil
	a.inner.size = 0
}

func (a LinkedList[T]) Size() int {
	return a.inner.size
}

func (a LinkedList[T]) IsEmpty() bool {
	return a.inner.size == 0
}

func (a LinkedList[T]) Iter() Iterator[T] {
	return &linkedListIterator[T]{a.inner.first}
}

func (a LinkedList[T]) ToSlice() []T {
	var arr = make([]T, a.Size())
	ForEach(func(t T) {
		arr = append(arr, t)
	}, a.Iter())
	return arr
}

func (a LinkedList[T]) Clone() LinkedList[T] {
	return LinkedListFrom[T](a)
}

func (a LinkedList[T]) isOutOfBounds(index int) bool {
	if index < 0 || index >= a.inner.size {
		return true
	}
	return false
}

func (a LinkedList[T]) at(index int) *twoWayNode[T] {
	if index < (a.inner.size >> 1) {
		var x = a.inner.first
		for i := 0; i < index; i++ {
			x = x.next
		}
		return x
	} else {
		var x = a.inner.last
		for i := a.inner.size - 1; i > index; i-- {
			x = x.prev
		}
		return x
	}
}

func (a LinkedList[T]) linkFirst(element T) {
	var first = a.inner.first
	var newNode = &twoWayNode[T]{value: element, prev: nil, next: first}
	a.inner.first = newNode
	if first == nil {
		a.inner.last = newNode
	} else {
		first.prev = newNode
	}
	a.inner.size++
}

func (a LinkedList[T]) linkLast(element T) {
	var last = a.inner.last
	var newNode = &twoWayNode[T]{value: element, next: nil, prev: last}
	a.inner.last = newNode
	if last == nil {
		a.inner.first = newNode
	} else {
		last.next = newNode
	}
	a.inner.size++
}

func (a LinkedList[T]) linkBefore(element T, succ *twoWayNode[T]) {
	var pred = succ.prev
	var newNode = &twoWayNode[T]{value: element, prev: pred, next: succ}
	succ.prev = newNode
	if pred == nil {
		a.inner.first = newNode
	} else {
		pred.next = newNode
	}
	a.inner.size++
}

func (a LinkedList[T]) unlink(x *twoWayNode[T]) T {
	var element = x.value
	var next = x.next
	var prev = x.prev
	if prev == nil {
		a.inner.first = next
	} else {
		prev.next = next
		x.prev = nil
	}

	if next == nil {
		a.inner.last = prev
	} else {
		next.prev = prev
		x.next = nil
	}
	var empty T
	x.value = empty
	a.inner.size--
	return element
}

func (a LinkedList[T]) unlinkFirst(x *twoWayNode[T]) T {
	var element = x.value
	var next = x.next
	var empty T
	x.value = empty
	x.next = nil
	a.inner.first = next
	if next == nil {
		a.inner.last = nil
	} else {
		next.prev = nil
	}
	a.inner.size--
	return element
}

func (a LinkedList[T]) unlinkLast(x *twoWayNode[T]) T {
	var element = x.value
	var prev = x.prev
	var empty T
	x.value = empty
	x.prev = nil
	a.inner.last = prev
	if prev == nil {
		a.inner.first = nil
	} else {
		prev.next = nil
	}
	a.inner.size--
	return element
}

type linkedListIterator[T any] struct {
	current *twoWayNode[T]
}

func (a *linkedListIterator[T]) Next() Option[T] {
	if a.current != nil {
		var item = a.current.value
		a.current = a.current.next
		return Some(item)
	}
	return None[T]()
}
