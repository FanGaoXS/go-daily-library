package main

import "log"

func main() {
	e := NewEmail()
	err := e.Send(EmailOption{
		From:    "wqk",
		To:      []string{"a954278478@gmail.com"},
		Cc:      []string{"954278478@qq.com"},
		Subject: "test",
		Text:    "test",
		Html:    "<h1>test</h1>",
	})
	if err != nil {
		log.Fatalln("send email err")
	}
}
