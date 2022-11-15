package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacciOriginal() func() int {
	// Track previous two and current index
	state := [3]int{0, 1, 0}

	return func() int {
		// always increment index regardless
		defer func() { state[2] += 1 }()

		// special cases for first two indices
		// No computation possible or needed
		switch state[2] {
		case 0:
			return state[0]
		case 1:
			return state[1]
		default:
			computed := state[0] + state[1]
			state[0] = state[1]
			state[1] = computed
			return computed
		}
	}
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	curr, next := 0, 1
	start := true

	return func() int {
		// special case to include 0 output
		// No computation possible or needed
		if start {
			start = false
			return curr
		} else {
			curr, next = next, curr+next
			return curr
		}
	}
}

// Recursive version
func fib_r(n int) int {
	if n <= 1 {
		return 1
	}

	return fib_r(n-1) + fib_r(n-2)
}

func main() {
	f := fibonacci()
	for i := 0; i < 13; i++ {
		fmt.Println(f())
	}
}
