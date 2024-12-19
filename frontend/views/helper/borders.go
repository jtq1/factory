package helper

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func CreateBorderCell(size float32) fyne.CanvasObject {
	rect := canvas.NewRectangle(&color.NRGBA{})
	rect.SetMinSize(fyne.NewSize(size, size))
	return rect
}

func CreateDefaultBorderCell() fyne.CanvasObject {
	rect := canvas.NewRectangle(&color.NRGBA{})
	rect.SetMinSize(fyne.NewSize(20, 20))
	return rect
}

func CreateBorderContainer(objects ...fyne.CanvasObject) fyne.CanvasObject {
	rect := CreateDefaultBorderCell()
	return container.NewBorder(rect, rect, rect, rect, objects...)
}
