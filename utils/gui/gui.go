package gui

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/unit"
	"gioui.org/widget/material"

)

var (
	w     *app.Window
	theme *material.Theme = material.NewTheme(gofont.Collection())
)

func InitGUI() {
	go func() {
		w = app.NewWindow(
			app.Title("Dano Master Thesis (2023)"),
			// app.Size(unit.Dp(1000), unit.Dp(600)),
			// app.Maximized.Option(),
			// app.Fullscreen.Option(),
		)

		theme.TextSize = unit.Sp(10)
		initFirstToolBarItems(theme)
		initSecondToolBar(theme)
		initThirdToolBar(theme)
		if err := renderMainWindow(); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

}
func RunGUI() {
	app.Main()
}
