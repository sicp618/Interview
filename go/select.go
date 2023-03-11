package main

import (
	"fmt"
	"time"
)

func fib(c, quit chan int) {
	x, y := 0, 1
	// sleep 1s
	for {
		select {
		// case x <- c <- x:
		case <- c:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		default:
			fmt.Println("default")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
			fmt.Println("aa")
		}
		// quit <- 0
	}()
	fib(c, quit)
	time.Sleep(time.Second)
}