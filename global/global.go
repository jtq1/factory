package global

import (
	"database/sql"

	"fyne.io/fyne/v2"
)

var (
	app    fyne.App
	globDB *sql.DB
)

func SetGlobalDB(db *sql.DB) {
	globDB = db
}

func GetDB() *sql.DB {
	return globDB
}

func SetApp(appIn fyne.App) {
	app = appIn
}

func GetApp() fyne.App {
	return app
}
