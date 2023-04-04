package handler

import (
	"net/http"
	"encoding/json"
	"TechRead/model"
	"fmt"
	"TechRead/database"
)

func FetchProfileHandler(w http.ResponseWriter, req *http.Request) {
	var resState model.ResInfo
	var userRecord model.User
	cookies := req.Cookies()
	if cookies != nil{
		for _, c := range cookies {
			email := c.Value
			rows,_ := database.DB.Query("SELECT user_id ,email, password FROM users WHERE email = ?", email)
			for rows.Next() {
				err := rows.Scan(&userRecord.UserID,&userRecord.Email,&userRecord.Password)
				if err != nil{
					fmt.Println("エラーです")
				}
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
