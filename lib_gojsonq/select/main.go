package main

import (
	"fmt"
	"github.com/thedevsaddam/gojsonq"
	"reflect"
)

func main() {
	gq := gojsonq.New().File("./data.json")

	// 查询user.items中的id和name属性，像极了sql
	// select id,name from user.items
	maps := gq.From("user.items").Select("id", "name").Get()
	fmt.Println("type of maps =", reflect.TypeOf(maps))
	if value, ok := maps.([]interface{}); ok {
		for _, item := range value {
			fmt.Println(item)
		}
	}
}
