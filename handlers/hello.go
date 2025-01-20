package handlers

import "net/http"

func HandleHello(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello, World!"))
}