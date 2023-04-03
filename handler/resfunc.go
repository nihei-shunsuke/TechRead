package handler

import (
	"TechRead/model"
)

func ResFail(m model.ResInfo){
	m.ResState = "fail"
	m.UserID = 0
	return
}