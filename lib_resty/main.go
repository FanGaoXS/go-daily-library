package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
)

func main() {
	// 1、create http client
	client := resty.New()
	// 2、创建请求对象（请求方法get）
	resp, err := client.R().Get("https://www.baidu.com/")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Response Info:")
	fmt.Println("Status Code:", resp.StatusCode())
	fmt.Println("Status:", resp.Status())
	fmt.Println("Proto:", resp.Proto())
	fmt.Println("Time:", resp.Time())
	fmt.Println("Received At:", resp.ReceivedAt())
	fmt.Println("Size:", resp.Size())
	fmt.Println("Headers:")
	for key, value := range resp.Header() {
		fmt.Println(key, "=", value)
	}
	fmt.Println("Cookies:")
	for i, cookie := range resp.Cookies() {
		fmt.Printf("cookie%d: name:%s value:%s\n", i, cookie.Name, cookie.Value)
	}
}
