package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
)

type User struct {
	Name    string
	Contact *Contact
}

type Contact struct {
	Email string
	Phone string
}

func main() {

	// cmp.Equal() 比较内容，如果比较的对象有未导出字段，则会发生panic
	// 可以使用cmdopts.IngoreUnexported来忽略未导出字段，从而不发生panic
	// == 比较内容和指针指向

	u1 := User{Name: "test"}
	u2 := User{Name: "test"}
	fmt.Println("u1 == u2?", u1 == u2)
	fmt.Println("u1 equals u2?", cmp.Equal(u1, u2))

	c := &Contact{Email: "xxx@qq.com", Phone: "10086"}
	u1.Contact = c
	u2.Contact = c
	fmt.Println("u1 == u2?", u1 == u2)
	fmt.Println("u1 equals u2?", cmp.Equal(u1, u2))

	c1 := &Contact{Email: "xxx@qq.com", Phone: "10086"}
	c2 := &Contact{Email: "xxx@qq.com", Phone: "10086"}
	u1.Contact = c1
	u2.Contact = c2
	fmt.Println("u1 == u2?", u1 == u2)
	fmt.Println("u1 equals u2?", cmp.Equal(u1, u2))
}
