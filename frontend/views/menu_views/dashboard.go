package menu_views

import (
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

	return container.NewVBox(
		widget.NewLabel("Dashboard"),
		stats,
	)
}

func createStatCard(title string, value string) fyne.CanvasObject {
	return widget.NewCard(
		title,
		"",
		widget.NewLabel(value),
	)
}
