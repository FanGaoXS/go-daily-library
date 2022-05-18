package main

import (
	"log"
	"net/smtp"

	"github.com/joho/godotenv"
	"github.com/jordan-wright/email"
)

type Email struct {
	host     string
	port     string
	username string
	password string
}

func NewEmail() Email {
	envMap, err := godotenv.Read(".env")
	if err != nil {
		log.Fatalln("read env err")
	}

	return Email{
		host:     envMap["smtp_host"],
		port:     envMap["smtp_port"],
		username: envMap["email_username"],
		password: envMap["email_password"],
	}
}

type EmailOption struct {
	From    string
	To      []string
	Cc      []string
	Subject string
	Text    string
	Html    string
}

func (e Email) Send(option EmailOption) error {
	em := email.NewEmail()
	em.From = option.From + "<" + e.username + ">"
	em.To = option.To
	em.Cc = option.Cc
	em.Subject = option.Subject
	em.Text = []byte(option.Text)
	em.HTML = []byte(option.Html)

	// 添加附件
	//em.AttachFile(".env")

	addr := e.host + ":" + e.port
	return em.Send(addr, smtp.PlainAuth("", e.username, e.password, e.host))
}
