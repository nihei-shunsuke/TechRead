package model

import "time"

type Event struct {
	EventID     int       `db:"event_id" form:"event_id" json:"event_id"`
	EventName   string    `db:"event_name" form:"event_name" json:"event_name"`
	BookName		string    `db:"book_name" form:"book_name" json:"book_name"`
	StartDate   time.Time `db:"start_date" form:"start_date" json:"start_date"`
	EndDate     time.Time `db:"end_date" form:"end_date" json:"end_date"`
	OrganizerID int64       `db:"organizer_id" form:"organizer_id" json:"organizer_id"`
	CreatedAt   time.Time `db:"created_at" form:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" form:"updated_at" json:"updated_at"`
}
