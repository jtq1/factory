package menu_views

import (
	"appTalleres/frontend/views/helper"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowDashboard() fyne.CanvasObject {
	titleHBox := helper.CreateMenuTitle("Dashboard")

	stats := container.NewGridWithColumns(3,
		createStatCard("Ventas Totales", "$10,000"),
		createStatCard("Clientes", "150"),
		createStatCard("Productos", "45"),
	)

	borderCont := container.NewVBox(titleHBox, helper.CreateBorderContainer(stats))

	return borderCont
}

func createStatCard(title string, value string) fyne.CanvasObject {
	return widget.NewCard(
		title,
		"",
		widget.NewLabel(value),
	)
}
