package main

import (
	"fmt"
	"sync/atomic"
)

// # Title
// 1. Item 1
// 2. Item 2
// find main features [google]
// [google]: http://google.com
func RunAtomic() {
	var i1 int64
	atomic.AddInt64(&i1, 10)
	fmt.Println(atomic.LoadInt64(&i1))

	var i2 atomic.Int64
	i2.Add(10)
	fmt.Println(i2.Load())
}

func main() {
	fmt.Println("Hello World")
	RunAtomic()
}