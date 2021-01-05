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

}
