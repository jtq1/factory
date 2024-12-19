package helper

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func CreateMenuTitle(titleStr string) fyne.CanvasObject {
	title := canvas.NewText(titleStr, color.NRGBA{R: 0, G: 100, B: 180, A: 255})
	title.TextSize = 20
	title.TextStyle.Bold = true

	return title
}
