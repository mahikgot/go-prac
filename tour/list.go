package main

import "fmt"

type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) Append(n *List[T]) {
	curr := l
	for {
		if curr.next == nil {
			curr.next = n
			curr.val = n.val
			return
		}
		curr = curr.next
	}
}

func (l *List[T]) String() string {
	return fmt.Sprintf("%v", l.val)
}

func main() {
	list := &List[int]{nil, 1}
	newList := &List[int]{nil, 2}
	list.Append(newList)
	fmt.Println(list)
}
