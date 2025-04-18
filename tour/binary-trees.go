package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func walkAndClose(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	c1, c2 := make(chan int, 10), make(chan int, 10)
	go walkAndClose(t1, c1)
	go walkAndClose(t2, c2)
	for {
		x, ok1 := <-c1
		y, ok2 := <-c2
		if x != y {
			return false
		}
		if !ok1 && !ok2 {
			return true
		}
		if !ok1 || !ok2 {
			return false
		}
	}
}

func main() {
	t := tree.New(1)
	s := tree.New(2)
	fmt.Println(Same(t, s))

	t = tree.New(1)
	s = tree.New(1)
	fmt.Println(Same(t, s))
}
