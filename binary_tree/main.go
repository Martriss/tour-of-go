// https://go.dev/tour/concurrency/8

package main

import (
	"fmt"
	"sort"
	"sync"

	"golang.org/x/tour/tree"
)

// type Tree struct {
//     Left  *Tree
//     Value int
//     Right *Tree
// }

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	ch <- t.Value
	Walk(t.Left, ch)
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch1)
		Walk(t1, ch1)
	}()

	go func() {
		defer close(ch2)
		Walk(t2, ch2)
	}()

	var slice1, slice2 []int

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for val := range ch1 {
			slice1 = append(slice1, val)
		}
	}()

	go func() {
		defer wg.Done()
		for val := range ch2 {
			slice2 = append(slice2, val)
		}
	}()

	wg.Wait()

	if len(slice1) != len(slice2) {
		return false
	}

	sort.Ints(slice1)
	sort.Ints(slice2)

	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
}
