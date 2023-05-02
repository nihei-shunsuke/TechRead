package handler

import (
	"net/http"
	"TechRead/model"
	"encoding/json"
	"fmt"
	"strconv"
	"TechRead/database"
)

func CreateEventHandler(w http.ResponseWriter, req *http.Request) {
	var reqEventData model.Event
	var resState model.ResInfo
	var userRecord model.User
	if err := json.NewDecoder(req.Body).Decode(&reqEventData); err != nil {
		fmt.Println(err)
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		ResFail(resState)
		json.NewEncoder(w).Encode(resState)
		return
	}
	fmt.Println(reqEventData)
	const sqlStr = `
		insert into events (event_name, book_name,organizer_id) values
		(?,?,?);
	`
	cookies := req.Cookies()
	if cookies != nil{
		for _, c := range cookies {
			ID,_ := strconv.ParseInt(c.Value, 10, 64)
			fmt.Println(ID)
			_ = database.DB.QueryRow("SELECT user_id FROM users WHERE user_id = ?", ID).Scan(&userRecord.UserID)
		}
		fmt.Println(userRecord.UserID)
		reqEventData.OrganizerID = userRecord.UserID
		result, err := database.DB.Exec(sqlStr, reqEventData.EventName, reqEventData.BookName,userRecord.UserID)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "fail internal exec \n", http.StatusInternalServerError)
			ResFail(resState)
			json.NewEncoder(w).Encode(resState)
			return
		}
		//ストリームにしてhttpレスポンスに出力している。
		id,_ := result.LastInsertId()
		_, err = database.DB.Exec("INSERT INTO events (event_id (?)", id,)
		resState.ResState = "success"
		resState.UserID = userRecord.UserID
		fmt.Println(id)
		json.NewEncoder(w).Encode(resState)
	}
}
