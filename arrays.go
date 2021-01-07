package main

import "fmt"

func main() {
	var array_int [10]int
	array_str := [3]string{"si", "soy", "D="}

	fmt.Println(array_int)
	fmt.Println(array_str)
	fmt.Printf("The lenght of the first array is: %d\n ", len(array_int))

	for i := 0; i < len(array_int); i++ {
		array_int[i] = i
		fmt.Println(array_int[i])
	}

	var matrix [3][2]int
	fmt.Println(matrix)
	matrix[2][0] = 21
	fmt.Println(matrix)

	//Copy

	slice := []int{1, 2, 3, 4, 5, 6}
	slice_copy := make([]int, len(slice), cap(slice)*2)

	/* Is very important to declare the length of the array that you are copied, is better if you write len(array).
	About the capacity sometimes is a good practice to declare cap(array) * 2  */

	fmt.Println(slice)
	fmt.Println(slice_copy)

	copy(slice_copy, slice)

	fmt.Println(slice)
	fmt.Println(slice_copy)
}
