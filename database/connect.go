package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)
var DB *sql.DB
func Connect() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
	}

	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := "3306"
	database_name := os.Getenv("MYSQL_DATABASE")

	dbconf := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database_name + "?parseTime=true"

	DB, err = sql.Open("mysql", dbconf)

	if err != nil {
		fmt.Println(err.Error())
	}
}
