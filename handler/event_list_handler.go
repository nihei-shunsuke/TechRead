package handler

import (
	"TechRead/database"
	"TechRead/model"
	"encoding/json"
	"fmt"
	"net/http"
	//"strconv"
)

type ResEventlist struct {
	Status    string
	Result    int
	EventID	  int
	EventName string
	BookName  string
	OrganizerID int64
}

func EventListHandler(w http.ResponseWriter, req *http.Request) {
	const sqlStr = `
	select event_id, event_name, book_name, organizer_id
	from events;
`
	//queryMap := req.URL.Query()
//
	//// クエリパラメータpageを取得
	var resState model.ResInfo
	var ResEvent ResEventlist
	//var page int
	//if p, ok := queryMap["page"]; ok && len(p) > 0 {
	//	var err error
	//	page, err = strconv.Atoi(p[0])
	//	if err != nil {
	//		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
	//		return
	//	}
	//	page = 1
	//}
	rows, err := database.DB.Query(sqlStr)
	if err != nil {
		fmt.Println(err)
		ResFail(resState)
		json.NewEncoder(w).Encode(resState)
		return
	}
	var EventlistData model.Event
	for rows.Next() {
		rows.Scan(&EventlistData.EventID, &EventlistData.EventName, &EventlistData.BookName, &EventlistData.OrganizerID)
		fmt.Println(EventlistData.BookName,EventlistData.EventName)
	}
	ResEvent.EventID,ResEvent.EventName,ResEvent.BookName,ResEvent.OrganizerID = EventlistData.EventID, EventlistData.EventName, EventlistData.BookName, EventlistData.OrganizerID
	ResEvent.Status,ResEvent.Result = "success",1
	json.NewEncoder(w).Encode(ResEvent)

}

