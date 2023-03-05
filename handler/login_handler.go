package handler

import (
	"TechRead/database"
	"TechRead/model"
	"encoding/json"
	"fmt"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, req *http.Request) {
	var reqUserData model.User
	if err := json.NewDecoder(req.Body).Decode(&reqUserData); err != nil {
		fmt.Println(err)
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}
	//入力されたEmailが存在しているかどうかを調べる
	row := database.DB.QueryRow("SELECT * FROM users WHERE email = ?", reqUserData.Email)
	if err := row.Err(); err != nil {
		fmt.Println("メールアドレスが違います")
		http.Error(w, "fail internal email \n", http.StatusInternalServerError)
		}
	rows, err := database.DB.Query("SELECT email, password FROM users WHERE email = ?" reqUserData.Email)
	if err != nil{
		http.Error(w, "fail internal read data \n", http.StatusInternalServerError)
	}
	_ := rows.Scan(&email)
	err := rows.Scan(&password)
	if err != nil{
		http.Error(w, "fail internal read data \n", http.StatusInternalServerError)
	}
	pass := password
	if reqUserData.Password != pass{
		fmt.Println("パスワードが違います")
		return
	}
	fmt.Println("ログインしました")
	json.NewEncoder(w).Encode(reqUserData)
}
