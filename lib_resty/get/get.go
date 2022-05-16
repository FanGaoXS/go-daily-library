package main

import "github.com/go-resty/resty/v2"

func main() {
	c := resty.New()
	// query: https://www.baidu.com?name=test&age=15
	c.R().
		SetQueryString("name=test1&age=18").
		Get("https://www.baidu.com")

	c.R().
		SetQueryParams(map[string]string{
			"name": "test1",
			"age":  "18",
		}).
		Get("https://www.baidu.com")

	// path: https://www.baidu.com/test/15
	c.R().
		SetPathParams(map[string]string{
			"name": "test1",
			"age":  "18",
		}).
		Get("https://www.baidu.com/{name}/{age}")
}
