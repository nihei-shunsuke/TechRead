package handler

import (
	"TechRead/database"
	"TechRead/model"
	"encoding/json"
	"net/http"
	"time"
	"log"
)

type ReqChapterList struct {
	EventID       int       `db:"event_id" form:"event_id" json:"event_id"`
}

type ResChapterList struct {
	ResState   string   `db:"res_state" form:"res_state" json:"res_state"`
	EventID    int       `db:"event_id" form:"event_id" json:"event_id"`
	EventName  string    `db:"event_name" form:"event_name" json:"event_name"`
	Chapters   Chapters `json:"chapters"`
}

type Chapters struct {
	ChapterList []*ChapterList `json:"chapter_list"`
}

type ChapterList struct {
	ChapterID     int       `db:"chapter_id" form:"chapter_id" json:"chapter_id"`
	ChapterNum    int       `db:"chapter_num" form:"chapter_num" json:"chapter_num"`
	Venue         string    `db:"venue" form:"venue" json:"venue"`
	EventDate     time.Time `db:"event_date" form:"event_date" json:"event_date"`
	Content       string    `db:"content" form:"content" json:"content"`
}

func ChapterListHandler(w http.ResponseWriter, req *http.Request) {
	// リクエストを構造体に入れる
	var reqChapterList ReqChapterList
	if err := json.NewDecoder(req.Body).Decode(&reqChapterList); err != nil {
		log.SetFlags(log.Lshortfile)
    log.Println(err)

		var resState model.ResInfo
		ResFail(resState)
		json.NewEncoder(w).Encode(resState)
		return
	}

	// Chaptersを取得
	chapters, err := getChapterList(reqChapterList)
	if err != nil {
		log.SetFlags(log.Lshortfile)
    log.Println(err)

		var resState model.ResInfo
		ResFail(resState)
		json.NewEncoder(w).Encode(resState)
		return
	}
	var resChapterList ResChapterList
	resChapterList.Chapters = chapters

	// event関連のデータを取得
	const sqlEventStr = `SELECT event_id, event_name FROM events WHERE event_id = ?;`
	err = database.DB.QueryRow(sqlEventStr, reqChapterList.EventID).Scan(&resChapterList.EventID, &resChapterList.EventName)
	if err != nil {
		log.SetFlags(log.Lshortfile)
    log.Println(err)

		var resState model.ResInfo
		ResFail(resState)
		json.NewEncoder(w).Encode(resState)
		return
	}

	// 成功レスポンスを返す
	resChapterList.ResState = "success"
	json.NewEncoder(w).Encode(resChapterList)
}

// ChapterListを取得
func getChapterList(req ReqChapterList) (Chapters, error) {
	// クエリ文を定義
	const sqlStr = `SELECT chapter_id, event_date, venue, chapter_num, content FROM chapters WHERE event_id = ?;`

	// dbから取得
	rows, err := database.DB.Query(sqlStr, req.EventID)
	if err != nil {
		log.SetFlags(log.Lshortfile)
    log.Println(err)
		return Chapters{}, err
	}
	defer rows.Close()

	// dbから取り出した値をモデルに入れる
	var chapters Chapters
	for rows.Next() {
		var chapterList = &ChapterList{}
		err := rows.Scan(&chapterList.ChapterID, &chapterList.EventDate, &chapterList.Venue, &chapterList.ChapterNum, &chapterList.Content)
		if err != nil {
			log.SetFlags(log.Lshortfile)
			log.Println(err)
			return Chapters{}, err
		}
		chapters.ChapterList = append(chapters.ChapterList, chapterList)
	}

	err = rows.Err()
	if err != nil {
		log.SetFlags(log.Lshortfile)
    log.Println(err)
		return Chapters{}, err
	}

	return chapters, nil
}
