package main

import (
	"fmt"
	"go.uber.org/dig"
	"lib_dig/env"
	"lib_dig/model"
)

func initUser(opt *env.Option) *model.User {
	return model.NewUser(opt.User.Name, opt.User.Age)
}

func initShoe(opt *env.Option) *model.Shoe {
	return model.NewShoe(opt.Shoe.Brand)
}

// Args 在invoke的时候，方便直接使用Args结构体对象的属性来进行调用
// 如arg.User来使用容器中的user实例
type Args struct {
	dig.In

	User *model.User
	Shoe *model.Shoe
}

// 使用dig.In
func print(args Args) {
	fmt.Printf("user = %#v\n", args.User)
	fmt.Printf("shoe = %#v\n", args.Shoe)
}

//// 不使用dig.In
//func print(user *model.User, shoe *model.Shoe) {
//	fmt.Printf("user = %#v\n", user)
//	fmt.Printf("shoe = %#v\n", shoe)
//}

func main() {
	container := dig.New()

	container.Provide(env.InitEnv)
	container.Provide(initUser)
	container.Provide(initShoe)

	container.Invoke(print)
}
