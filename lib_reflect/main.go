package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	name string
}

func main() {
	var i int64 = 5
	var a Animal
	a.name = "cat"

	// i的类型
	t1 := reflect.TypeOf(i)
	// a的类型
	t2 := reflect.TypeOf(a)

	fmt.Println(t1.String())
	fmt.Println(t2.String())
}
