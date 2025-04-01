package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) (m map[string]int) {
	m = make(map[string]int)

	w := strings.Fields(s)
	for _, v := range w {
		m[string(v)] += 1
	}
	return
}

func main() {
	wc.Test(WordCount)
}
