package main

import (
	"fmt"
	"learning-go/7-tags/transform"
)

type User struct {
	ID    int    `transform:"-"`
	Name  string `transform:"upper"`
	Email string `transform:"lower"`
}

func main() {

	user := User{ID: 24, Name: "Pessoa", Email: "email@email.com"}
	err := transform.Tag(&user) // utilizando o método criado

	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
