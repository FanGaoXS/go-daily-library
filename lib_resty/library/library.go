package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
)

// Libraries 结构体对应：https://api.cdnjs.com/libraries
type Libraries struct {
	Results   []*Library
	Total     int64
	Available int64
}

type Library struct {
	Name   string
	Latest string
}

func main() {
	c := resty.New()
	libraries := &Libraries{}
	_, err := c.R().SetResult(libraries).Get("https://api.cdnjs.com/libraries")
	if err != nil {
		log.Fatalln(err)
	}
	for _, library := range libraries.Results {
		fmt.Println("name: "+library.Name+",latest: ", library.Latest)
	}
	fmt.Println("total: ", libraries.Total)
	fmt.Println("available: ", libraries.Available)
}
