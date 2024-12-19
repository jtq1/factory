package global

import (
	"database/sql"
)

var (
	globDB *sql.DB
)

func SetGlobalDB(db *sql.DB) {
	globDB = db
}

func GetDB() *sql.DB {
	return globDB
}
