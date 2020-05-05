package main

import (
	"sync"
)

var wg sync.WaitGroup

func doStuff(c1, c2 chan int) {
	<-c1
	c2 <- 1
	wg.Done()
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	wg.Add(2)
	// This creates a deadlock condition
	go doStuff(ch1, ch2)
	go doStuff(ch2, ch1)
	wg.Wait()
}
