package handler

import (
	"io"
	"net/http"
)

func CreateChapterHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "CreateChapterHandler\n")
}
