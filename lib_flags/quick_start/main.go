package main

import (
	"fmt"

	"github.com/jessevdk/go-flags"
)

type Option struct {
	Names    []string       `short:"n" long:"name" description:"names of users"`
	Age      int            `short:"a" long:"age" description:"age of user"`
	Pointers []*string      `short:"p" long:"pointer" description:"pointers of *string"`
	Call     func(int)      `short:"c" long:"call" description:"call function"`
	Values   map[string]int `short:"v" long:"value" description:"values of user"`
}

func main() {
	var opt Option
	opt.Call = func(v int) {
		// 预先定义好call属性的函数类型的实现，然后parse后会自动进行调用
		fmt.Println("call function:", v)
	}

	if _, err := flags.Parse(&opt); err != nil {
		panic("flags parse failed")
	}

	fmt.Println(opt.Names)
	fmt.Println(opt.Age)
	for _, p := range opt.Pointers {
		fmt.Println(*p)
	}
	fmt.Println(opt.Values)
}
