package main

import (
	"net/http"

	"github.com/folospior/guestbook2/handlers"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", handlers.HandleIndex)
	router.HandleFunc("/hello", handlers.HandleHello)
	router.HandleFunc("POST /writeIn", handlers.HandleWriteIn)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}
