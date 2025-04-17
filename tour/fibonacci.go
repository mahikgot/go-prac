package main

import "fmt"

func fibonacci() func() int {
	first, second := 0, 1
	ctr := 0
	return func() int {
		ctr++
		if ctr > 2 {
			ret := first + second
			first = second
			second = ret
			return ret
		}

		switch ctr {
		case 1:
			return first
		case 2:
			return second
		default:
			return -1
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
