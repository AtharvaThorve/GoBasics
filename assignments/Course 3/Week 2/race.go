package main

import (
	"fmt"
	"time"
)

func appendToSlice1(a []int) {
	for i := 10; i < 20; i++ {
		time.Sleep(10 * time.Millisecond)
		a = append(a, i)
	}
	fmt.Println(a)
}

func appendToSlice2(a []int) {
	for i := 1; i < 10; i++ {
		time.Sleep(10 * time.Millisecond)
		a = append(a, i)
	}
	fmt.Println(a)
}

func main() {
	fmt.Println("Started main goroutine")
	a := make([]int, 0, 3)
	go appendToSlice1(a)
	go appendToSlice2(a)
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("Main routine ended")
}

/********************************************************************************************************
In the above code there are 2 goroutines running other than the main goroutine known as appendToSlice1
and appendToSlice2. The functions as the name suggests append elements to a slice which is common between
both the goroutines. Execute the code to see the actual results.
Also instead of go run race.go if go run -race race.go is used then more details of the race conditio are
visible.
********************************************************************************************************/
