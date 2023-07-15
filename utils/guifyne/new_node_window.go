package guifyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var (
	new_node_window        fyne.Window
	new_node_window_IsOpen binding.Bool = binding.NewBool()
)
var (
	_new_node_auto             binding.Bool   = binding.NewBool()
	_new_node_auto_count       binding.Int    = binding.NewInt()
	_new_node_auto_countString binding.String = binding.NewSprintf("%d", _new_node_auto_count)
)

func init_new_node_window_toolbar() fyne.CanvasObject {
	_toolbar_auto_nodes_count := NewCustomToolbarEntryItem()
	_new_node_auto_count.Set(1)
	_toolbar_auto_nodes_count.SetBinding(_new_node_auto_countString)
	_toolbar_auto_nodes_count.Filter("^[0-9]+$")
	_toolbar_auto_nodes_count.OnChange(func(str string) {
		_toolbar_auto_nodes_count.Resize(fyne.NewSize(120, 34))
		_toolbar_auto_nodes_count.Get().Validate()
	})
	_toolbar_auto_nodes_count.SetText("1")
	_toolbar_auto_nodes := NewCustomToolbarCheckItem("Auto", func(b bool) {
		if b {

		}
	})
	_toolbar_auto_nodes.SetBinding(_new_node_auto)
	_toolbar := widget.NewToolbar(
		_toolbar_auto_nodes,
		widget.NewToolbarSeparator(),
		NewCustomToolBarLabelItem("<Count>:"),
		_toolbar_auto_nodes_count,
	)
	_new_node_auto_count.Set(12)
	return _toolbar
}

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
		new_node_window.SetContent(
			container.NewVBox(
				init_new_node_window_toolbar(),
			),
		)
		new_node_window.Show()
		new_node_window.SetOnClosed(func() {
			new_node_window_IsOpen.Set(false)
		})
	}
}
