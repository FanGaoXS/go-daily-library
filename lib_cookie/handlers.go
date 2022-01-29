package main

import (
	"fmt"
	"log"
	"net/http"
)

func SetCookieHandler(w http.ResponseWriter, r *http.Request) {
	u := &User{
		Name: "wqk",
		Age:  18,
		Role: "admin",
	}
	cookie, err := EncodeCookie(u)
	if err != nil {
		log.Fatal("encode cookies error: ", err)
	}
	// write into response
	http.SetCookie(w, cookie)
}

func GetCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(CookieKey)
	if err != nil {
		log.Fatal("get cookie from request error: ", err)
	}
	user, err := ReadCookie(cookie)
	if err != nil {
		log.Fatal("decode cookie error: ", err)
	}
	fmt.Fprintf(w, "user = %#v\n", user)
}
