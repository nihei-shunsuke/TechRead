<<<<<<< Updated upstream
package model

import "time"

type EventGroup struct {
	UserID    int64       `db:"user_id" form:"user_id" json:"user_id"`
	EventID   int       `db:"event_id" form:"event_id" json:"event_id"`
	CreatedAt time.Time `db:"created_at" form:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" form:"updated_at" json:"updated_at"`
}
=======
package model

import "time"

type EventGroup struct {
	UserID    int       `db:"user_id" form:"user_id" json:"user_id"`
	EventID   int       `db:"event_id" form:"event_id" json:"event_id"`
	CreatedAt time.Time `db:"created_at" form:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" form:"updated_at" json:"updated_at"`
}
>>>>>>> Stashed changes
