package handler

import (
	"TechRead/database"
	"TechRead/model"
	"encoding/json"
	"fmt"
	"net/http"
	//"strconv"
)
//リストでそのユーザーの参加した、してるイベントの詳細(イベントID、名前、本の名前、メンバーの数)を表示する
//クッキーからuseridを取る→event-groupテーブルからeventidとuseridの数を取得→eventidからbooknameを取得
type ResEventlist struct {
	Status    string
	event	  []eventdetail
}
type eventdetail struct {
	EventID		int
	BookName	string
	MemberCount	int
}

func EventListHandler(w http.ResponseWriter, req *http.Request) {
	var userRecord model.User
	var resEvent ResEventlist
	var loopNum int
	//cookies := req.Cookies()
	//if cookies != nil{
	//	for _, c := range cookies {
	//		ID,_ := strconv.ParseInt(c.Value, 10, 64)
	//		fmt.Println(ID)
	//		_ = database.DB.QueryRow("SELECT user_id FROM users WHERE user_id = ?", ID).Scan(&userRecord.UserID)
	//	}
	//}
	userRecord.UserID = 1
	rows,err := database.DB.Query("SELECT event_id FROM event_groups WHERE user_id = ? ",userRecord.UserID)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		i := 0
		fmt.Println(b)
		rows.Scan(&resEvent.event[i].EventID)
		i = i + 1
		loopNum = i
	}
	fmt.Println(c)
	for i := 0; i < loopNum; i++ {
		err := database.DB.QueryRow("SELECT book_name FROM events WHERE event_id = ?", resEvent.event[i].EventID).Scan(&resEvent.event[i].BookName)
		if err != nil {
			var resState model.ResInfo
			fmt.Println(err)
			ResFail(resState)
			json.NewEncoder(w).Encode(resState)
			return
		}
		fmt.Println(d)
		err = database.DB.QueryRow("SELECT COUNT ( user_id ) FROM event_groups WHERE event_id = ?", resEvent.event[i].EventID).Scan(&resEvent.event[i].MemberCount)
		if err != nil {
			var resState model.ResInfo
			fmt.Println(err)
			ResFail(resState)
			json.NewEncoder(w).Encode(resState)
			return
		}
	}
	fmt.Println(e)
	resEvent.Status = "success"
	json.NewEncoder(w).Encode(resEvent)

}