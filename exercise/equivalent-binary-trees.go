package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Structure of the Tree struct.
// type Tree struct {
//     Left  *Tree
//     Value int
//     Right *Tree
// }

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	// Declare closure to call itself recursive
	var walker func(*tree.Tree)
	walker = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walker(t.Left)
		ch <- t.Value
		walker(t.Right)
	}
	walker(t)
	close(ch)
}

// Same determines whether trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		switch {
		case !ok1, !ok2:
			// In the case of...
			// 1. Two channels were closed at the same time -> true
			// 2. Two channels were closed at different times -> false
			return ok1 == ok2
		case v1 != v2:
			return false
		}
	}
}

func main() {
	// Test for Walk
	// ch := make(chan int)
	// go Walk(tree.New(1), ch)
	// for v := range ch {
	// 	fmt.Println(v)
	// }

	// Test for Same
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
