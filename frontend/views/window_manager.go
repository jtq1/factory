package views

import (
	models "appTalleres"

	"fyne.io/fyne/v2"
)

type FrontManager struct {
	window  fyne.Window
	clients models.ClientService
	events  models.EventService
}

func NewFrontManager(window fyne.Window, clients models.ClientService, events models.EventService) *FrontManager {
	return &FrontManager{
		window:  window,
		clients: clients,
		events:  events,
	}
}

func (fm *FrontManager) Run() {
	fm.ShowLogin()
}

func (fm *FrontManager) Window() fyne.Window {
	return fm.window
}

func (fm *FrontManager) SetContent(canvas fyne.CanvasObject) {
	fm.window.SetContent(canvas)
}
