package model

type ResInfo struct {
	ResState  string   `db:"res_state" form:"res_state" json:"res_state"`
	UserID    int64      `db:"user_id" form:"user_id" json:"user_id"`
}
