package main

import (
	"appTalleres/backend/db"
	"appTalleres/frontend/views"
	"appTalleres/frontend/views/helper"
	"appTalleres/global"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

func main() {
	window := SetConfiguration()

	views.ShowLogin(window)
	window.ShowAndRun()
}

func SetConfiguration() fyne.Window {
	app := app.New()
	global.SetApp(app)
	window := app.NewWindow("Aragón Gestión")
	con, err := db.ConnectDB()
	if err != nil || con == nil {
		helper.CreateCriticalPopUp(window, "error trying to connect the DB", err)
	}

	app.Settings().SetTheme(theme.LightTheme())
	//window.Resize(fyne.NewSize(800, 600))
	//window.CenterOnScreen()
	window.SetFullScreen(true)

	return window
}
