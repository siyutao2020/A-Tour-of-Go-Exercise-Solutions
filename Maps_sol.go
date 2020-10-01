package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	res := make(map[string]int)

	for _, word := range words {
		res[word] = res[word] + 1
	}

	return res
}

func main() {
	wc.Test(WordCount)
}
