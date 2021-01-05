package main

import (
	"bufio"
	"fmt" //input and output
	"os"
)

func main() {
	fmt.Print("Hi my friend")
	fmt.Println("Hi my friend") //This add you a endl

	followers := 319
	fmt.Printf("My numbers or followers is %d\n", followers)

	name := "Gino"
	fmt.Printf("My name is %v\n", name)

	soy := true
	fmt.Printf("I'm telling the %t\n", soy)

	price := 100.99
	fmt.Printf("The price is %f\n", price)

	var current_year int
	fmt.Printf("Write the current year: ")
	fmt.Scanf("%d\n", &current_year)
	fmt.Printf("The current year is %d\n", current_year)

	var lastname string
	fmt.Printf("Write your lastname: ")
	fmt.Scanf("%s\n", &lastname)
	fmt.Printf("My lastname is %s\n", lastname)

	//Other way, using reader of bufio
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Write the number of your credit card :")
	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The number is " + text)
	}
}
