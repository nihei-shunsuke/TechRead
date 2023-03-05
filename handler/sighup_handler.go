package handler

import (
	"TechRead/database"
	"TechRead/model"
	"encoding/json"
	"fmt"
	"net/http"
)

func SignUpHandler(w http.ResponseWriter, req *http.Request) {
	var reqUserData model.User
	//httpリクエスト(req.Body)に入ってるjsonデータ(model.User)を基にUser構造体にデコードする
	if err := json.NewDecoder(req.Body).Decode(&reqUserData); err != nil {
		fmt.Println(err)
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}
	fmt.Println(reqUserData)
	//User構造体にクライアントが記入したデータを入れる処理を書く、またはその関数を記述する//
	const sqlStr = `
		insert into users (user_name, password, email) values
		(?,?,?);
	`
	//記入されたデータのuser_Emailに被りがないか調べる
	row := database.DB.QueryRow("SELECT * FROM users WHERE email = ?", reqUserData.Email)
	if err := row.Err(); err != nil {
		//Exec文で戻り値としてレコードを期待しないクエリを実行する
		//_をresultにすることで変更した行やデータを取得できる
		_, err = database.DB.Exec(sqlStr, reqUserData.UserName, reqUserData.Password, reqUserData.Email)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "fail internal exec \n", http.StatusInternalServerError)
			return
		}
		//ストリームにしてhttpレスポンスに出力している。
		json.NewEncoder(w).Encode(reqUserData)
		return
	} else {
		fmt.Println("このメールアドレスは使用されています")
		return
	}
}
