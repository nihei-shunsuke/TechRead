package handler

import (
	"TechRead/database"
	"TechRead/model"
	"encoding/json"
	"fmt"
	"net/http"
	"database/sql"
)

func SignUpHandler(w http.ResponseWriter, req *http.Request) {
	var reqUserData model.User
	var resState model.ResInfo
	//httpリクエスト(req.Body)に入ってるjsonデータ(model.User)を基にUser構造体にデコードする
	if err := json.NewDecoder(req.Body).Decode(&reqUserData); err != nil {
		ResFail(resState)
		json.NewEncoder(w).Encode(resState)
		return
	}
	fmt.Println(reqUserData)
	//User構造体にリクエストが記入したデータを入れる処理を書く
	const sqlStr = `
		insert into users (user_name, password, email) values
		(?,?,?);
	`
	//リクエストのメールアドレスに被りがないか調べる
	row := database.DB.QueryRow("SELECT email FROM users WHERE email = ?", reqUserData.Email).Scan(&reqUserData.Email)
	if row == sql.ErrNoRows {//被りがない場合
		result, err := database.DB.Exec(sqlStr, reqUserData.UserName, reqUserData.Password, reqUserData.Email)
		if err != nil {
			ResFail(resState)
			json.NewEncoder(w).Encode(resState)
			return
		}
		//ユーザーIDを追加された行数番号で登録している
		id,_ := result.LastInsertId()
		_, _ = database.DB.Exec("INSERT INTO users (user_id VALUES (?)", id)
		resState.ResState = "success"
		resState.UserID = id
		json.NewEncoder(w).Encode(resState)
		fmt.Println("アカウントが作成されました")
		return
	}
	resState.ResState = "failed"
	resState.UserID = 0
	json.NewEncoder(w).Encode(resState)
}
