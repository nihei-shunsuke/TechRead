package handler

import (
	"TechRead/database"
	"TechRead/model"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ReqChapter struct {
	ChapterNum    int       `db:"chapter_num" form:"chapter_num" json:"chapter_num"`
	Venue         string    `db:"venue" form:"venue" json:"venue"`
	EventID       int       `db:"event_id" form:"event_id" json:"event_id"`
	EventDate     time.Time `db:"event_date" form:"event_date" json:"event_date"`
	Content       string       `db:"content" form:"content" json:"content"`
}

func CreateChapterHandler(w http.ResponseWriter, req *http.Request) {
	// リクエストを構造体に入れる
	var reqChapter ReqChapter
	if err := json.NewDecoder(req.Body).Decode(&reqChapter); err != nil {
		fmt.Println(err)

		var resState model.ResInfo
		ResFail(resState)
		json.NewEncoder(w).Encode(resState)
		return
	}

	// TODO: フロントエンド側でcookieを入れる処理ができたら、cookieの値を読み取る処理に変更する
	user_id := 1

	// クエリ文を定義
	const sqlStr = `
		INSERT into chapters (chapter_num, venue, event_id, event_date, content, last_updater_id, created_at, updated_at) value
		(?, ?, ?, ?, ?, ?, now(), now());
	`

	// dbに入れる処理
	result, err := database.DB.Exec(sqlStr, reqChapter.ChapterNum, reqChapter.Venue, reqChapter.EventID, reqChapter.EventDate, reqChapter.Content, user_id)
	if err != nil {
		fmt.Println(err)

		var resState model.ResInfo
		ResFail(resState)
		json.NewEncoder(w).Encode(resState)
		return
	}

	// 保存したchapter_idを取得
	chapterId, _ := result.LastInsertId()
	fmt.Println(chapterId)

	// 成功のレスポンスを返す
	var resSuccess model.ResCreateChapter
	resSuccess.ResState = "success"
	resSuccess.ChapterID = chapterId
	json.NewEncoder(w).Encode(resSuccess)
}
