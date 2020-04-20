package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	// For file handling multiple packages are allowed but ioutil is good for basics
	// It returns 2 things data in a byte array format and error
	data, _ := ioutil.ReadFile("test.txt")
	fmt.Println("Data returned by the ReadFile:", string(data))
	// The read file function reads the complete file irrespective of spaces or new line characters
	// read file is not good for large size files

	data = []byte("Hello, World")
	ioutil.WriteFile("outfile.txt", data, 0777)
	// The write file has three arguments. First 2 are self explanatory. Third argument is the permission
	// argument which gives the program to create a new file and write to it

	ioutil.WriteFile("test.txt", data, 0777)
	// This will overwrite everything that is already present in the file
}
