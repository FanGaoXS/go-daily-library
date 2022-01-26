package main

import (
	"github.com/gorilla/mux"
	"lib_goth/handlers"
	"lib_goth/providers"
	"log"
	"net/http"
)

const (
	Port = ":9091"
)

func main() {
	initialize()
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/login/github", handlers.LoginHandler)
	r.HandleFunc("/logout/github", handlers.LogoutHandler)
	r.HandleFunc("/auth/github", handlers.AuthHandler)
	r.HandleFunc("/auth/github/callback", handlers.CallbackHandler)

	log.Println("listening on localhost" + Port)
	log.Fatal(http.ListenAndServe(Port, r))
}

func initialize() {
	handlers.Init()
	providers.Init()
}
