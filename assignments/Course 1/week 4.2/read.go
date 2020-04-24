package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter name of the file to be read: ")
	var s string
	fmt.Scan(&s)
	type name struct {
		fname string
		lname string
	}
	slice := make([]name, 0)

	f, err := os.Open(s)
	defer f.Close()
	if err != nil {
		fmt.Printf("Error occured: %s", err)
		return
	}
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		srr := strings.Split(string(fileScanner.Text()), " ")
		var n name
		n.fname = srr[0]
		n.lname = srr[1]
		slice = append(slice, n)
	}

	for _, v := range slice {
		fmt.Println(v.fname, v.lname)
	}
}
