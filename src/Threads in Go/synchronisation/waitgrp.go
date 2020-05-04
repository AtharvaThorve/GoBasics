package main

import (
	"fmt"
	"sync"
)

func foo1(wg *sync.WaitGroup) {
	fmt.Println("Hello from goroutine 1")
	wg.Done() // This method tells the WaitGroup that this goroutine has completed execution
}

func foo2(wg *sync.WaitGroup) {
	fmt.Println("Hello from goroutine 2")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup // sync package consists lot of constructs for synchronisation
	// WaitGroup waits for specific number of goroutines to complete before executing a specific
	// goroutine
	wg.Add(2) // This value tells number of goroutines to wait for before execution
	go foo1(&wg)
	go foo2(&wg)
	wg.Wait() // This method tell the goroutine to wait before all the specified goroutines hav completed
	// their execution
	fmt.Println("Hello from main")
}
