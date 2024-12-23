package main

import (
	models "appTalleres"
	"appTalleres/frontend/views"
	"appTalleres/frontend/views/helper"
	"appTalleres/inmem"
	"appTalleres/managers"
	db "appTalleres/mysql"
	"database/sql"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

type Main struct {
	Config  Config
	Window  fyne.Window
	clients models.ClientService
}

type Config struct {
	Cache struct {
		Enabled bool `json:"enabled"`
	}
	DB struct {
		Enabled bool `json:"enabled"`
	}
}

func NewMain() *Main {
	return &Main{
		Config: DefaultConfig(),
	}
}

func main() {
	m := NewMain()
	window := m.setConfiguration()

	fm := views.NewFrontManager(window, m.clients)
	fm.Run()

	fm.Window().ShowAndRun()
}

func DefaultConfig() Config {
	var config Config
	config.DB.Enabled = false
	config.Cache.Enabled = true
	return config
}

func (m *Main) setConfiguration() fyne.Window {
	myapp := app.New()
	window := myapp.NewWindow("Aragón Gestión")
	m.Window = window

	myapp.Settings().SetTheme(theme.LightTheme())

	window.Resize(fyne.NewSize(float32(900), float32(700)))
	window.CenterOnScreen()

	var dbOb *sql.DB
	if m.Config.DB.Enabled {
		dbOb, err := db.ConnectDB()
		if err != nil || dbOb == nil {
			helper.CreateCriticalPopUp(window, "error trying to connect the DB", err)
		}
	}

	// Set clientService to use MasterClient which is a combination of ClientDB and ClientCache with Validation
	cliCache := inmem.NewClientCache()
	cliDB := db.NewClientDB(dbOb)
	m.clients = managers.NewManagerClient(cliDB, cliCache, m.Config.DB.Enabled, m.Config.Cache.Enabled)

	return window
}
