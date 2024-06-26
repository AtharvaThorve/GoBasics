package main

import "fmt"

func main() {
	var x [5]int
	x[0] = 2
	fmt.Println(x)

	primes := [5]int{2, 3, 5, 7, 11} // [...] syntax can be used here too
	y := [...]int{1, 2, 3, 4, 5}     // Explicitly specifying size is not necessary
	// fmt.Println(y + primes)

	for i, v := range primes {
		fmt.Printf("ind: %d, val: %d ", i, v)
	}

	fmt.Println()

	for i := 0; i < len(y); i++ {
		fmt.Printf("ind: %d, val: %d ", i, y[i])
	}

	fmt.Println()

	for _, v := range y {
		fmt.Printf("val: %d ", v) // If index is not required
	}
}
