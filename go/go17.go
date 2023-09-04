package main

import (
	"fmt"
	"os"
	"os/signal"
)

// use go vet to check for errors
func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	fmt.Println("Waiting for signal")
	<-c
	fmt.Println("Got signal")
}
