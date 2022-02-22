package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"lib_sessions/users"
	"net/http"
)

func setUserHandler(w http.ResponseWriter, r *http.Request) {

	users.SetUser(r, &users.User{})
	err := sessions.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "Hello World")
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	user := users.GetUser(r)
	fmt.Fprintf(w, "user = %#v\n", user)
}
