package main

import "fmt"

func main() {
	var p1, p2 *int
	num := 5

	p1 = &num
	p2 = &num

	*p1 = 6

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(*p1)
	fmt.Println(*p2)
}
