package main

import (
	"fmt"
	"math"
)

// New interface with name shape2D and two methods in it
type shape2D interface {
	Area() float64
	Perimeter() float64
}

// A struct for triangle
// triangle implements all the methods in the interface so it is said that triangle type satisfies
// shape2D interface
type triangle struct {
	height float64
	base   float64
}

// A struct for circle
// circle implements all the methods in the interface so it is said that triangle type satisfies
// shape2D interface
type circle struct {
	radius float64
}

// Method Area with same signature as that of the one in the interface but with receiver type as triangle
func (t triangle) Area() float64 {
	return (0.5 * t.height * t.base)
}

// Method Perimeter with same signature as that of the one in the interface but with receiver type as
// triangle
func (t triangle) Perimeter() float64 {
	return (t.base * math.Sqrt(math.Pow(t.base, 2)+4*math.Pow(t.height, 2)))
}

// Method Area with same signature as that of the one in the interface but with receiver type as circle
func (c circle) Area() float64 {
	return (math.Pi * c.radius * c.radius)
}

// Method Perimeter with same signature as that of the one in the interface but with receiver type as
// circle
func (c circle) Perimeter() float64 {
	return (2 * math.Pi * c.radius)
}

type speaker interface {
	Speak()
}

type dog struct {
	name string
}

func (d *dog) Speak() {
	if d == nil {
		fmt.Println("<noise>")
	} else {
		fmt.Println(d.name)
	}
}

/***************************************************
Empty interface specifies no methods.
All types satisfy the empty interface.
Use it to a function to accept any type as parameter.
****************************************************/
func printMe(val interface{}) {
	fmt.Println(val)
}

/********************************************************************************************************
Example of type assertion
func DrawShape(s shape2D) bool {
	rect, ok := s.(rectangle) This line returns first the rectangle and other is the boolean value
	stating whether type assertion is correct
	if ok {
		DrawRect(rect)
	}
	tri, ok := s.(triangle)
	if ok {
		DrawTri(tri)
	}
}
Switch construct for the above
func DrawShape(s shape2D) bool {
	switch := sh := s.(type) { sh contains the type returned by the s.(type)
	case rectangle:
		DrawRect(sh)
	case triangle:
		DrawTri(sh)
	}
}
********************************************************************************************************/

// This is an error interface which is very common and built into many packages in go
type error interface {
	Error() string
}

func main() {
	// One way to use interface
	// Dynamic type of s1 is circle and dynamic value is c1
	var s1 shape2D
	c1 := circle{5}
	s1 = c1
	fmt.Println(s1.Area())
	fmt.Println(s1.Perimeter())

	// Second way to use interface
	s2 := dog{"Brian"}
	s2.Speak()

	// Nil dynamic value of interface
	var s3 speaker
	var d1 *dog
	s3 = d1
	s3.Speak()
}
