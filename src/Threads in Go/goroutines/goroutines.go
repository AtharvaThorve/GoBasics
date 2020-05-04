package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func hello() {
	fmt.Println("Hello from goroutine function")
}

// This is a very simple example of using goroutines

func main() {
	// To make a goroutine just type the keyword go before the function call
	go hello() // This created a goroutine but then instantly control is given back to main goroutine
	// So to stop this from happening we need to make the main goroutine sleep so other goroutines can
	// execute
	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("Hello from main function")
}
