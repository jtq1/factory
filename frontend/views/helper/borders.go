package helper

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func CreateBorderCell(size float32) fyne.CanvasObject {
	rect := canvas.NewRectangle(&color.NRGBA{R: 0, G: 100, B: 180, A: 255})
	rect.SetMinSize(fyne.NewSize(size, size))
	return rect
}

func CreateDefaultBorderCell() fyne.CanvasObject {
	rect := canvas.NewRectangle(&color.NRGBA{R: 0, G: 100, B: 180, A: 255})
	rect.SetMinSize(fyne.NewSize(20, 20))
	return rect
}

func CreateBorderContainer(objects ...fyne.CanvasObject) *fyne.Container {
	rect := CreateDefaultBorderCell()
	return container.NewBorder(nil, rect, rect, rect, objects...)
}

func CreateCustomBorderContainer(border int, objects ...fyne.CanvasObject) *fyne.Container {
	rect := CreateBorderCell(float32(border))
	return container.NewBorder(nil, rect, rect, rect, objects...)
}
