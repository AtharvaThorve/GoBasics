package main

import (
	"fmt"
)

func main() {
	for x := 1; x <= 10; x++ {
		if x > 5 {
			fmt.Printf("Yup if %d\n", x)
		} else {
			fmt.Printf("Yup else %d\n", x)
		}
	}

	x := 1
	for x <= 10 { // Similar to while loop
		if x > 5 {
			fmt.Printf("Yup if %d\n", x)
		} else {
			fmt.Printf("Yup else %d\n", x)
		}
		x++
	}
	x = 1
	switch x {
	case 1: // No need to add break statement as it automatically breaks at the end of the case
		fmt.Printf("case 1\n")
	case 2:
		fmt.Printf("case 2\n")
	default:
		fmt.Printf("nocase\n")
	}
	switch {
	case x > 1: // This is a tagless switch so everything after case keyword is a boolean expression
		fmt.Printf("case 1\n")
	case x < -1:
		fmt.Printf("case 2\n")
	default:
		fmt.Printf("nocase\n")
	}
}
