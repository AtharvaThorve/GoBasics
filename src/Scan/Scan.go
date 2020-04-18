package main

import "fmt"

func main() {
	var appleNum, bananaNum int
	var name string
	fmt.Printf("Enter number of apples, bananas and name: ")
	// We need to pass the pointer to the variable where data must be stored
	num, err := fmt.Scan(&appleNum, &bananaNum, &name)
	// Scan returns number of tokens that user inputs and error
	fmt.Println(appleNum, bananaNum, name, num, err)
}
