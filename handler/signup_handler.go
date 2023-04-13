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
		fmt.Println(err)
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		ResFail(resState)
		json.NewEncoder(w).Encode(resState)
		return
	}
	fmt.Println(reqUserData)
	//User構造体にクライアントが記入したデータを入れる処理を書く、またはその関数を記述する//
	const sqlStr = `
		insert into users (user_name, password, email) values
		(?,?,?);
	`
	//記入されたデータのuser_Emailに被りがないか調べる
	row := database.DB.QueryRow("SELECT email FROM users WHERE email = ?", reqUserData.Email).Scan(&reqUserData.Email)
	if row == sql.ErrNoRows {
		//Exec文で戻り値としてレコードを期待しないクエリを実行する
		//_をresultにすることで変更した行やデータを取得できる
		result, err := database.DB.Exec(sqlStr, reqUserData.UserName, reqUserData.Password, reqUserData.Email)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "fail internal exec \n", http.StatusInternalServerError)
			ResFail(resState)
			json.NewEncoder(w).Encode(resState)
			return
		}
		//ストリームにしてhttpレスポンスに出力している。
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
	fmt.Println("このメールアドレスは使用されています")
}
