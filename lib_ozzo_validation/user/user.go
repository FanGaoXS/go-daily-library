package main

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type User struct {
	Name  string
	Age   int
	Email string
}

func (u *User) validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Name, validation.Required),
		validation.Field(&u.Age, validation.Required, validation.Min(1), validation.Max(100)),
		validation.Field(&u.Email, validation.Required, is.Email),
	)
}
