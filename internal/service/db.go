package service

import "database/sql"

var db *sql.DB

func SetDB(database *sql.DB) {
	db = database
}
