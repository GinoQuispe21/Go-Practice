//This concept is very similar to inheritance in POO

package main

import "fmt"

type Human struct {
	name string
}

func (this Human) talk() string {
	return "bla bla bla"
}

type Dummy struct {
	name string
}

type Student struct {
	Human
	Dummy
}

func (this Student) talk() string {
	return "I'm student"
}

func main() {
	Gino := Student{Human{"Gino"}, Dummy{"Gino"}}

	// Gino is talking as a student
	fmt.Println(Gino.talk())

	// Gino is talking as a human
	fmt.Println(Gino.Human.talk())

	// Ambigous selector
	//fmt.Println(Gino.name)
}
