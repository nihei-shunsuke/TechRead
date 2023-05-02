package handler

import (
	"net/http"
	"TechRead/model"
	"encoding/json"
	"fmt"
	"strconv"
	"TechRead/database"
)

type reqEvent struct {
	EventName  string
	BookName	string
	MemberList	[]string  `db:"memberlist" form:"memberlist" json:"memberlist"`
}

func CreateEventHandler(w http.ResponseWriter, req *http.Request) {
	var reqEventData reqEvent
	var resState model.ResInfo
	var userRecord model.User
	if err := json.NewDecoder(req.Body).Decode(&reqEventData); err != nil {
		fmt.Println("a")
		ResFail(resState)
		json.NewEncoder(w).Encode(resState)
		return
	}
	const sqlStr = `
		insert into events (event_name, book_name, organaized_id) values
		(?,?,?);
	`
	const sqlStrMember = `
		insert into event_groups (user_id, event_id) values
		(?,?);
	`
	cookies := req.Cookies()
	if cookies != nil{
		for _, c := range cookies {
			ID,_ := strconv.ParseInt(c.Value, 10, 64)
			fmt.Println(ID)
			_ = database.DB.QueryRow("SELECT user_id FROM users WHERE user_id = ?", ID).Scan(&userRecord.UserID)
		}
		result, _ := database.DB.Exec(sqlStr, reqEventData.EventName, reqEventData.BookName,userRecord.UserID)
		//if err != nil {
		//	fmt.Println(err)
		//	ResFail(resState)
		//	json.NewEncoder(w).Encode(resState)
		//	return
		//}
		id,_ := result.LastInsertId()
		for i := 0; i < len(reqEventData.MemberList); i++{
			_ = database.DB.QueryRow("SELECT user_id FROM users WHERE user_name = ?", reqEventData.MemberList[i]).Scan(&userRecord.UserID)
			_, err := database.DB.Exec(sqlStrMember, userRecord.UserID, id)
			if err != nil{
				fmt.Println("c")
				ResFail(resState)
				json.NewEncoder(w).Encode(resState)
				return
			}
		}
		resState.ResState = "success"
		resState.UserID = id//ユーザーIDと書いてあるがイベントID
		json.NewEncoder(w).Encode(resState)
	}
	ResFail(resState)
	fmt.Println("d")
	json.NewEncoder(w).Encode(resState)
}
