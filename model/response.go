<<<<<<< Updated upstream
package model

type ResInfo struct {
	ResState  string   `db:"res_state" form:"res_state" json:"res_state"`
	UserID    int64      `db:"user_id" form:"user_id" json:"user_id"`
}
=======
package model

type ResInfo struct {
	ResState  string   `db:"res_state" form:"res_state" json:"res_state"`
	UserID    int64      `db:"user_id" form:"user_id" json:"user_id"`
}

type ResCreateEvent struct {
	ResState string `db:"res_state" form:"res_state" json:"res_state"`
	EventID	int		`db:"event_id" form:"event_id" json:"event_id"`
}

type ResCreateChapter struct {
	ResState  string   `db:"res_state" form:"res_state" json:"res_state"`
	ChapterID int64    `db:"chapter_id" form:"res_state" json:"chapter_id"`
}
>>>>>>> Stashed changes
