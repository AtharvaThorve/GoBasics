package main

import (
	"fmt"
	"math"
)

// Function can be considered as first class values similar to int, float, string, etc

// declaring a var as function
// This is how it is defined...firstly var keyword then the name of the var then type
var funcVar func(int) int

func incFunc(x int) int {
	return x + 1
}

func decFunc(x int) int {
	return x - 1
}

// functions as arguments
// argFunc is the function passed as argument to the function
func applyIt(argFunc func(int) int, val int) int {
	return argFunc(val)
}

// function with return value of a function
// The below function returns a function which takes in 2 float64 values and returns a float64 value
func makeDistOrigin(oX, oY float64) func(float64, float64) float64 {
	return func(x, y float64) float64 {
		return math.Sqrt(math.Pow(x-oX, 2) + math.Pow(y-oY, 2))
	}
}

// Variable number of arguments
func getMax(vals ...int) int {
	maxV := -1
	for _, val := range vals {
		if val > maxV {
			maxV = val
		}
	}
	return maxV
}

func main() {

	// Deferred function calls
	// These functions get executed after the surrounding functions are done
	// The function below will get executed at the end so useful for cleanup code
	defer fmt.Println("Goodbye")

	funcVar = incFunc // assigning the variable value of the function
	// Note no round paranthesis on the RHS stating funcVar is function itself and not the value returned
	// by the function
	fmt.Println(funcVar(1))

	fmt.Println(applyIt(incFunc, 2))
	fmt.Println(applyIt(decFunc, 2))

	// Anonymous function call
	// The below declaration consists of a anonymous function
	v := applyIt(func(x int) int {
		return x + 1
	}, 2)
	fmt.Println(v)

	Dist1 := makeDistOrigin(0, 0)
	Dist2 := makeDistOrigin(2, 2)
	// Now Dist1 and Dist2 are 2 different functions which calculate distance from different origins
	fmt.Println(Dist1(2, 2))
	fmt.Println(Dist2(2, 2))

	// Variadic slice argument
	fmt.Println(getMax(1, 3, 7, 4))
	vSlice := []int{1, 3, 7, 4}
	fmt.Println(getMax(vSlice...)) // ... unzips the values stored inside the list

	// Arguments of deferred functions are evaluated immediately but function is call later
	i := 1
	defer fmt.Println(i + 1)
	i++
	fmt.Println("Hello!")
}
