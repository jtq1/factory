package main

import (
	"mi-aplicacion/frontend/views"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	window := myApp.NewWindow("Mi Aplicación")
	window.Resize(fyne.NewSize(800, 600))

	views.ShowLogin(window)
	window.ShowAndRun()
}
