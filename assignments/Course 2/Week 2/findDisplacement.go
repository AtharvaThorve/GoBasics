package main

import (
	"fmt"
)

// GenDisplaceFunc ...A function that returns another function
func GenDisplaceFunc(a, v, s float64) func(float64) float64 {
	return func(t float64) float64 {
		return (a * t * t / 2) + (v * t) + (s)
	}
}

func main() {
	var a, v, s float64
	fmt.Println("Enter acceleration a, initial velocity v, and initial displacement s:")
	fmt.Scan(&a, &v, &s)
	fn := GenDisplaceFunc(a, v, s)
	var t float64
	fmt.Println("Now enter value for time:")
	fmt.Scan(&t)
	fmt.Println(fn(t))
}
