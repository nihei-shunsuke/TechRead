package handler

import (
	"io"
	"net/http"
)

func ChapterListHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "ChapterListHandler\n")
}
