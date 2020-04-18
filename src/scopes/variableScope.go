package main

import "fmt"

var x int = 4

func f() {
	// var x int = 4   by adding the variable here the scope is limited only to the function block
	fmt.Printf("%d\n", x)
}

func g() {
	fmt.Printf("%d", x)
}

func main() {
	f()
	g()
}
