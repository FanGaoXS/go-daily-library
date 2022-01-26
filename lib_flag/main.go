package main

import (
	"flag"
	"fmt"
)

var (
	name string
	sex  bool
	age  int
)

func init() {
	flag.StringVar(&name, "name", "guest", "姓名")
	flag.BoolVar(&sex, "sex", true, "性别（T：男性；F：女性）")
	flag.IntVar(&age, "age", 0, "年龄")
}

func main() {
	flag.Parse()
	fmt.Println("name =", name)
	fmt.Println("age =", age)
	fmt.Println("sex =", sex)
}
