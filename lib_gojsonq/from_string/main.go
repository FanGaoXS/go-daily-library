package main

import (
	"fmt"
	"github.com/thedevsaddam/gojsonq"
	"io/ioutil"
	"log"
)

const fileName = "./data.json"

func main() {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln("read file err")
	}

	// json from string
	gq := gojsonq.New().FromString(string(bytes))
	name := gq.Find("user.name")
	fmt.Println("name =", name)

	// gojsonq.Find() 会记录当前查询的节点，下次查询则会从当前记录的节点开始
	gq.Reset()

	hobbies := gq.Find("user.hobbies")
	fmt.Println("hobbies =", hobbies)

	gq.Reset()
	hobbies0 := gq.Find("user.hobbies.[0]")
	fmt.Println("hobbies[0] =", hobbies0)
}
