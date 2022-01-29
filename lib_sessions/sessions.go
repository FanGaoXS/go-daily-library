package main

import (
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

var (
	store = sessions.NewFilesystemStore("./", securecookie.GenerateRandomKey(32), securecookie.GenerateRandomKey(32))
)

const (
	SessionKey = "session_key"
)

type User struct {
	Name string
	Age  int
}

func SetSession(r *http.Request, user *User) {
	session, err := store.Get(r, SessionKey)
	if err != nil {
		log.Fatal("set session error: ", err)
	}
	session.Values["user"] = user
}

func GetSession(r *http.Request) *User {
	session, err := store.Get(r, SessionKey)
	if err != nil {
		log.Fatal("get session error: ", err)
	}
	user, ok := session.Values["user"].(*User)
	if !ok {
		log.Fatal("get user from session error: ", err)
		return nil
	}
	return user
}
