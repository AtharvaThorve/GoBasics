package main

import "fmt"

func swap(inputSlice []int, index int) {
	temp := inputSlice[index]
	inputSlice[index] = inputSlice[index+1]
	inputSlice[index+1] = temp
}

func bubbleSort(inputSlice []int) {
	n := len(inputSlice)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if inputSlice[j] > inputSlice[j+1] {
				swap(inputSlice, j)
			}
		}
	}
}

func main() {
	slice := make([]int, 0, 10)
	fmt.Print("Enter number of values to be sorted: ")
	var n int
	fmt.Scan(&n)
	fmt.Println("Enter values to be sorted")
	var a int
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		slice = append(slice, a)
	}
	bubbleSort(slice)
	fmt.Println("Sorted slice:", slice)
}
