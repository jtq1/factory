package views

import (
	"appTalleres/frontend/views/helper"
	"appTalleres/frontend/views/menu_views"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func (fm *FrontManager) ShowMainWindow() {
	content := container.NewStack()

	menuItems := []string{
		"  📊 Dashboard",
		"  👥 Clientes",
		"  📦 Productos",
		"  💰 Ventas",
		"  📄 Visor PDF",
		"  🖨️ Imprimir",
		"  ⚙️ Configuración",
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

	/* for i := 0; i < 10; i++ {
		id, err := clientManager.CreateClient(models.Client{
			Name:    fmt.Sprintf("Nombre %d", i),
			Email:   fmt.Sprintf("Mail %d", i),
			Phone:   fmt.Sprintf("%09d", rand.Intn(1000000000)),
			Address: fmt.Sprintf("Address %d", i),
		})
		if err != nil {
			helper.CreateErrorPopUp(window, fmt.Sprintf("error create client %d", id), err)
		}
	} */

	clientView := menu_views.NewClientView(fm.Window(), fm.clients, fm.events)

	menuList.Resize(fyne.NewSize(200, 0))
	menuList.Select(0)
	content.Objects = []fyne.CanvasObject{menu_views.ShowDashboard()} // Select directly dashboard
	menuList.OnSelected = func(id widget.ListItemID) {
		switch id {
		case 0:
			content.Objects = []fyne.CanvasObject{menu_views.ShowDashboard()}
		case 1:
			content.Objects = []fyne.CanvasObject{clientView.ShowClientList()}
		case 2:
			content.Objects = []fyne.CanvasObject{menu_views.ShowProducts()}
		case 3:
			content.Objects = []fyne.CanvasObject{menu_views.ShowSales()}
		case 4:
			content.Objects = []fyne.CanvasObject{menu_views.ShowPDFViewer(fm.Window())}
		case 5:
			content.Objects = []fyne.CanvasObject{menu_views.ShowPrintFile(fm.Window())}
		case 6:
			content.Objects = []fyne.CanvasObject{menu_views.ShowSettings()}
		}
		content.Refresh()
	}

	// Barra superior
	topBar := fm.createTopBar()

	// Layout principal usando container.NewBorder con un split fijo
	split := container.NewHSplit(menuList, content)
	split.SetOffset(0.2)

	mainContent := container.NewBorder(
		topBar, nil, nil, nil,
		split,
	)

	fm.SetContent(mainContent)
}

func (fm *FrontManager) createTopBar() fyne.CanvasObject {
	title := canvas.NewText("Panel de Control", color.NRGBA{R: 0, G: 100, B: 180, A: 255})
	title.TextSize = 20
	title.TextStyle.Bold = true

	border := helper.CreateBorderCell(float32(40))

	logoutBtn := widget.NewButton("Cerrar Sesión", func() {
		fm.ShowLogin()
	})

	return container.NewHBox(border, title, layout.NewSpacer(), container.NewHBox(logoutBtn, helper.CreateManagementButtons(fm.Window())))
}
