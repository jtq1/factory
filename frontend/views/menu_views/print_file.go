package menu_views

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func ShowPrintFile(window fyne.Window) fyne.CanvasObject {
	var selectedFilePath string
	selectedFile := widget.NewLabel("Ningún archivo seleccionado")

	// Obtener lista de impresoras del sistema
	printers := getPrinters()
	selectedPrinter := widget.NewSelect(printers, nil)
	if len(printers) > 0 {
		selectedPrinter.SetSelected(printers[0])
	}

	selectButton := widget.NewButton("Seleccionar Archivo", func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, window)
				return
			}
			if reader == nil {
				return
			}
			selectedFilePath = reader.URI().Path()
			selectedFile.SetText("Archivo seleccionado: " + reader.URI().Name())
		}, window)
	})

	printButton := widget.NewButton("Imprimir", func() {
		if selectedFilePath == "" {
			dialog.ShowInformation("Error", "Por favor seleccione un archivo", window)
			return
		}

		if selectedPrinter.Selected == "" {
			dialog.ShowInformation("Error", "Por favor seleccione una impresora", window)
			return
		}

		go func() {
			err := printFile(selectedFilePath, selectedPrinter.Selected)
			if err != nil {
				dialog.ShowError(err, window)
				return
			}
			dialog.ShowInformation("Éxito", "Documento enviado a imprimir", window)
		}()
	})

	return container.NewVBox(
		widget.NewLabel("Imprimir Archivo"),
		selectButton,
		selectedFile,
		widget.NewLabel("Seleccionar Impresora:"),
		selectedPrinter,
		printButton,
	)
}

func getPrinters() []string {
	switch runtime.GOOS {
	case "windows":
		/*names, err := printer.ReadNames()
		if err != nil {
			return []string{"Impresora predeterminada"}
		}
		return names*/

	case "darwin", "linux": // macOS and Linux
		cmd := exec.Command("lpstat", "-p")
		output, err := cmd.Output()
		if err != nil {
			return []string{"Impresora predeterminada"}
		}

		// Parse printer names from lpstat output
		printers := []string{}
		lines := strings.Split(string(output), "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "printer ") {
				// Format is typically "printer PrinterName"
				parts := strings.Fields(line)
				if len(parts) >= 2 {
					printers = append(printers, parts[1])
				}
			}
		}

		if len(printers) == 0 {
			return []string{"Impresora predeterminada"}
		}
		return printers

	default:
		return []string{"Impresora predeterminada"}
	}

	return []string{"Impresora predeterminada"}
}

func printFile(filePath, printer string) error {
	switch runtime.GOOS {
	case "windows":
		/*
			p, err := printer.Open(printer)
			if err != nil {
				return err
			}
			defer p.Close()

			return p.PrintFile(filePath)
		*/

	case "darwin", "linux": // macOS and Linux
		cmd := exec.Command("lpr", "-P", printer, filePath)
		return cmd.Run()

	default:
		return fmt.Errorf("unsupported operating system")
	}

	return fmt.Errorf("unsupported operating system")
}
