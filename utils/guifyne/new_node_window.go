package guifyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
)

var (
	new_node_window        fyne.Window
	new_node_window_IsOpen binding.Bool = binding.NewBool()
)

func init_new_node_window() {
	if ok, _ := new_node_window_IsOpen.Get(); ok {
		dlg := dialog.NewInformation("Node Window", "New Node Already Open", gui_MainWindow)
		dlg.Show()
		new_node_window.RequestFocus()
	} else {
		new_node_window = gui_Application.NewWindow("New Node")
		new_node_window_IsOpen.Set(true)
		new_node_window.Resize(fyne.NewSize(600, 400))
		new_node_window.CenterOnScreen()
		new_node_window.Show()
		new_node_window.SetOnClosed(func() {
			new_node_window_IsOpen.Set(false)
		})
	}
}
