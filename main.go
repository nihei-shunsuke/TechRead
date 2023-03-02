package main

import (
	"TechRead/database"
	"TechRead/handler"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	//Docker内でDBが初期化されるまで待つ
	time.Sleep(time.Second * 10)
	db := database.Connect()
	defer db.Close()

	err := db.Ping()

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
	r.HandleFunc("/create-abstract", handler.CreateAbstractHandler).Methods(http.MethodPost)
	r.HandleFunc("/edit-abstract", handler.EditAbstractHandler).Methods(http.MethodPost)

	http.ListenAndServe(":8080", r)
}
