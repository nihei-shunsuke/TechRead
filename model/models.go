package model

type User struct{
	UserID string		`json:"user_id"`
	UserName string		`json:"user_name"`
	UserPass string		`json:"password"`
	UserEmail string	`json:"email"`
}