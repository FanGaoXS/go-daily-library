package main

import (
	"flag"
	"fmt"
)

var (
	name   string
	age    int
	gender bool
)

func main() {
	flag.Parse()

	fmt.Println("name :", name)
	fmt.Println("age :", age)
	fmt.Println("gender :", gender)
}

func init() {
	flag.StringVar(&name, "name", "", "name")
	flag.IntVar(&age, "age", 0, "age")
	flag.BoolVar(&gender, "gender", false, "gender: false is man")
}
