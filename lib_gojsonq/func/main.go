package main

import (
	"fmt"
	"github.com/thedevsaddam/gojsonq"
)

func main() {
	gq := gojsonq.New().File("./data.json").From("user.items")

	fmt.Println("the number of user.items = ", gq.Count())
	fmt.Println("sum of count of user.items = ", gq.Sum("count"))
	fmt.Println("max price of user.items = ", gq.Max("price"))
	fmt.Println("min price of user.items = ", gq.Min("price"))
	fmt.Println("avg price of user.items = ", gq.Avg("price"))
}
