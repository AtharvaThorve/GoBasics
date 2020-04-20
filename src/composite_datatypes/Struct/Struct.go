package main

import "fmt"

func main() {
	// Aggregate datatype
	// Groups together other objects of arbitary type
	type Person struct {
		name    string
		address string
		phone   string
	}

	var p1 Person
	fmt.Println(p1)
	p1.name = "Joe"
	fmt.Println(p1)

	// Another way to initialise struct
	p2 := new(Person) // new will return a pointer to the struct so use * to dereference it
	p2.name = "Jane"
	fmt.Println(*p2)

	// Another way to initialise struct
	structLiteral := Person{name: "Joe", address: "a St.", phone: "123"}
	fmt.Println(structLiteral)
}
