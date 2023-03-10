package main

import (
	"fmt"
	"unicode/utf8"

	"github.com/rivo/uniseg"
)

func main() {
	var m map[string]int
	if m == nil {
		fmt.Println("m is nil")
		fmt.Println("m get key ", m["key"])
	}

	s := "你好👍🏻"

	fmt.Println("Hello World", len([]rune(s)), utf8.RuneCountInString(s), uniseg.GraphemeClusterCount(s))
}
