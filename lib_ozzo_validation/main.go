package main

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"log"
)

func main() {
	url := "https://www.baidu.com"
	err := validation.Validate(url,
		validation.Required,
		is.URL)
	if err != nil {
		log.Fatalln(err)
	}
}
