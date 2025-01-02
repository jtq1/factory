package main

import (
	models "appTalleres"
	"appTalleres/cache"
	"appTalleres/events"
	"appTalleres/frontend/views"
	"appTalleres/frontend/views/helper"
	"appTalleres/inmem"
	"appTalleres/managers"
	db "appTalleres/mysql"
	"database/sql"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

type Main struct {
	Config   Config
	Window   fyne.Window
	events   models.EventService
	clients  models.ClientService
	cacheMan models.CacheManager
}

func NewMain() *Main {
	return &Main{
		Config: DefaultConfig(),
	}
}

type Config struct {
	Cache struct {
		Enabled bool `json:"enabled"`
	}
	DB struct {
		Enabled bool `json:"enabled"`
	}
}

func DefaultConfig() Config {
	var config Config
	config.DB.Enabled = true
	config.Cache.Enabled = true
	return config
}

func main() {
	m := NewMain()
	window := m.setConfiguration()

	fm := views.NewFrontManager(window, m.clients, m.events)
	fm.Run()

	fm.Window().ShowAndRun()
}

func (m *Main) setConfiguration() fyne.Window {
	myapp := app.New()
	window := myapp.NewWindow("Aragón Gestión")
	m.Window = window

	myapp.Settings().SetTheme(theme.LightTheme())

	window.Resize(fyne.NewSize(float32(900), float32(700)))
	window.CenterOnScreen()

	var dbOb *sql.DB
	var err error
	if m.Config.DB.Enabled {
		dbOb, err = db.ConnectDB()
		if err != nil || dbOb == nil {
			helper.CreateCriticalPopUp(window, "error trying to connect the DB", err)
		}
	}

	// Set EventManager
	eventManager := events.NewEventManager()
	eventManager.Run()
	m.events = eventManager

	// Init Subscriptions
	m.subscribeEvents()

	// Set clientService to use MasterClient which is a combination of ClientDB and ClientCache with Validation
	cliCache := inmem.NewClientCache()
	cliDB := db.NewClientDB(dbOb)
	cliManager := managers.NewManagerClient(cliDB, cliCache, cliCache, eventManager, m.Config.DB.Enabled, m.Config.Cache.Enabled)
	m.clients = cliManager
	m.cacheMan = cliManager

	return window
}

func (m *Main) subscribeEvents() {
	if m.events != nil {
		err := m.events.Subscribe(models.LoginSuccessSubject, m.startLoggedServices)
		if err != nil {
			fmt.Printf("initial subscribing error %v", err)
		}
	}
}

func (m *Main) startLoggedServices() error {
	// Set CacheSync Manager
	if m.Config.Cache.Enabled {
		cacheManager := cache.NewCacheManager(m.cacheMan)
		cacheManager.SyncCache()
	}

	return nil
}
