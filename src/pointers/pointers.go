package main

import "fmt"

func main() {
	var x int = 1
	var y int  // initial value is 0
	var p *int // p is a pointer to int

	fmt.Println(p) // at start it is nil
	p = &x         // p now points to x
	y = *p         // y is now 1 but addresses of x and y are different
	fmt.Println(x, y, &x, &y, *p)

	ptr := new(int) // returns a pointer to a variable which is at start intitialised to 0
	fmt.Println(ptr, *ptr)
	*ptr = 3 // Now the data at the address stroed in ptr is 3
	fmt.Println(ptr, *ptr) // The address remains the same but the data stored there changes
}
