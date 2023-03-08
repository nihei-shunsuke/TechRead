package main

import (
	"TechRead/database"
	"TechRead/handler"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	c := cors.Default()
	database.Connect()
	sqlDB := database.DB
	defer sqlDB.Close()

	err := sqlDB.Ping()

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("データベース接続成功")
	}
	r.HandleFunc("/login", handler.LoginHandler).Methods(http.MethodPost)
	r.HandleFunc("/sign-up", handler.SignUpHandler).Methods(http.MethodPost)
	r.HandleFunc("/mypage", handler.FetchProfileHandler).Methods(http.MethodGet)
	r.HandleFunc("/update-profile", handler.EditProfileHandler).Methods(http.MethodPost)
	r.HandleFunc("/create-event", handler.CreateEventHandler).Methods(http.MethodPost)
	r.HandleFunc("/edit-event", handler.EditEventHandler).Methods(http.MethodPost)
	r.HandleFunc("/event-list", handler.EventListHandler).Methods(http.MethodGet)
	r.HandleFunc("/event-detail", handler.EventDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/create-chapter", handler.CreateChapterHandler).Methods(http.MethodPost)
	r.HandleFunc("/edit-chapter", handler.EditChapterHandler).Methods(http.MethodPost)
	handler := c.Handler(r)
	http.ListenAndServe(":8080", handler)
}
