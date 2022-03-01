package main

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func main() {
	u1 := &User{
		Name:  "test",
		Age:   100,
		Email: "954278478@qq.com",
	}
	fmt.Println("err =", u1.Validate())
	// or
	fmt.Println("err =", validation.Validate(u1, validation.NotNil))

	u2 := &User{
		Name:  "test",
		Age:   150,
		Email: "954278478@qq.com",
	}
	fmt.Println("err =", u2.Validate())
	// or
	fmt.Println("err =", validation.Validate(u2, validation.NotNil))
}
