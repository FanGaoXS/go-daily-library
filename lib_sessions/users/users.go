package users

import (
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"reflect"
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

func fields() []string {
	var fields []string
	types := reflect.TypeOf(&User{})
	for i := 0; i < types.NumField(); i++ {
		field := types.Field(i)
		fields = append(fields, field.Name)
	}
	return fields
}

func SetUser(r *http.Request, user *User) {
	session, err := store.Get(r, SessionKey)
	if err != nil {
		log.Fatal("set session error: ", err)
	}
	session.Values["user"] = user
}

func GetUser(r *http.Request) *User {
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
