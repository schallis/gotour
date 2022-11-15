//https://stackoverflow.com/questions/12224042/go-tour-exercise-7-binary-trees-equivalence

package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	// NOTE: Prefer using defer as idiomatic
	defer close(ch)

	var walker func(t *tree.Tree)
	walker = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walker(t.Left)
		ch <- t.Value
		walker(t.Right)
	}
	walker(t)

	// The following attempt doesn't read in the right order
	// ch <- t.Value

	// switch {
	// case t.Left != nil:
	// 	Walk(t.Left, ch)
	// case t.Right != nil:
	// 	Walk(t.Right, ch)
	// default:
	// 	// 	fmt.Println("closing...")
	// 	// 	close(ch)
	// }
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		val1, ok1 := <-ch1
		val2, ok2 := <-ch2

		//fmt.Printf("%v %v %v\n", val1, val2, val1 == val2)

		if !ok1 || !ok2 {
			break
		}

		if val1 != val2 {
			return false
		}
	}

	return true
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(1)

	fmt.Printf("%v\n", Same(t1, t2))
}
