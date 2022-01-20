package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
)

type Option struct {
	Name string `short:"n" long:"name" description:"姓名"`
	Age  int    `short:"a" long:"age" description:"年龄"`
	Sex  bool   `short:"s" long:"sex" description:"性别"`
}

func main() {
	option := Option{}
	flags.Parse(&option)
	fmt.Printf("option = %#v\n", option)
}
