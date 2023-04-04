package handler

import (
	"TechRead/database"
	"TechRead/model"
	"encoding/json"
	"fmt"
	"net/http"
	"database/sql"
)

func LoginHandler(w http.ResponseWriter, req *http.Request) {
	var reqUserData model.User
	var resState model.ResInfo
	if err := json.NewDecoder(req.Body).Decode(&reqUserData); err != nil {
		fmt.Println(err)
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		ResFail(resState)
		json.NewEncoder(w).Encode(resState)
		return
	}
	//入力されたEmailが存在しているかどうかを調べる
	row := database.DB.QueryRow("SELECT email FROM users WHERE email = ?", reqUserData.Email).Scan(&reqUserData.Email)
	if row == sql.ErrNoRows {
		fmt.Println("メールアドレスが違います")//入力されたメールアドレスが存在しない場合
		http.Error(w, "fail internal email \n", http.StatusInternalServerError)
		ResFail(resState)
		json.NewEncoder(w).Encode(resState)
		return
	}
	rows,err := database.DB.Query("SELECT user_id ,email, password FROM users WHERE email = ?", reqUserData.Email)//存在する場合
	if err != nil {
		http.Error(w, "fail internal email \n", http.StatusInternalServerError)
		ResFail(resState)
		json.NewEncoder(w).Encode(resState)
		return
	}
	fmt.Println(rows)

	var userRecord model.User
	for rows.Next() {
		err = rows.Scan(&userRecord.UserID,&userRecord.Email,&userRecord.Password)
		if err != nil {
			fmt.Println("データの読み込みに失敗しました",err)
			ResFail(resState)
			json.NewEncoder(w).Encode(resState)
			return
		}
		fmt.Println(userRecord.Email,userRecord.Password)
	}
	if reqUserData.Password != userRecord.Password {
		fmt.Println("パスワードが違います")
		ResFail(resState)
		json.NewEncoder(w).Encode(resState)
		return
	}
	resState.ResState = "success"
	resState.UserID = userRecord.UserID
	fmt.Println("ログインしました")
	cookie := http.Cookie{Name: "ID", Value: reqUserData.Email, Path: "/", HttpOnly: true}
	fmt.Println(cookie)
	http.SetCookie(w, &cookie)
	json.NewEncoder(w).Encode(resState)
}
