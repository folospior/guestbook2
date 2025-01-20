package handlers

import (
	"net/http"

	"github.com/a-h/templ"
)

func render(writer http.ResponseWriter, request *http.Request, component templ.Component) {
	component.Render(request.Context(), writer)
}
