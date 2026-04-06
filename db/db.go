package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

)

var DB

func InitDB() {
	sql.Open("sqlite3", "api.db")

}

