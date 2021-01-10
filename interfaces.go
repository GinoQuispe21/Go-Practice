package main

import "fmt"

type User interface {
	Permission() int
	Name() string
}

type Admin struct {
	name string
}

type Editor struct {
	name string
}

func (this Admin) Permission() int {
	return 5
}

func (this Admin) Name() string {
	return this.name
}

func (this Editor) Permission() int {
	return 4
}

func (this Editor) Name() string {
	return this.name
}

func Auth(user User) string {
	if user.Permission() > 4 {
		return user.Name() + " is allowed"
	} else {
		return user.Name() + " is not allowed"
	}
}

func main() {
	admin := Admin{"Gino"}
	editor := Editor{"Pedro"}

	fmt.Println(Auth(admin))
	fmt.Println(Auth(editor))

	users := []User{Admin{"Juanca"}, Editor{"Roro"}}
	for _, user := range users {
		fmt.Println(Auth(user))
	}
}
