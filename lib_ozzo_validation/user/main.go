package main

import "fmt"

func main() {
	u1 := &User{
		Name:  "test",
		Age:   100,
		Email: "954278478@qq.com",
	}
	fmt.Println("err =", u1.validate())

	u2 := &User{
		Name:  "test",
		Age:   150,
		Email: "954278478@qq.com",
	}
	fmt.Println("err =", u2.validate())
}
