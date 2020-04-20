package main

import "fmt"

func main() {
	// Slice is like a window on an underlying array. It is of variable size
	/* Properties of a slice
	*  Contains a pointer to the start of the slice
	*  Length is the number of elements in the slice
	*  Capacity is the max number of elements in the slice
		* It is the difference between the start pointer and size of the whole array
	*/

	arr := [...]string{"a", "b", "c", "d", "e", "f", "g"}
	slice1 := arr[1:3] // This is the how we define slices.
	// slice1 includes array element 1 and 2 but not 3
	fmt.Println(slice1)
	slice2 := arr[2:5]
	fmt.Println(slice2)

	primes := [...]int{2, 3, 5, 7}
	primeSlice := primes[0:2]
	// len returns length and cap returns capacity
	fmt.Println(primeSlice, len(primeSlice), cap(primeSlice))

	// Writing to a slice changes the underlying array so anyother slice that refers to that same
	// array would also get changed
	fmt.Println(slice1[1], slice2[0], slice1[1] == slice2[0])
	slice1[1] = "a"
	fmt.Println("Changed slice[1] to \"a\" so value of slice2[0] also changes to \"a\"")
	fmt.Println(slice1[1], slice2[0], slice1[1] == slice2[0])

	// When you use slice literal to define a slice then the slice points to the start of the array
	// and the length is equal to capacity

	sliceLiteral := []int{1, 2, 3} /* In array literal here the size of the array needs to be specified
	or ... operator is need to be used*/
	fmt.Println(len(sliceLiteral) == cap(sliceLiteral))

	fmt.Print("Enter a number: ")
	var n int
	fmt.Scan(&n)
	makeSlice1 := make([]int, n) // If you want to make a slice of variable size then this syntax is used
	/* This can take two or three arguments.
	*  The first arg is which type of slice this is.
	*  The second is the length of the slice.
	*  Third is the capacity of the slice (optional)
	*  If capacity is not defined then length and capacity are the same
	 */
	makeSlice2 := make([]int, n, n+10)
	fmt.Printf("Length and capacity of two arg make function slice: %d, %d\n", len(makeSlice1), cap(makeSlice1))
	fmt.Printf("Length and capacity of three arg make function slice: %d, %d\n", len(makeSlice2), cap(makeSlice2))

	// append function can be used to increase the size of a slice.
	// if the size of the slice exceeds the size of the underlying array then size of the underlying
	// array is increased so append function never has to stop appending.

	sli := make([]int, 0, 3) // Length of the slice is 0 and capacity is 3
	sli = append(sli, 0)     // append the value 100 to the end of the slice so in this case the size of
	// slice increases but the capacity of the underlying array remains the same
	fmt.Println("After appending an element to a slice of length 0:", len(sli), cap(sli))
	sli = append(sli, 1, 2, 3, 4, 5)
	fmt.Println(sli, len(sli), cap(sli)) // Both length and capacity of slice increased
	// Appending a slice to another slice or just itself syntax is same for both
	sli = append(sli, sli...) // This appends slice sli to itself
	// The ... operator is an unpacking operator so the elements of slice are appended and not the
	// slice itself
	fmt.Println(sli, len(sli), cap(sli))
}
