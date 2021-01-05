package main

import "fmt"

func main() {

	var x, y int

	fmt.Print("Write the first number: ")
	fmt.Scanf("%d\n", &x)

	fmt.Print("Write the second number: ")
	fmt.Scanf("%d\n", &y)

	if x < y {
		fmt.Printf("number %d is less than number %d\n", x, y)
	} else if x > y {
		fmt.Printf("number %d is less than number %d\n", y, x)
	} else {
		fmt.Printf("You wrote the same numbers: %d = %d", x, y)
	}
}

// ==, !=, &&, ||
//NOTICE
//In conditionals, you must write the curly braces in the same line that you write the condition.
//Also, never forget to type the curly braces, because without this the program cannot run.
