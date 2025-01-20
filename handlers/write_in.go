package handlers

import (
	"net/http"

	"github.com/folospior/guestbook2/views/components"
)

func HandleWriteIn(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		http.Error(writer, "failed to parse form", http.StatusInternalServerError)
		return
	}

	content := request.FormValue("content")
	if content == "" {
		http.Error(writer, "text is required", http.StatusExpectationFailed)
		return
	}

	render(writer, request, components.WriteIn(content))
}

func HandleGetWriteIns(writer http.ResponseWriter, request *http.Request) {
	render(writer, request, components.WriteIns([]string{"hello", "world", "fuck"}))
}
