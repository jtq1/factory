package views

import (
	models "appTalleres"
	"fyne.io/fyne/v2"
)

type FrontManager struct {
	window  fyne.Window
	clients models.ClientService
}

func NewFrontManager(window fyne.Window, clients models.ClientService) *FrontManager {
	return &FrontManager{
		window:  window,
		clients: clients,
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
