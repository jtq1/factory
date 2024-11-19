package menu_views

import (
	"fmt"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func ShowPDFViewer(window fyne.Window) fyne.CanvasObject {
	pageNum := 0
	var totalPages int
	var currentPDFPath string
	var tempDir string

	// Controles de navegación
	prevButton := widget.NewButton("Anterior", nil)
	nextButton := widget.NewButton("Siguiente", nil)
	pageLabel := widget.NewLabel("Página: 0 / 0")

	// Contenedor para la página actual
	pageContainer := container.NewCenter(widget.NewLabel("No hay PDF cargado"))

	// Función para mostrar una página específica
	showPage := func(pageIndex int) {
		if currentPDFPath == "" {
			return
		}

		imagePath := filepath.Join(tempDir, fmt.Sprintf("page_%d.jpg", pageIndex+1))

		// Leer la imagen
		reader, err := os.Open(imagePath)
		if err != nil {
			dialog.ShowError(err, window)
			return
		}
		defer reader.Close()

		// Crear y mostrar la imagen
		canvasImg := canvas.NewImageFromFile(imagePath)
		canvasImg.FillMode = canvas.ImageFillOriginal
		canvasImg.Resize(fyne.NewSize(600, 800))

		pageContainer.Objects = []fyne.CanvasObject{canvasImg}
		pageContainer.Refresh()

		// Actualizar etiqueta de página
		pageLabel.SetText(fmt.Sprintf("Página: %d / %d", pageIndex+1, totalPages))
	}

	// Configurar botones de navegación
	prevButton.OnTapped = func() {
		if pageNum > 0 {
			pageNum--
			showPage(pageNum)
		}
	}

	nextButton.OnTapped = func() {
		if pageNum < totalPages-1 {
			pageNum++
			showPage(pageNum)
		}
	}

	// Botón para seleccionar PDF
	selectButton := widget.NewButton("Seleccionar PDF", func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, window)
				return
			}
			if reader == nil {
				return
			}
			defer reader.Close()

			// Crear directorio temporal para las imágenes
			var err2 error
			tempDir, err2 = os.MkdirTemp("", "pdf_viewer_*")
			if err2 != nil {
				dialog.ShowError(err2, window)
				return
			}

			// Guardar la ruta del PDF
			currentPDFPath = reader.URI().Path()

			// Configuración para la extracción de imágenes
			conf := model.NewDefaultConfiguration()
			conf.ValidationMode = model.ValidationRelaxed

			// Extraer todas las páginas como imágenes
			err = api.ExtractImagesFile(currentPDFPath, tempDir, nil, conf)
			if err != nil {
				dialog.ShowError(err, window)
				return
			}

			// Obtener número total de páginas
			ctx, err := api.ReadContextFile(currentPDFPath)
			if err != nil {
				dialog.ShowError(err, window)
				return
			}
			totalPages = ctx.PageCount

			// Resetear página y mostrar primera página
			pageNum = 0
			showPage(pageNum)
		}, window)

		fd.SetFilter(storage.NewExtensionFileFilter([]string{".pdf"}))
		fd.Show()
	})

	controls := container.NewHBox(
		prevButton,
		pageLabel,
		nextButton,
	)

	return container.NewVBox(
		widget.NewLabel("Visor de PDF"),
		selectButton,
		controls,
		pageContainer,
	)
}
