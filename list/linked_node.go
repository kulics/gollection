package list

import . "github.com/kulics/gollection/tuple"

type node[T any] struct {
	Value T
	Next  *node[T]
}

func getNode[T any](n *node[T], i int) T {
	if i == 0 {
		return n.Value
	} else {
		return getNode(n.Next, i-1)
	}
}

func setNode[T any](n *node[T], i int, v T) T {
	if i == 0 {
		var oldValue = n.Value
		n.Value = v
		return oldValue
	} else {
		return setNode(n.Next, i-1, v)
	}
}

func getAndSetNode[T any](n *node[T], i int, set func(oldElement T) T) Pair[T, T] {
	if i == 0 {
		var oldValue = n.Value
		var newValue = set(oldValue)
		n.Value = newValue
		return PairOf(newValue, oldValue)
	} else {
		return getAndSetNode(n.Next, i-1, set)
	}
}

func prependNode[T any](n *node[T], v T) {
	if n == nil {
		*n = node[T]{v, nil}
	} else {
		*n = node[T]{v, n}
	}
}

func appendNode[T any](n *node[T], v T) {
	if n == nil {
		*n = node[T]{v, nil}
	} else {
		appendNode(n.Next, v)
	}
}

func insertNode[T any](n *node[T], pre *node[T], i int, e T) {
	if i == 0 {
		pre.Next = &node[T]{e, n}
	} else {
		insertNode(n.Next, n, i-1, e)
	}
}

func removeNode[T any](n *node[T], pre *node[T], i int) T {
	if i == 0 {
		var item = n.Value
		pre.Next = n.Next
		n.Next = nil
		return item
	} else {
		return removeNode(n.Next, n, i-1)
	}
}
