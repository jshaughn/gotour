package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if nil == t {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < 10; i++ {
		v1, v2 := <-ch1, <-ch2
		fmt.Println(v1, v2)
		if v1 != v2 {
			return false
		}
	}
	return true
}

func main() {
	if Same(tree.New(1), tree.New(1)) && !Same(tree.New(1), tree.New(2)) {
		fmt.Println("OK!")
	} else {
		fmt.Println("Fail...")
	}
}

