package menu_views

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowClients() fyne.CanvasObject {
	table := widget.NewTable(
		func() (int, int) { return 10, 4 },
		func() fyne.CanvasObject {
			return widget.NewLabel("Cell")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			label := o.(*widget.Label)
			switch i.Col {
			case 0:
				label.SetText(fmt.Sprintf("Cliente %d", i.Row+1))
			case 1:
				label.SetText("email@ejemplo.com")
			case 2:
				label.SetText("123-456-789")
			case 3:
				label.SetText("Dirección ejemplo")
			}
		},
	)

	return container.NewVBox(
		widget.NewLabel("Gestión de Clientes"),
		widget.NewButton("Nuevo Cliente", nil),
		table,
	)
}
