package menu_views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowSales() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel("Registro de Ventas"),
		widget.NewButton("Nueva Venta", nil),
		widget.NewLabel("Gráfico de ventas aquí"),
	)
}
