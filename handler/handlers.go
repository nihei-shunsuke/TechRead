package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"fmt"
	"TechRead/database"
	"TechRead/model"
)

func LoginHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "LoginHandler\n")
}

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
	//記入されたデータのuser_idに被りがないか調べる
	row := database.DB.QueryRow("SELECT * FROM users WHERE user_id = ?",reqUserData.Email)
	if err := row.Err(); err != nil{
		_, err = database.DB.Exec(sqlStr, reqUserData.UserName, reqUserData.Password, reqUserData.Email)	//戻り値としてレコードを期待しないクエリを実行する
	/////////
		if err != nil{
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

func FetchProfileHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "FetchProfileHandler\n")
}

func EditProfileHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "EditProfileHandler\n")
}

func CreateEventHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "CreateEventHandler\n")
}

func EditEventHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "EditEventHandler\n")
}

func EventListHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "EventListHandler\n")
}

func EventDetailHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "EventDetailHandler\n")
}

func CreateChapterHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "CreateChapterHandler\n")
}

func EditChapterHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "EditChapterHandler\n")
}
