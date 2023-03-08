package handler

import (
	"io"
	"net/http"
)

func EventListHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "EventListHandler\n")
}
