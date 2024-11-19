package views

import (
	"image/color"
	"mi-aplicacion/backend/auth"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowLogin(window fyne.Window) {
	// Título
	title := canvas.NewText("Sistema de Gestión", color.NRGBA{R: 0, G: 100, B: 180, A: 255})
	title.TextSize = 24
	title.TextStyle.Bold = true

	// Campos de entrada
	username := widget.NewEntry()
	username.SetPlaceHolder("Usuario")

	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Contraseña")

	// Mensaje de error
	errorMsg := widget.NewLabel("")
	errorMsg.Hide()

	// Botón de login
	loginBtn := widget.NewButton("Iniciar Sesión", nil)
	loginBtn.Importance = widget.HighImportance

	// Función de validación
	validate := func() {
		if username.Text == "" || password.Text == "" {
			errorMsg.SetText("Por favor complete todos los campos")
			errorMsg.Show()
			return
		}

		creds := auth.Credentials{
			Username: username.Text,
			Password: password.Text,
		}

		if auth.ValidateLogin(creds) {
			errorMsg.Hide()
			ShowMainWindow(window)
		} else {
			errorMsg.SetText("Usuario o contraseña incorrectos")
			errorMsg.Show()
		}
	}

	// Asignar la función al botón
	loginBtn.OnTapped = validate

	// También permitir enviar con Enter
	username.OnSubmitted = func(string) { validate() }
	password.OnSubmitted = func(string) { validate() }

	// Crear el formulario
	formContainer := container.NewVBox(
		widget.NewLabel("Usuario:"),
		username,
		widget.NewLabel("Contraseña:"),
		password,
		container.NewCenter(loginBtn),
		container.NewCenter(errorMsg),
	)

	// Contenedor principal con padding y centrado
	content := container.NewCenter(
		container.NewVBox(
			container.NewCenter(title),
			widget.NewSeparator(),
			container.NewPadded(formContainer),
		),
	)

	window.SetContent(content)

	// Dar foco inicial al campo de usuario
	username.FocusGained()
}
