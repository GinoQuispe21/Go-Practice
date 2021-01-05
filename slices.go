package main

import "fmt"

func main() {

	array := [3]int{1, 2, 3}
	slice1 := array[1:2]
	//slice2 := array[1:3]
	//slice3 := array[2:3]
	fmt.Println(array)

	fmt.Println(slice1)
	//fmt.Println(slice2)
	//fmt.Println(slice3)

	// Pointer to array
	// Length of the array  that is pointing
	// Capacity

	slice := make([]int, 0, 5) // pointer to an array, lenght : 3, capacity : 5
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = append(slice, 21)

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// It is better if you write the capacity of your slice.
	// If you don't write the capacity, the slice will be more inefficient
	// because when you want to use "append" on your slice, it will use more capacity than it probably needs
}

//nil = null
