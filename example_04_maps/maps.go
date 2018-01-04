package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wc := make(map[string]int)
	for _, word := range strings.Fields(s) {
		c, _ := wc[word]
		wc[word] = c + 1
	}

	return wc
}

func main() {
	wc.Test(WordCount)
}

