package main

import (
	"fmt"
	"strconv" // package that permit convert variables
)

func main() {
	var x int
	x = 69
	g := 21 //Operator type dynamic
	name := "My name is Gino"
	fmt.Println(x + g)
	fmt.Println(name)

	//casting in go

	age := 20
	age_str := strconv.Itoa(age)

	n := "19"
	n_int, _ := strconv.Atoi(n)

	fmt.Println("I'm " + age_str + " years old")
	fmt.Println(n_int + 2)
}
