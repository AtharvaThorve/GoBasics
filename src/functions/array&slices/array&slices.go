package main

import (
	"fmt"
)

func foo1(x [3]int) {
	fmt.Println("Pass by value")
	x[0] = x[0] + 1
}

func foo2(x *[3]int) {
	fmt.Println("Pass by reference")
	(*x)[0] = (*x)[0] + 1
}

func foo3(sli []int) {
	fmt.Println("Passing slice")
	sli[0] = sli[0] + 1
}

func main() {
	/*******************************
	Call by value
	********************************/
	a := [3]int{1, 2, 3}
	fmt.Println("Value of a before function call", a)
	foo1(a)
	fmt.Println("Value of a after function call", a)

	/*******************************
	Call by reference
	********************************/
	fmt.Println("Value of a before function call", a)
	foo2(&a)
	fmt.Println("Value of a after function call", a)
	// Dont use pass by reference but start using slices

	/*******************************
	Passing slice
	********************************/
	b := []int{1, 2, 3}
	fmt.Println("Value of slice b before function call", b)
	foo3(b)
	fmt.Println("value of slice b after function call", b)
}
