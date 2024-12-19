package helper

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func CreateMenuTitle(titleStr string) *fyne.Container {
	title := canvas.NewText(titleStr, color.NRGBA{R: 0, G: 100, B: 180, A: 255})
	title.TextSize = 20
	title.TextStyle.Bold = true

	hbox := container.NewHBox(
		CreateBorderCell(float32(25)), title,
	)

	return hbox
}
