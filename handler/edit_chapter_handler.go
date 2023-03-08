package handler

import (
	"io"
	"net/http"
)

func EditChapterHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "EditChapterHandler\n")
}
