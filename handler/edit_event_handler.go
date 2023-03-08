package handler

import (
	"io"
	"net/http"
)

func EditEventHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "EditEventHandler\n")
}
