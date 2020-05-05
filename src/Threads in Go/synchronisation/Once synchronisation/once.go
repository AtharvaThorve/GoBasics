package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var once sync.Once

func setup() {
	fmt.Println("init")
}

func doStuff() {
	once.Do(setup) // This ensures that initialisation only happens once and not multiple times even
	// though it is called by two separate goroutines. If one goroutine calls this method then other
	// goroutines are blocked automatically so initialisation only happens once.
	fmt.Println("Hello")
	wg.Done()
}

func main() {
	wg.Add(2)
	// Two goroutines but initialisation only once so sync.Once object is going to be used
	go doStuff()
	go doStuff()
	wg.Wait()
}
