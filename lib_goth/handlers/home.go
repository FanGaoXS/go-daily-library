package handlers

import (
	"github.com/markbates/goth/gothic"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// 如果用户已经登陆，则会返回一个非空的user对象
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		http.Redirect(w, r, "/login/github", http.StatusTemporaryRedirect)
		return
	}
	ptTemplate.ExecuteTemplate(w, "home.tpl", user)
}
