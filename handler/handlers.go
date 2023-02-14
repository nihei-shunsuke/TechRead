package handler

import (
	"io"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "LoginHandler\n")
}

func SignUpHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "SignUpHandler\n")
}

func FetchProfileHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "FetchProfileHandler...test\n")
}

func EditProfileHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "EditProfileHandler\n")
}

func CreateEventHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "CreateEventHandler\n")
}

func EditEventHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "EditEventHandler\n")
}

func EventListHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "EventListHandler\n")
}

func EventDetailHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "EventDetailHandler\n")
}

func CreateAbstractHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "CreateAbstractHandler\n")
}

func EditAbstractHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "EditAbstractHandler\n")
}
