package main

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"log"
)

func main() {
	dir, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Home dir:", dir)
}
