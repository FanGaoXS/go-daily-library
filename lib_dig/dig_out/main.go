package main

import (
	"fmt"
	"go.uber.org/dig"
	"lib_dig/env"
	"lib_dig/model"
	"log"
)

type Result struct {
	dig.Out

	User *model.User
	Shoe *model.Shoe
}

// 使用dig.out的时候，返回值直接返回包含dig.out的结构体实例
func initUserAndShoe(opt *env.Option) Result {
	return Result{
		User: &model.User{
			Name: opt.User.Name,
			Age:  opt.User.Age,
		},
		Shoe: &model.Shoe{
			Brand: opt.Shoe.Brand,
		},
	}
}

// 使用dig.out：直接使用具体的对象实例，而不是Result实例的属性：
// 如直接使用User实例，而不是Result.User实例
func printInfo(user *model.User, shoe *model.Shoe) {
	fmt.Printf("user = %#v\n", user)
	fmt.Printf("shoe = %#v\n", shoe)
}

// 不使用dig.out
//func printInfo(r *Result) {
//	fmt.Printf("user = %#v\n", r.User)
//	fmt.Printf("shoe = %#v\n", r.Shoe)
//}

func main() {
	container := dig.New()

	container.Provide(env.InitEnv)
	container.Provide(initUserAndShoe)

	err := container.Invoke(printInfo)
	if err != nil {
		log.Fatal(err)
	}
}
