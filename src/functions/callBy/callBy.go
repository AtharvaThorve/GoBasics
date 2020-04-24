package main

import "fmt"

func foo1(y int) {
	fmt.Println("Call by value")
	y = y + 1
}

func foo2(y *int) {
	fmt.Println("Call by reference")
	*y = *y + 1
}

func main() {
	/*******************************
	Call by value
	********************************/
	x := 1
	fmt.Println("Value of x before function call", x)
	foo1(x)
	fmt.Println("Value of x after function call", x)

	/*******************************
	Call by reference
	********************************/
	fmt.Println("Value of x before function call", x)
	foo2(&x)
	fmt.Println("Value of x after function call", x)

}
