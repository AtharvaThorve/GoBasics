package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x int32 = 2
	var y int16 = 1
	x = int32(y)                                      // y is now converted to int32
	fmt.Printf("%T %T\n", x, y)                       // to get the type of variable use %T in Printf
	fmt.Println(reflect.TypeOf(x), reflect.TypeOf(y)) // Another way of doing the same thing but by using reflect package
}
