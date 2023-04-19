package handler

import (
	"net/http"
	"encoding/json"
	"TechRead/model"
	"fmt"
	"strconv"
	"TechRead/database"
)

func FetchProfileHandler(w http.ResponseWriter, req *http.Request) {
	var resState model.ResInfo
	var userRecord model.User
	var reqUserData model.User
	if err := json.NewDecoder(req.Body).Decode(&reqUserData); err != nil {
		fmt.Println(err)
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		ResFail(resState)
		json.NewEncoder(w).Encode(resState)
		return
	}
	cookies := req.Cookies()
	if cookies != nil{
		for _, c := range cookies {
			ID,_ := strconv.ParseInt(c.Value, 10, 64)
			fmt.Println(ID)
			rows,_ := database.DB.Query("SELECT user_id ,user_name,email, password FROM users WHERE user_id = ?", ID)
			for rows.Next() {
				err := rows.Scan(&userRecord.UserID,&userRecord.UserName,&userRecord.Email,&userRecord.Password)
				if err != nil{
					fmt.Println("エラーです")
				}
			userRecord.UserName,userRecord.Email,userRecord.Password = reqUserData.UserName,reqUserData.Email,reqUserData.Password
			fmt.Println(userRecord)
			fmt.Println(userRecord)
			resState.ResState = "success"
			resState.UserID = userRecord.UserID
			json.NewEncoder(w).Encode(resState)
			return
		}
	}
	ResFail(resState)
	json.NewEncoder(w).Encode(resState)
	}
}

