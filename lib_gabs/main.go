package main

import (
	"fmt"
	"github.com/Jeffail/gabs/v2"
)

func main() {
	jsonStringBytes := []byte(`{
"info": {
  "name": {
	"quick_start": "lee",
	"last": "darjun"
  },
  "age": 18,
  "hobbies": [
	"game",
	"programming"
  ]
}
}`)
	jsonObj, _ := gabs.ParseJSON(jsonStringBytes)
	// 查询的三种方法：
	// 1、Path()：以.分割
	first := jsonObj.Path("info.name.quick_start").Data().(string)
	fmt.Println("quick_start =", first)
	// 2、Search()：传多个参数，可以简写为S()
	last := jsonObj.Search("info", "name", "last").Data().(string)
	fmt.Println("last =", last)
	// 3、JSONPointer()：以/分割，并且必须以/开头
	pointer, _ := jsonObj.JSONPointer("/info/age")
	age := pointer.Data().(float64)
	fmt.Println("age =", age)

	// 获取数组
	hobbies := jsonObj.Path("info.hobbies").Data().([]interface{})
	fmt.Println("hobbies =", hobbies)

	// 获取数组中具体元素
	game := jsonObj.Search("info", "hobbies", "0").Data().(string)
	fmt.Println("game =", game)
	game = jsonObj.Path("info.hobbies.0").Data().(string)
	fmt.Println("game =", game)
	pointer, _ = jsonObj.JSONPointer("/info/hobbies/0")
	game = pointer.Data().(string)
	fmt.Println("game =", game)

	// 如果JSON字符串是对象数组
	jsonObj, _ = gabs.ParseJSON([]byte(`{
    "user": {
      "name": "dj",
      "age": 18,
      "members": [
        {
          "name": "hjw",
          "age": 20,
          "relation": "spouse"
        },
        {
          "name": "lizi",
          "age": 3,
          "relation": "son"
        }
      ]
    }
  }`))
	names := jsonObj.Path("user.members.*.name").Data().([]interface{})
	fmt.Println("names =", names)

	ages := jsonObj.Search("user", "members", "*", "age").Data().([]interface{})
	fmt.Println("ages =", ages)

	pointer, _ = jsonObj.JSONPointer("/user/members/*/relation")
	relations := pointer.Data().([]interface{})
	fmt.Println("relations =", relations)

}
