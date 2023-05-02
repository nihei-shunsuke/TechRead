package model

import "time"

type Comment struct {
	ChapterID  int       `db:"chapter_id" form:"chapter_id" json:"chapter_id"`
	ChapterNum int       `db:"chapter_num" form:"chapter_num" json:"chapter_num"`
	Comment    string    `db:"comment" form:"comment" json:"comment"`
	UserID     int64       `db:"user_id" form:"user_id" json:"user_id"`
	CreatedAt  time.Time `db:"created_at" form:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" form:"updated_at" json:"updated_at"`
}
