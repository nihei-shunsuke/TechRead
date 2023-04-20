package handler

import (
	"TechRead/model"
)

func ResFail(m model.ResInfo){
	m.ResState = "failed"
	m.UserID = 0
	return
}