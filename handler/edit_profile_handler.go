package handler

import (
	"io"
	"net/http"
)

func EditProfileHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "EditProfileHandler\n")
}
