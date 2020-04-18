package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var z float32 = 1

func main() {
	var x float32 = 100 // IF type is not specified then it infers it automatically but it is a good
	// practice to specify the type of the variable
	fmt.Println(x + z) // If no value is assigned then 0 is default valuex

	// Shorten variable declarations uses ":=" syntax where the variable is declared and initialised
	// at with the value at the right hand side and type is inferred automatically
	// Short declaration is allowed only inside the function
	y := 150.1 // This becomes float64
	fmt.Println(y)

	x = 1.234e2
	fmt.Println(x)

	var z complex128 = complex(2, 3) // This is a complex number which is created using the complex function whose first arg is real part of number and second arg is imaginary part of number

	fmt.Println(z)

	a := "Hi, there"
	// strings package contains a lot of utility functions which are very usefull
	b := strings.Contains(a, "Hi")
	fmt.Println(b)

	c := "124i"
	// string conversion(strconv) contains a lot of usefull utility functions
	// This function converts ascii to integer which returns 2 values which are handled this way
	// one value is final answer and other is error. The below example throws an error as the
	// string to be converted is not an integer value. So d will be 0 and err will contain the
	// error
	d, err := strconv.Atoi(c)
	fmt.Printf("%d %T", d, err)

	// constant are declared using the keyword const. The value of a constant is inferred from right hand
	// side of the expression
	const l = 1.4
	// Multiple constants can be declared using the syntax
	const (
		m = 10
		n = "Hi"
	)

	// iota is a function to generate a set of related but distinct constants
	// Often represents a property which has several distinct values
	// Ex-
	// Days of the week
	// Months of the year
	// It is similar to enum in many languages
	type Grades int // Grades is now alias for int now
	// All grades are unique and we only care that they are different and not about their actual values
	// since iota is written for first value so it need not be repeated again and all other values are
	// assigned automatically
	const (
		A Grades = iota
		B
		C
		D
		F
	)

	fmt.Print("Enter string: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
