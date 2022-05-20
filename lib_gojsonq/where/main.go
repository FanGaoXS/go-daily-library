package main

import (
	"fmt"
	"github.com/thedevsaddam/gojsonq"
)

func main() {
	gq := gojsonq.New().File("./data.json")

	r := gq.From("user.items").
		Select("id", "name").
		Where("id", "=", 1).
		Get()
	fmt.Println("select id, name from user.items where id = 1: ", r)

	gq.Reset()

	r = gq.From("user.items").
		Select("id", "name", "count", "price").
		Where("id", ">", 0).
		Where("id", "<=", 3).
		Get()
	fmt.Println("select id, name, count, price from user.items where id > 0 and id <= 3: ", r)

	gq.Reset()

	r = gq.From("user.items").
		Select("id", "name", "count", "price").
		Where("id", "=", 3).
		OrWhere("id", "=", 1).
		Get()
	fmt.Println("select id, name, count, price from user.items where id = 3 or id = 1: ", r)

	// 还有whereIn whereStartsWith ...
}
