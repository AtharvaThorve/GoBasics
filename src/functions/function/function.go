package main

import (
	"fmt"
)

func foo1() {
	fmt.Println("foo1 with no parameters and return value")
}

func foo2(x int) int {
	fmt.Println("foo2 with 1 parameter and 1 return value")
	return x + 1
}

func foo3(x int) (int, int) {
	fmt.Println("foo3 with 1 parameter and 2 return values")
	return x, x + 1
}

func foo4(x, y int) int {
	fmt.Println("foo4 with 2 parameters and 1 return value")
	return x * y
}

func main() {
	// Execute the code to see what each function does
	foo1()
	a := foo2(1)
	fmt.Println("foo2 result", a)
	a, b := foo3(2)
	fmt.Println("foo3 result", a, b)
	a = foo4(3, 2)
	fmt.Println("foo4 result", a)
}
