package handlers

import (
	"net/http"

	"github.com/folospior/guestbook2/views/components"
)

func HandleIndex(writer http.ResponseWriter, request *http.Request) {
	render(writer, request, components.Index())
}
