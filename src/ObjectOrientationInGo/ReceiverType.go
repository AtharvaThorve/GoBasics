package main

import (
	"fmt"
	"math"
)

/****************************************************************************************************
If you associate a method with a receiver type then you have to make sure that type is defined in the
same package as of the method.
This cannot be done with built in types like strings or ints, etc.
If name of a variable or a function starts with a capital letter then it is public in nature. If it
starts with lowercase letter then it is private
*****************************************************************************************************/

// MyInt of type int
type MyInt int

// Double function whose receiver type is MyInt. This function is now associated with MyInt
func (mi MyInt) Double() int {
	return int(mi * 2)
}

type Point struct {
	x float64
	y float64
}

func (p *Point) DistToOrigin() float64 {
	t := math.Pow(p.x, 2) + math.Pow(p.y, 2)
	return math.Sqrt(t)
}

func (p *Point) InitMe(x, y float64) {
	p.x = x
	p.y = y
}

func (p *Point) Scale(v float64) {
	p.x = p.x * v
	p.y = p.y * v
}

func (p *Point) PrintMe() {
	fmt.Println(p.x, p.y)
}

func main() {
	v := MyInt(3)
	fmt.Println(v.Double())

	var p1 Point
	p1.InitMe(3, 4)
	fmt.Println(p1.DistToOrigin())
	p1.Scale(2)
	p1.PrintMe()
}
