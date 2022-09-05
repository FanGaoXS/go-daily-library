package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"log"
)

type Config struct {
	Name   string `short:"n" required_default:"true"`
	Gender string `short:"g" default:"man"`
}

func main() {
	var config Config

	if _, err := flags.Parse(&config); err != nil {
		log.Fatal("parse failed:", err)
	}

	fmt.Println(config.Name)
	fmt.Println(config.Gender)
}
