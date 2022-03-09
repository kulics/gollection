package main

import (
	"fmt"

	. "github.com/kulics/gollection"
)

func main() {
	var stack Stack[int] = LinkedStackOf(1, 2, 3)
	stack = ArrayStackOf(1, 2, 3)
	fmt.Println(stack)
	var list List[int] = ArrayListOf(1, 2, 3, 4, 5)
	list = ArrayListOf(1, 2, 3, 4, 5)
	printItem := func(i int) {
		fmt.Println(i)
	}
	ForEach(printItem, list)
	list.Append(5)
	list.Append(4)
	list.Append(3)
	list.Append(2)
	list.Append(1)
	ForEach(printItem, Mapper(func(t int) int {
		return t * 3
	}, Filter(func(t int) bool {
		return t%2 == 0
	}, list)))
	list = ArrayListFrom[int](list)
	ForEach(printItem, list)
	fmt.Println(Product[int](list))
}
