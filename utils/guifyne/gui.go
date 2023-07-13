package guifyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

// Define Variables across the system
var (
	// gui main application that should run on the main routine
	gui_Application fyne.App = app.New()
	gui_MainWindow           = gui_Application.NewWindow("Dano Thesis")
)

func RunGUI() {
	// gui_Application.Settings().SetTheme(theme.LightTheme())
	gui_MainWindow.SetContent(mainContent())
	gui_MainWindow.Resize(fyne.NewSize(800, 600))
	gui_MainWindow.CenterOnScreen()
	// gui_MainWindow.SetCloseIntercept(func() {

	// })
	gui_MainWindow.SetOnClosed(func() {
		gui_Application.Quit()
	})
	gui_MainWindow.Show()
	gui_Application.Run()

}
