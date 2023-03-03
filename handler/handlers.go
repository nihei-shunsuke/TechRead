package handler

import (
	"io"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "LoginHandler\n")
}

func SignUpHandler(w http.ResponseWriter, req *http.Request) {
	var reqUserData model.User
	//httpリクエスト(req.Body)に入ってるjsonデータ(model.User)を基にUser構造体にデコードする
	if err := json.NewDecoder(req.Body).Decode(&reqUserData); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}
	//User構造体にクライアントが記入したデータを入れる処理を書く、またはその関数を記述する//
	user := model.User
	const sqlStr = `
		insert into users (user_id, user_name, password, email, created_at) values
		(?,?,?,?,now());
	`
	//記入されたデータのuser_idに被りがないか調べる
	row := db.QueryRow("SELECT * FROM users WHERE user_id = ?",user.UserID)
	if err := row.Err(); err != nil{
		fmt.Println(err)
		return
	}
	//追加したデータや行が知りたい場合_をresultにして取り出せる。
	_, err := db.Exec(sqlStr, user.UserID, user.UserName, user.UserPass, user.UserEmail, user.CreatedAt)	//戻り値としてレコードを期待しないクエリを実行する
	/////////
	if err != nil{
		http.Error(w, "fail internal exec \n", http.StatusInternalServerError)
		return
	}
	//ストリームにしてhttpレスポンスに出力している。
	json.NewEncoder(w).Encode(userdata)
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

func CreateAbstractHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "CreateAbstractHandler\n")
}

func EditAbstractHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "EditAbstractHandler\n")
}
