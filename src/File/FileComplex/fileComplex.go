package main

import (
	"fmt"
	"os"
)

func main() {
	// Opening and reading
	// Open would return the file handler f
	f, err := os.Open("test.txt")
	barr := make([]byte, 10)

	// This will fill the byte array with first 10 bytes of the file

	/**********************************************
	If the read function is called again using the same byte array then the array is filled with the
	next 10 bytes of the file. This means the read head actually moves until the file is closed using
	the close function
	**********************************************/
	// nb contains number of bytes read and all the read data is stored in the byte array we pass
	nb, err := f.Read(barr)
	fmt.Println("Data read, number of characters read and error:", string(barr), nb, err)

	nb, err = f.Read(barr)
	fmt.Printf("The read head moved here\n")
	fmt.Println("Data read, number of characters read and error:", string(barr), nb, err)
	f.Close()

	// New file or appending to same file
	f, err = os.Create("outfile.txt")
	barr = []byte{1, 2, 3}
	nb, err = f.Write(barr)
	nb, err = f.WriteString("Hi")

	f, err = os.OpenFile("test.txt", os.O_APPEND, 0644)
	defer f.Close() // defer statements delay the execution of the function or method or an anonymous
	// method until the nearby functions returns.
	f.WriteString("\nHello, World!\n")
}
