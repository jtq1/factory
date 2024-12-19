package main

import (
	"appTalleres/backend/db"
	"appTalleres/frontend/views"
	"appTalleres/frontend/views/helper"
	"appTalleres/global"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func main() {
	window := setConfiguration()
	defer global.GetDB().Close()

	views.ShowLogin(window)

	window.ShowAndRun()
}

func setConfiguration() fyne.Window {
	app := app.New()
	window := app.NewWindow("Aragón Gestión")

	app.Settings().SetTheme(theme.LightTheme())

	width, height, err := getMonitorResolution()
	if err != nil {
		helper.CreateCriticalPopUp(window, "error getMonitorResolution", err)
	}
	window.Resize(fyne.NewSize(float32(width), float32(height)))
	window.CenterOnScreen()
	//window.SetFullScreen(true)

	con, err := db.ConnectDB()
	if err != nil || con == nil {
		helper.CreateCriticalPopUp(window, "error trying to connect the DB", err)
	}
	global.SetGlobalDB(con)

	return window
}

func getMonitorResolution() (int, int, error) {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	// Get the primary monitor
	monitor := glfw.GetPrimaryMonitor()
	if monitor == nil {
		return 0, 0, fmt.Errorf("no monitor found")
	}

	videoMode := monitor.GetVideoMode()

	// Print screen width and height
	fmt.Printf("Screen Width: %d px, Screen Height: %d px\n", videoMode.Width, videoMode.Height)
	return videoMode.Width, videoMode.Height, nil
}

func initializeMasters() {

}
