package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	distinctWords := make(map[string]int)

	for _, word := range words {
		existing := distinctWords[word]
		distinctWords[word] = existing + 1
	}

	return distinctWords
}

func main() {
	wc.Test(WordCount)
}
