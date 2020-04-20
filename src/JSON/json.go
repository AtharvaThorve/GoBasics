package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Employee struct {
		Name   string
		Age    int
		Salary int
	}

	// Object of the struct Employee
	empObj1 := Employee{Name: "Rachel", Age: 24, Salary: 344444}

	// Marshal function returns a json representation as a byte array
	// It returns two things first is the byte array and the other is the error
	emp, _ := json.Marshal(empObj1)
	fmt.Println("Struct and its equivalent json form:", empObj1, string(emp))

	// Unmarshal function works opposite of marshal function
	var empObj2 Employee
	json.Unmarshal(emp, &empObj2)
	fmt.Println("Unmarshalling the byte array recieved earlier from marshal function:", empObj2)
}
