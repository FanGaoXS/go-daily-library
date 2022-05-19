package main

import (
	"fmt"
	"go.uber.org/dig"
	"lib_dig/model"
)

func initUser(name string, age int) func() *model.User {
	return func() *model.User {
		return &model.User{Name: name, Age: age}
	}
}

type Args struct {
	dig.In

	Users []*model.User `group:"user"`
}

func printInfo(args Args) {
	for i, u := range args.Users {
		fmt.Printf("user[%d] = %#v\n", i, u)
	}
}

func main() {
	container := dig.New()

	container.Provide(initUser("u1", 18), dig.Group("user"))
	container.Provide(initUser("u2", 18), dig.Group("user"))

	container.Invoke(printInfo)
}
