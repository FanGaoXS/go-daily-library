package main

import (
	"log"
	"net/http"
	"time"
)

const (
	Port = ":9090"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/set", Set)
	mux.HandleFunc("/get", Get)
	server := &http.Server{
		Addr:         Port,
		Handler:      mux,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
