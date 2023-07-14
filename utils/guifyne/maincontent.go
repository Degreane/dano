package guifyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// these variables used to define control of the batch if set or not
var (
	_mainWindowBatchSet bool = false
	mainWindowBatchSet       = binding.BindBool(&_mainWindowBatchSet)
)

// border layout containers
var (
	top    fyne.CanvasObject
	bottom fyne.CanvasObject
	left   fyne.CanvasObject
	right  fyne.CanvasObject
	// content is the container that holds the data inside
	content *fyne.Container = container.NewVBox()
)

// Content of main window should display the following:
// batch name
// batch threshold
// batch precision
// they all are bound to the actual data
// and format binding for display
var (
	// batch name
	content_toolbar_BatchName       binding.String = binding.NewString()
	content_toolbar_BatchNameString binding.String = binding.NewSprintf("Batch (Name: <%s>)", content_toolbar_BatchName)
	// threshold
	content_toolbar_BatchThreshold       binding.Float  = binding.NewFloat()
	content_toolbar_BatchThresholdString binding.String = binding.NewSprintf("(Threshold: <%f>)", content_toolbar_BatchThreshold)
	// Precision
	content_toolbar_BatchPrecision       binding.Int    = binding.NewInt()
	content_toolbar_BatchPrecisionString binding.String = binding.NewSprintf("(Precision: <%d>)", content_toolbar_BatchPrecision)
	// Info
	content_toolbar_BatchInfo binding.String = binding.NewString()

	// No Of Nodes
	content_toolbar_no_of_nodes        binding.Int    = binding.NewInt()
	content_toolbar_no_of_nodes_string binding.String = binding.NewSprintf("No. Of Nodes <%d>", content_toolbar_no_of_nodes)

	// Add New Node Button toolbar
	content_toolbar_add_nodes PsyButtonToolbar

	first_content_toolbar  fyne.CanvasObject
	second_content_toolbar fyne.CanvasObject
)

// top bar toolbar icons and actions definition
var (
	topbar_save PsyButtonToolbar
	topbar_new  PsyButtonToolbar
	topbar_load PsyButtonToolbar
)

// center first_content toolbar
func initFirstContent_toolbar() fyne.CanvasObject {
	_bname := NewCustomToolBarLabelItem("")
	_bname.SetBinding(content_toolbar_BatchNameString)
	_bthresh := NewCustomToolBarLabelItem("")
	_bthresh.SetBinding(content_toolbar_BatchThresholdString)
	_bprecision := NewCustomToolBarLabelItem("")
	_bprecision.SetBinding(content_toolbar_BatchPrecisionString)
	first_content_toolbar = widget.NewToolbar(
		_bname,
		widget.NewToolbarSpacer(),
		_bprecision, widget.NewToolbarSpacer(),
		_bthresh,
	)
	return first_content_toolbar
}

// center second_content_toolbar
func initSecondContent_toolbar() fyne.CanvasObject {
	_no_of_nodes := NewCustomToolBarLabelItem("")
	_no_of_nodes.SetBinding(content_toolbar_no_of_nodes_string)

	content_toolbar_add_nodes = NewCustomToolbarButtonItem("Add", theme.ContentAddIcon(), func() {
		init_new_node_window()
	})
	content_toolbar_add_nodes.SetDisable(true)
	content_toolbar_add_nodes.SetNormalPriority()
	second_content_toolbar = widget.NewToolbar(
		_no_of_nodes,
		widget.NewToolbarSeparator(),
		content_toolbar_add_nodes,
		widget.NewToolbarSpacer(),
	)
	return second_content_toolbar
}

func initContents() {
	setMainWindowContent()
	set_TopToolbar()

}

// mainContent should be of type border{top,bottom,left,right,center}
func mainContent() fyne.CanvasObject {
	initContents()
	return container.NewBorder(top, bottom, left, right, content)
}

// Sets the main window Content
func setMainWindowContent() fyne.CanvasObject {
	initFirstContent_toolbar()
	initSecondContent_toolbar()
	content = container.NewVBox(
		first_content_toolbar,
		second_content_toolbar,
		container.NewPadded(widget.NewLabel("Some Values Left Inside")),
	)
	return content

}

func set_TopToolbar() fyne.CanvasObject {
	// define save top bar toolbar
	topbar_save = NewCustomToolbarButtonItem("", theme.DocumentSaveIcon(), nil)
	topbar_save.SetNormalPriority()
	topbar_save.SetDisable(true)
	// define new top bar toolbar
	topbar_new = NewCustomToolbarButtonItem("", theme.FolderNewIcon(), func() {
		topbar_save.SetNormalPriority()
		openNewDataSamplingWindow()
	})
	topbar_new.SetNormalPriority()

	// define load top bar toolbar
	topbar_load = NewCustomToolbarButtonItem("", theme.FolderOpenIcon(), nil)
	topbar_load.SetNormalPriority()
	// define top bar toolbar
	topToolbar := widget.NewToolbar(
		topbar_new,
		topbar_load,
		topbar_save,
		widget.NewToolbarSeparator(),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.LogoutIcon(), func() {
			gui_Application.Quit()
		}),
	)
	top = topToolbar
	return top
}
