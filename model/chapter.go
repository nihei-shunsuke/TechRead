package model

import "time"

type Chapter struct {
	ChapterID     int       `db:"chapter_id" form:"chapter_id" json:"chapter_id"`
	ChapterNum    int       `db:"chapter_num" form:"chapter_num" json:"chapter_num"`
	Venue         string    `db:"venue" form:"venue" json:"venue"`
	EventID       int       `db:"event_id" form:"event_id" json:"event_id"`
	EventDate     time.Time `db:"event_date" form:"event_date" json:"event_date"`
	Content       int       `db:"content" form:"content" json:"content"`
	LastUpdaterID int       `db:"last_updater_id" form:"last_updater_id" json:"last_updater_id"`
	CreatedAt     time.Time `db:"created_at" form:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" form:"updated_at" json:"updated_at"`
}
