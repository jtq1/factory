package views

import (
	"appTalleres/backend/db"
	"appTalleres/frontend/views/menu_views"
	"appTalleres/global"
	"image/color"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func ShowMainWindow(window fyne.Window) {
	content := container.NewStack()

	menuItems := []string{
		"📊 Dashboard",
		"👥 Clientes",
		"📦 Productos",
		"💰 Ventas",
		"📄 Visor PDF",
		"🖨️ Imprimir",
		"⚙️ Configuración",
	}

	menuList := widget.NewList(
		func() int { return len(menuItems) },
		func() fyne.CanvasObject {
			return widget.NewLabel("Template")
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(menuItems[id])
		},
	)

	clientManager := db.NewClientDB(global.GetDB())

	menuList.Resize(fyne.NewSize(200, 0))
	content.Objects = []fyne.CanvasObject{menu_views.ShowDashboard()} // Select directly dashboard
	menuList.OnSelected = func(id widget.ListItemID) {
		switch id {
		case 0:
			content.Objects = []fyne.CanvasObject{menu_views.ShowDashboard()}
		case 1:
			content.Objects = []fyne.CanvasObject{menu_views.ShowClientList(clientManager)}
		case 2:
			content.Objects = []fyne.CanvasObject{menu_views.ShowProducts()}
		case 3:
			content.Objects = []fyne.CanvasObject{menu_views.ShowSales()}
		case 4:
			content.Objects = []fyne.CanvasObject{menu_views.ShowPDFViewer(window)}
		case 5:
			content.Objects = []fyne.CanvasObject{menu_views.ShowPrintFile(window)}
		case 6:
			content.Objects = []fyne.CanvasObject{menu_views.ShowSettings()}
		}
		content.Refresh()
	}

	// Barra superior
	topBar := createTopBar(window)

	// Layout principal usando container.NewBorder con un split fijo
	split := container.NewHSplit(menuList, content)
	split.SetOffset(0.2) // Fijar la división

	mainContent := container.NewBorder(
		topBar, nil, nil, nil,
		split,
	)

	window.SetContent(mainContent)

}

func createTopBar(window fyne.Window) fyne.CanvasObject {
	title := canvas.NewText("Panel de Control", color.NRGBA{R: 0, G: 100, B: 180, A: 255})
	title.TextSize = 20
	title.TextStyle.Bold = true

	logoutBtn := widget.NewButton("Cerrar Sesión", func() {
		ShowLogin(window)
	})

	return container.NewHBox(title, layout.NewSpacer(), container.NewHBox(logoutBtn, createCrossExitButton()))
}

func createCrossExitButton() fyne.CanvasObject {
	closeButton := widget.NewButton("X", func() {
		os.Exit(1)
	})

	// Align the button to the top-right
	overlay := container.New(layout.NewHBoxLayout(),
		layout.NewSpacer(),
		container.NewVBox(
			closeButton,
			layout.NewSpacer(),
		),
	)

	return overlay
}
