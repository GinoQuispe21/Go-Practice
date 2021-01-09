package main

import "fmt"

type User struct {
	age             int
	name, last_name string
}

//Method of the struct
func (user User) full_name() string {
	return user.name + " " + user.last_name
}

//Use pointer to modify the object
func (this *User) set_name(n string) {
	this.name = n
}

func main() {
	var Gino User

	fmt.Println(Gino)

	Gino.age = 19
	Gino.name = "Gino"
	Gino.last_name = "Quispe Calixto"

	fmt.Println(Gino)

	Roro := User{age: 20, name: "Rodrigo", last_name: "Ramirez Bracamonte"}

	fmt.Println(Roro)

	Yuanca := User{18, "Juan Carlos", "Hernandez  Castillo"}

	fmt.Println(Yuanca)

	leyva := new(User) //Declara una estructura que retorna un puntero

	fmt.Println(leyva)
	fmt.Println(leyva.age, leyva.name, leyva.last_name)

	(*leyva).age = 18

	leyva.name = "Juan Alonso"
	leyva.last_name = "Leyva Calle"

	fmt.Println(leyva.age, leyva.name, leyva.last_name)
	fmt.Println(leyva.full_name())
	leyva.set_name("Pepito")
	fmt.Println(leyva.full_name())
}
