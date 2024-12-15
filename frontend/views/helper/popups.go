package helper

import (
	"image/color"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
)

func CreateInfoPopUp(window fyne.Window, msg string, err error) {
	dialog.ShowInformation("Popup Title", "This is a popup message!", window)
}

func CreateConfirmPopUp(window fyne.Window, msg string, fun func()) {
	confDialog := dialog.NewConfirm("Info", msg, func(confirmed bool) {
		if confirmed {
			fun()
		}
	}, window)

	confDialog.Show()
}

func CreateErrorPopUp(window fyne.Window, msg string, err error) {
	dialog.ShowError(err, window)
}

func CreateCriticalPopUp(window fyne.Window, msg string, err error) {
	title := canvas.NewText(msg, color.NRGBA{R: 0, G: 100, B: 180, A: 255})
	title.TextSize = 30
	title.TextStyle.Bold = true
	errorLine := canvas.NewText("Error: "+err.Error(), color.NRGBA{R: 0, G: 100, B: 180, A: 255})
	errorLine.TextSize = 20
	errorLine.TextStyle.Bold = true
	box := container.NewVBox(title, errorLine)

	popup := dialog.NewCustom("Error", "Close", box, window)
	popup.SetOnClosed(func() {
		os.Exit(1)
	})
	popup.Show()
}
