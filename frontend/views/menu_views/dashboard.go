package menu_views

import (
	"appTalleres/frontend/views/helper"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowDashboard() fyne.CanvasObject {

	stats := container.NewGridWithColumns(3,
		createStatCard("Ventas Totales", "$10,000"),
		createStatCard("Clientes", "150"),
		createStatCard("Productos", "45"),
	)

	hbox := container.NewHBox(
		helper.CreateBorderCell(float32(30)), helper.CreateMenuTitle("Dashboard"),
	)

	borderCont := container.NewBorder(hbox, nil, nil, nil, helper.CreateBorderContainer(stats))

	return borderCont
}

func createStatCard(title string, value string) fyne.CanvasObject {
	return widget.NewCard(
		title,
		"",
		widget.NewLabel(value),
	)
}
