package menu_views

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowProducts() fyne.CanvasObject {
	list := widget.NewList(
		func() int { return 10 },
		func() fyne.CanvasObject {
			return widget.NewLabel("Template")
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(fmt.Sprintf("Producto %d", id+1))
		},
	)

	return container.NewVBox(
		widget.NewLabel("Catálogo de Productos"),
		widget.NewButton("Nuevo Producto", nil),
		list,
	)
}
