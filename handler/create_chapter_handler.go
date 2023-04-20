package handler

import (
	"net/http"
	"TechRead/model"
	"encoding/json"
	"fmt"
	"strconv"
	"TechRead/database"
)

func CreateChapterHandler(w http.ResponseWriter, req *http.Request) {
	var reqChapterData model.Chapter
	var resState model.ResInfo
	var userRecord model.User
	if err := json.NewDecoder(req.Body).Decode(&reqChapterData); err != nil {
		fmt.Println(err)
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		ResFail(resState)
		json.NewEncoder(w).Encode(resState)
		return
	}
	fmt.Println(reqChapterData)
	const sqlStr = `
		insert into chapters (venue,event_id,content) values
		(?,?,?);
	`
	cookies := req.Cookies()
	if cookies != nil{
		for _, c := range cookies {
			ID,_ := strconv.ParseInt(c.Value, 10, 64)
			fmt.Println(ID)
			_ = database.DB.QueryRow("SELECT user_id FROM users WHERE user_id = ?", ID).Scan(&userRecord.UserID)
		}
		result, err := database.DB.Exec(sqlStr, reqChapterData.Venue, 1,reqChapterData.Content)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "fail internal exec \n", http.StatusInternalServerError)
			ResFail(resState)
			json.NewEncoder(w).Encode(resState)
			return
		}
		id,_ := result.LastInsertId()
		_, err = database.DB.Exec("INSERT INTO chapters (chapter_num (?)", id,)
		resState.ResState = "success"
		resState.UserID = userRecord.UserID
		fmt.Println(id)
		json.NewEncoder(w).Encode(resState)
	}
}