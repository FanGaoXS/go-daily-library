package main

import (
	"fmt"
	mid "lib_net-http/middleware"
	"log"
	"net/http"
	"time"
)

const (
	Port = ":9090"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/ping", mid.PanicMiddle(mid.TimeMiddle(mid.LoggerMiddle(http.HandlerFunc(ping)))))
	mux.Handle("/panic", mid.PanicMiddle(mid.TimeMiddle(mid.LoggerMiddle(http.HandlerFunc(panics)))))
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

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

func panics(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}
