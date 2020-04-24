package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	slice := make([]int, 0, 3)
	var s string
	var i int
	for {
		fmt.Print("Enter the number to be appended or press x to exit: ")
		_, err := fmt.Scan(&s)
		i, err = strconv.Atoi(s)
		if err != nil {
			s = strings.ToLower(s)
			// fmt.Println(s)
			if s == "x" {
				fmt.Println("Goodbye :)")
				break
			} else {
				fmt.Println("Enter a valid input :(")
			}
		} else {
			slice = append(slice, i)
			sort.Ints(slice)
			fmt.Println(slice)
		}
	}
}
