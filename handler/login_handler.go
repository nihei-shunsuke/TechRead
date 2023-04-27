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
	var userRecord model.User
	var resState model.ResInfo
	//リクエストをreqUserDataに書き込む
	if err := json.NewDecoder(req.Body).Decode(&reqUserData); err != nil {
		ResFail(resState)
		json.NewEncoder(w).Encode(resState)
		return
	}
	//入力されたEmailが存在しているかどうかを調べる
	row := database.DB.QueryRow("SELECT email FROM users WHERE email = ?", reqUserData.Email).Scan(&reqUserData.Email)
	if row == sql.ErrNoRows {//もしメールアドレスが存在しない場合
		ResFail(resState)
		json.NewEncoder(w).Encode(resState)
		return
	}
	//ユーザーが入力したメールアドレスからそれらの情報をrowsに入れる
	rows,err := database.DB.Query("SELECT user_id ,email, password FROM users WHERE email = ?", reqUserData.Email)//存在する場合
	if err != nil {//もし上の処理でエラーが出た場合
		ResFail(resState)
		json.NewEncoder(w).Encode(resState)
		return
	}

	for rows.Next() {//上の情報をここでuserRecordに書き込む
		err = rows.Scan(&userRecord.UserID,&userRecord.Email,&userRecord.Password)
		if err != nil {
			ResFail(resState)
			json.NewEncoder(w).Encode(resState)
			return
		}
		fmt.Println(userRecord.Email,userRecord.Password)
	}
	//ここでパスワードの比較
	if reqUserData.Password != userRecord.Password {
		ResFail(resState)
		json.NewEncoder(w).Encode(resState)
		return
	}
	//レスポンスに成功とユーザーのIDを送信する
	resState.ResState = "success"
	resState.UserID = userRecord.UserID
	fmt.Println("ログインしました")
	json.NewEncoder(w).Encode(resState)
}
