package handler

import (
	"io"
	"net/http"
)

func FetchProfileHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "FetchProfileHandler\n")
}
