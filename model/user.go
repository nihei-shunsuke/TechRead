package model

import "time"

type User struct {
	UserID    int64       `db:"user_id" form:"user_id" json:"user_id"`
	UserName  string    `db:"user_name" form:"user_name" json:"user_name"`
	Email     string    `db:"email" form:"email" json:"email"`
	Password  string    `db:"password" form:"password" json:"password"`
	CreatedAt time.Time `db:"created_at" form:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" form:"updated_at" json:"updated_at"`
}
