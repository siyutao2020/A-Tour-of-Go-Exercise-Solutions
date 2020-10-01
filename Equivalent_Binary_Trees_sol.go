package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	var walk func(t *tree.Tree)
	walk = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walk(t.Left)
		ch <- t.Value
		walk(t.Right)
	}
	walk(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		val1, ok1 := <-ch1
		val2, ok2 := <-ch2
		if ok1 != ok2 {
			return false
		}
		if ok1 == false {
			return true
		}
		if val1 != val2 {
			return false
		}
	}
}

func main() {
	if Same(tree.New(1), tree.New(1)) == true {
		fmt.Println("OK")
	} else {
		fmt.Println("no")
	}

	if Same(tree.New(1), tree.New(2)) == false {
		fmt.Println("OK")
	} else {
		fmt.Println("no")
	}
}
