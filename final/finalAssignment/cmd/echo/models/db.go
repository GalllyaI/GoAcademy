package models

import "github.com/jmoiron/sqlx"

var db *sqlx.DB

func SetDatabase(database *sqlx.DB) {
	db = database
}
