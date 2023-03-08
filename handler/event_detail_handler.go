package handler

import (
	"io"
	"net/http"
)

func EventDetailHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "EventDetailHandler\n")
}
