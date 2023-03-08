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
	if err := json.NewDecoder(req.Body).Decode(&reqUserData); err != nil {
		fmt.Println(err)
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}
	//入力されたEmailが存在しているかどうかを調べる
	u := reqUserData
	row := database.DB.QueryRow("SELECT email FROM users WHERE email = ?", u.Email).Scan(&u.Email)
	if row == sql.ErrNoRows {
		fmt.Println("メールアドレスが違います")//入力されたメールアドレスが存在しない場合
		http.Error(w, "fail internal email \n", http.StatusInternalServerError)
		return
	}
	rows, err := database.DB.Query("SELECT email, password FROM users WHERE email = ?", u.Email)//存在する場合
	if err != nil {
		http.Error(w, "fail internal email \n", http.StatusInternalServerError)
		return
	}
	for rows.Next() {
		err := rows.Scan(&u.Email,&u.Password)
		if err != nil {
			fmt.Println("データの読み込みに失敗しました",err)
			return
		}
		if reqUserData.Password != u.Password {
			fmt.Println("パスワードが違います")
			return
		}
	fmt.Println("ログインしました")
	json.NewEncoder(w).Encode(reqUserData)
	}
}
