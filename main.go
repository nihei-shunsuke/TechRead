package main

import (
	"TechRead/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", handler.LoginHandler).Methods(http.MethodPost)
	r.HandleFunc("/sign-up", handler.SignUpHandler).Methods(http.MethodPost)
	r.HandleFunc("/mypage", handler.FetchProfileHandler).Methods(http.MethodGet)
	r.HandleFunc("/update-profile", handler.EditProfileHandler).Methods(http.MethodPost)
	r.HandleFunc("/create-event", handler.CreateEventHandler).Methods(http.MethodPost)
	r.HandleFunc("/edit-event", handler.EditEventHandler).Methods(http.MethodPost)
	r.HandleFunc("/event-list", handler.EventListHandler).Methods(http.MethodGet)
	r.HandleFunc("/event-detail", handler.EventDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/create-abstract", handler.CreateAbstractHandler).Methods(http.MethodPost)
	r.HandleFunc("/edit-abstract", handler.EditAbstractHandler).Methods(http.MethodPost)

	http.ListenAndServe(":8080", r)
}
