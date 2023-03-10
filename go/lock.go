package main

import (
	"fmt"
	"sync"
)

func main() {
	print()
}

func print() {
	typeCount := 3
	loopCount := 10
	dogChan := make(chan struct{}, 1)
	catChan := make(chan struct{}, 1)
	fishChan := make(chan struct{}, 1)
	var wg = sync.WaitGroup{}
	wg.Add(typeCount * loopCount)
	for i := 0; i < loopCount; i++ {
		go cat(catChan, fishChan, &wg)
		go dog(dogChan, catChan, &wg)
		go fish(fishChan, dogChan, &wg)
	}
	dogChan <- struct{}{}
	wg.Wait()
}

func dog(ch, nextCh chan struct{}, wg *sync.WaitGroup) {
	<-ch
	fmt.Println("dog")
	nextCh <- struct{}{}
	wg.Done()
}

func cat(ch, nextCh chan struct{}, wg *sync.WaitGroup) {
	<-ch
	fmt.Println("cat")
	nextCh <- struct{}{}
	wg.Done()
}

func fish(ch, nextCh chan struct{}, wg *sync.WaitGroup) {
	<-ch
	fmt.Println("fish")
	nextCh <- struct{}{}
	wg.Done()
}
