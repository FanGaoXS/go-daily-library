package main

import (
	"fmt"
	"go.uber.org/dig"
	"lib_dig/model"
)

// 构造函数不能返回结构体对象，应当返回函数
func initUser(name string, age int) func() *model.User {
	return func() *model.User {
		return &model.User{Name: name, Age: age}
	}
}

type UsersInParam struct {
	dig.In

	// 指定名称来使用
	U1 *model.User `name:"u1"`
	U2 *model.User `name:"u2"`
}

func PrintUsers(param UsersInParam) {
	fmt.Printf("user1 = %#v\n", param.U1)
	fmt.Printf("user2 = %#v\n", param.U2)
}

func main() {
	container := dig.New()

	// 将同一结构体的不同对象放入容器当中，需要命名，并且再取用的时候需要利用dig.in并且指定名称来使用
	container.Provide(initUser("t1", 18), dig.Name("u1"))
	container.Provide(initUser("t2", 20), dig.Name("u2"))

	container.Invoke(PrintUsers)
}
