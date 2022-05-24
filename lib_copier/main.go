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
	u := &User{
		Name:     "test",
		Age:      18,
		Username: "username",
		Password: "password",
	}

	// 将user实例复制到people实例
	var p People
	err := copier.Copy(&p, u)
	if err != nil {
		panic(err)
	}
	fmt.Printf("p = %#v\n", p)
}
