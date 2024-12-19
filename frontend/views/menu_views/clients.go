package menu_views

import (
	"appTalleres/frontend/interfaces"
	"appTalleres/frontend/views/helper"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func ShowClientList(window fyne.Window, clientMaster interfaces.ClientMaster) fyne.CanvasObject {
	clients, err := clientMaster.GetClients()
	if err != nil {
		helper.CreateErrorPopUp(window, "getClients error", err)
	}

	data := make([]string, len(clients))
	for i := range clients {
		data[i] = clients[i].Name
	}

	icon := widget.NewIcon(nil)
	label := widget.NewLabel("Clientes")
	hbox := container.NewHBox(icon, label)

	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(theme.DocumentIcon()), widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			if id == 5 || id == 6 {
				item.(*fyne.Container).Objects[1].(*widget.Label).SetText(data[id] + "\ntaller")
			} else {
				item.(*fyne.Container).Objects[1].(*widget.Label).SetText(data[id])
			}
		},
	)
	list.OnSelected = func(id widget.ListItemID) {
		label.SetText(data[id])
		icon.SetResource(theme.DocumentIcon())
	}
	list.OnUnselected = func(id widget.ListItemID) {
		label.SetText("Select An Item From The List")
		icon.SetResource(nil)
	}
	list.Select(125)
	list.SetItemHeight(5, 50)
	list.SetItemHeight(6, 50)

	return container.NewHSplit(list, container.NewCenter(hbox))
}
