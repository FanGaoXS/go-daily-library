package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

func Set(w http.ResponseWriter, r *http.Request) {
	SetSession(r, &User{})
	err := sessions.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "Hello World")
}

func Get(w http.ResponseWriter, r *http.Request) {
	user := GetSession(r)
	fmt.Fprintf(w, "user = %#v\n", user)
}
