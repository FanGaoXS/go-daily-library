package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"log"
)

type Option struct {
	Name    string        `long:"name"`
	Info    InfoOption    `group:"info"`
	Contact ContactOption `group:"contact"`
}

type InfoOption struct {
	Gender bool `long:"gender"`
	Age    int  `long:"age"`
}

type ContactOption struct {
	Email string `long:"email"`
	Phone string `long:"phone"`
}

func main() {
	var opt Option
	if _, err := flags.Parse(&opt); err != nil {
		log.Fatalln("flags parse failed:", err)
	}

	fmt.Println(opt.Name)
	fmt.Println(opt.Info)
	fmt.Println(opt.Contact)
}
