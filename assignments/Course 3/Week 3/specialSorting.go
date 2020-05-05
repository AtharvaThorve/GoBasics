package main

import (
	"fmt"
	"sort"
)

func sort1(arr []int, arrayChannel chan []int) {
	sort.Ints(arr)
	fmt.Println("Sorted first quarter:", arr)
	arrayChannel <- arr
}

func sort2(arr []int, arrayChannel chan []int) {
	sort.Ints(arr)
	fmt.Println("Sorted second quarter:", arr)
	arrayChannel <- arr
}

func sort3(arr []int, arrayChannel chan []int) {
	sort.Ints(arr)
	fmt.Println("Sorted third quarter:", arr)
	arrayChannel <- arr
}

func sort4(arr []int, arrayChannel chan []int) {
	sort.Ints(arr)
	fmt.Println("Sorted fourth quarter:", arr)
	arrayChannel <- arr
}

func minInt(vals ...int) int {
	var index int
	min := -2147483648
	for i, val := range vals {
		if val > min {
			min = val
			index = i
		}
	}
	return index
}

func mergeTwoArrays(arr1, arr2 []int, arrayChannel chan []int) {
	n := len(arr1)
	m := len(arr2)
	arr := make([]int, n+m)
	var i, j, k int
	for i < n && j < m {
		if arr1[i] <= arr2[j] {
			arr[k] = arr1[i]
			i++
		} else {
			arr[k] = arr2[j]
			j++
		}
		k++
	}
	for i < n {
		arr[k] = arr1[i]
		i++
		k++
	}
	for j < m {
		arr[k] = arr2[j]
		j++
		k++
	}
	arrayChannel <- arr
}

func main() {
	var n int
	arrayChannel := make(chan []int)
	fmt.Print("Enter size of the array to be sorted: ")
	fmt.Scan(&n)
	arr := make([]int, n)
	fmt.Println("Enter the values of the array to be sorted:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	go sort1(arr[0:n/4], arrayChannel)
	go sort2(arr[n/4:n/2], arrayChannel)
	go sort3(arr[n/2:3*n/4], arrayChannel)
	go sort4(arr[3*n/4:n], arrayChannel)

	sorted1 := <-arrayChannel
	sorted2 := <-arrayChannel
	sorted3 := <-arrayChannel
	sorted4 := <-arrayChannel

	go mergeTwoArrays(sorted1, sorted2, arrayChannel)
	go mergeTwoArrays(sorted3, sorted4, arrayChannel)

	firstHalf := <-arrayChannel
	secondHalf := <-arrayChannel

	go mergeTwoArrays(firstHalf, secondHalf, arrayChannel)

	arr = <-arrayChannel

	fmt.Println("Final sorted array is:", arr)
}
