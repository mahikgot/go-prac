package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	res := make(map[string]int)
	for _, str := range strings.Fields(s) {
		res[str] = res[str] + 1
	}
	return res
}

func main() {
	wc.Test(WordCount)
}
