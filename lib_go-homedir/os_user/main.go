package main

import (
	"fmt"
	"log"
	"os/user"
)

func main() {
	uesr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("user =", uesr)
	fmt.Println("user.HomeDir =", uesr.HomeDir)
}
