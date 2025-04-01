package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	prev, cur := 0, 1
	return func() (ret int) {
		ret, prev, cur = prev, cur, prev+cur
		return
	}
}

func main() {
	f := fibonacci()
	for range 10 {
		fmt.Println(f())
	}
}
