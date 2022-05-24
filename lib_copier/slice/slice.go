package main

import (
	"fmt"
	"github.com/jinzhu/copier"
)

type User struct {
	Name     string
	Age      int
	Username string
	Password string
}

type People struct {
	Name string
	Age  int
}

func main() {
	users := []*User{
		{
			Name:     "test0",
			Age:      18,
			Username: "username0",
			Password: "password0",
		},
		{
			Name:     "test1",
			Age:      18,
			Username: "username1",
			Password: "password1",
		},
		{
			Name:     "test0",
			Age:      18,
			Username: "username2",
			Password: "password2",
		},
	}

	var people []*People
	copier.Copy(&people, users)
	for _, person := range people {
		fmt.Printf("p = %#v\n", person)
	}
}
