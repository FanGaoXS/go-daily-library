package main

import (
	"github.com/gorilla/securecookie"
	"net/http"
)

const (
	CookieKey = "cookie_user"
)

type User struct {
	Name string
	Age  int
	Role string
}

var (
	hashKey  = securecookie.GenerateRandomKey(16)
	blockKey = securecookie.GenerateRandomKey(16)
	s        = securecookie.New(hashKey, blockKey)
)

func EncodeCookie(user *User) (*http.Cookie, error) {
	encode, err := s.Encode(CookieKey, user)
	if err != nil {
		return nil, err
	}
	cookie := &http.Cookie{
		Name:     CookieKey,
		Value:    encode,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
	}
	return cookie, nil
}

func ReadCookie(cookie *http.Cookie) (User, error) {
	var user User
	err := s.Decode(CookieKey, cookie.Value, &user)
	return user, err
}
