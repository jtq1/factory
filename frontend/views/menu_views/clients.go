package menu_views

import (
	models "appTalleres"
	"appTalleres/frontend/views/helper"
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type clientView struct {
	clientsListData []string
	clientsList     *widget.List
	clientMaster    models.ClientService
	window          fyne.Window
	rendered        fyne.CanvasObject
	eventManager    models.EventService
}

func NewClientView(window fyne.Window, cm models.ClientService, ev models.EventService) *clientView {
	cv := &clientView{
		clientMaster: cm,
		window:       window,
		eventManager: ev,
	}

	cv.subscribeEvents()
	return cv
}

func (cv *clientView) subscribeEvents() {
	if cv.eventManager != nil {
		err := cv.eventManager.Subscribe(models.ClientCacheRefreshedSubject, cv.RefreshClientsList)
		if err != nil {
			fmt.Printf("clientManager subscribing error %v", err)
		}
	}
}

func (cv *clientView) ShowClientList() fyne.CanvasObject {
	cv.RefreshClientsList()
	if cv.rendered != nil {
		return cv.rendered
	}

	titleHBox := helper.CreateMenuTitle("Clientes")
	icon := widget.NewIcon(nil)
	label := widget.NewLabel("Clientes")
	hbox := container.NewHBox(icon, label)

	cv.clientsList = widget.NewList(
		func() int {
			return len(cv.clientsListData)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(theme.DocumentIcon()), widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			if id == 5 || id == 6 {
				item.(*fyne.Container).Objects[1].(*widget.Label).SetText(cv.clientsListData[id] + "\ntaller")
			} else {
				item.(*fyne.Container).Objects[1].(*widget.Label).SetText(cv.clientsListData[id])
			}
		},
	)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("Entra")
			cv.clientsListData = append(cv.clientsListData, fmt.Sprintf("Test %d", i))
			cv.clientsList.Refresh()
			time.Sleep(10 * time.Second)
		}
	}()

	cv.clientsList.OnSelected = func(id widget.ListItemID) {
		label.SetText(cv.clientsListData[id])
		icon.SetResource(theme.DocumentIcon())
	}
	cv.clientsList.OnUnselected = func(id widget.ListItemID) {
		label.SetText("Select An Item From The List")
		icon.SetResource(nil)
	}
	cv.clientsList.Select(125)
	cv.clientsList.SetItemHeight(5, 50)
	cv.clientsList.SetItemHeight(6, 50)

	borderCont := helper.CreateCustomBorderContainer(10, cv.clientsList)

	clientContainer := container.NewBorder(titleHBox, nil, nil, nil, borderCont)
	split := container.NewHSplit(clientContainer, container.NewCenter(hbox))
	split.SetOffset(0.3)

	cv.rendered = split

	return split
}

func (cv *clientView) RefreshClientsList() error {
	clients, err := cv.clientMaster.GetClients()
	if err != nil {
		if cv.window != nil {
			helper.CreateErrorPopUp(cv.window, "getClients error", err)
		}
		fmt.Println("error fetching client list")
		return err
	}

	cv.clientsListData = make([]string, len(clients))
	for i := range clients {
		cv.clientsListData[i] = clients[i].Name
	}

	if cv.clientsList != nil {
		cv.clientsList.Refresh()
	}

	return nil
}
