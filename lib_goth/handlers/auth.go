package handlers

import (
	"fmt"
	"github.com/markbates/goth/gothic"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	ptTemplate.ExecuteTemplate(w, "login.tpl", nil)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	gothic.Logout(w, r)
	ptTemplate.ExecuteTemplate(w, "logout.tpl", nil)
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	gothic.BeginAuthHandler(w, r)
}

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	ptTemplate.ExecuteTemplate(w, "home.tpl", user)
}
