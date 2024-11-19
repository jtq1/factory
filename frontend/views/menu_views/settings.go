package menu_views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowSettings() fyne.CanvasObject {
	form := widget.NewForm(
		widget.NewFormItem("Nombre de la empresa", widget.NewEntry()),
		widget.NewFormItem("Email de contacto", widget.NewEntry()),
		widget.NewFormItem("Moneda", widget.NewSelect([]string{"USD", "EUR", "MXN"}, nil)),
	)

	return container.NewVBox(
		widget.NewLabel("Configuración del Sistema"),
		form,
	)
}
