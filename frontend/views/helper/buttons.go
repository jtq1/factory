package helper

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"os"
)

func CreateManagementButtons(window fyne.Window) *fyne.Container {
	closeButton := widget.NewButton("X", func() {
		os.Exit(1)
	})

	miniMax := widget.NewButton("🗖", func() {
		window.SetFullScreen(!window.FullScreen())
	})

	// Align the button to the top-right
	overlay := container.New(layout.NewHBoxLayout(),
		layout.NewSpacer(),
		container.NewHBox(
			miniMax,
			closeButton,
		),
	)

	return overlay
}
