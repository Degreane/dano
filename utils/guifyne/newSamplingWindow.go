package guifyne

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var (
	newSamplingWindow fyne.Window
	//(108, 122, 137);
	_nameEntry               = widget.NewEntryWithData(newBatchName)
	_formbgColor color.NRGBA = color.NRGBA{
		R: 108,
		G: 122,
		B: 137,
		A: 32,
	}

	// sets a boolean to prevent multiple opening of the same window
	newSamplingWindowIsOpen binding.Bool   = binding.NewBool()
	newBatchName            binding.String = binding.NewString()
	newBatchInfo            binding.String = binding.NewString()
	// newBatchData            binding.String = binding.NewSprintf("NewBatch: \n(%s)\n", newBatchName)
	newBatchThreshold     binding.Float  = binding.NewFloat()
	newBatchThresholdData binding.String = binding.NewSprintf("%f", newBatchThreshold)
	newBatchPrecision     binding.Int    = binding.NewInt()
	newBatchPrecisionData binding.String = binding.NewSprintf("%d", newBatchPrecision)
)

func _drawTopFrame() fyne.CanvasObject {
	topFrameText := canvas.NewText(" New Batch Creation ", color.NRGBA{R: 250, A: 250})
	topFrameText.Alignment = fyne.TextAlignCenter
	topFrameText.TextSize = 20
	topFrameData := widget.NewLabelWithData(newBatchName)
	topFrameData.TextStyle.Bold = true
	topFrameTextWidget := container.NewHBox(topFrameText, topFrameData)
	sizeConsumed := fyne.NewSize(topFrameTextWidget.MinSize().Width+10, topFrameTextWidget.MinSize().Height+10)
	topFrameRect := canvas.NewRectangle(_formbgColor)
	topFrameRect.StrokeColor = color.NRGBA{B: 250, A: 250}
	topFrameRect.StrokeWidth = 0.5
	topFrameRect.SetMinSize(sizeConsumed)
	return container.NewCenter(container.NewHBox(container.NewPadded(topFrameRect, topFrameTextWidget)))
}

// func _validateNumeric(str string) fyne.StringValidator {
// 	 _rexObj:=regexp.MustCompile(`^[0-9]+\.?[0-9]*$`)
// 	 _rexObj.
// }

func _drawFormEntry() fyne.CanvasObject {
	_infoEntry := widget.NewMultiLineEntry()
	_infoEntry.Bind(newBatchInfo)
	_thresholdEntry := widget.NewEntryWithData(newBatchThresholdData)
	_thresholdEntry.Validator = validation.NewRegexp(`^[0-9]+\.?[0-9]*$`, "Invalid (Threshold) Data")
	_precisionEntry := widget.NewEntryWithData(newBatchPrecisionData)
	_precisionEntry.Validator = validation.NewRegexp(`^[0-9]+$`, "Invalid (Precision) Data")
	_nameEntry.PlaceHolder = "New Batch Name?"
	_nameEntry.Validator = validation.NewRegexp(`^[a-zA-Z0-9\_\-\.]+$`, " Invalid Name ")
	_entryForm := widget.NewForm(
		widget.NewFormItem("Batch Name", _nameEntry),
		widget.NewFormItem("Batch Info", _infoEntry),
		widget.NewFormItem("Batch Threshold", _thresholdEntry),
		widget.NewFormItem("Batch Precision", _precisionEntry),
	)

	_submitEntry := widget.NewFormItem("", widget.NewButtonWithIcon("Submit", theme.DocumentSaveIcon(), func() {
		err := _entryForm.Validate()
		if err != nil {
			dlg := dialog.NewError(err, newSamplingWindow)
			dlg.Show()
		} else {
			_batch_name, _ := newBatchName.Get()
			_batch_thresh, _ := newBatchThreshold.Get()
			_batch_prec, _ := newBatchPrecision.Get()
			_batch_info, _ := newBatchInfo.Get()
			content_toolbar_BatchInfo.Set(_batch_info)
			content_toolbar_BatchName.Set(_batch_name)
			content_toolbar_BatchPrecision.Set(_batch_prec)
			content_toolbar_BatchThreshold.Set(_batch_thresh)
			mainWindowBatchSet.Set(true)

			// enabling high priority marks on gadgets in main window
			topbar_save.SetHighPriority()
			topbar_save.Settext("*")
			//
			newSamplingWindow.Close()
		}
	}))
	_entryForm.AppendItem(_submitEntry)

	_container := container.NewPadded(
		_entryForm,
	)
	_rect := canvas.NewRectangle(_formbgColor)
	_rect.SetMinSize(fyne.NewSize(_container.MinSize().Width+10, _container.MinSize().Height+10))
	_rect.StrokeColor = color.NRGBA{G: 236, B: 236, A: 250}
	_rect.StrokeWidth = 0.5
	_returnedContainer := container.NewPadded(_rect, _container)
	return _returnedContainer
}
func _drawBottomFrame() fyne.CanvasObject {

	return container.NewGridWithColumns(
		3,
		widget.NewLabelWithData(newBatchName),
		widget.NewLabelWithData(newBatchThresholdData),
		widget.NewLabelWithData(newBatchPrecisionData),
	)
}
func newWindowContent() fyne.CanvasObject {

	_frame := container.NewBorder(
		_drawTopFrame(),
		_drawBottomFrame(),
		nil,
		nil,
		_drawFormEntry(),
	)
	return _frame
}
func _setNewSamplingDefaults() {
	newBatchName.Set("")
	newBatchPrecision.Set(2)
	newBatchThreshold.Set(3.0)
}
func openNewDataSamplingWindow() {
	// newSamplingWindow.SetCloseIntercept(func() {

	// })
	ok, _ := newSamplingWindowIsOpen.Get()
	if ok {
		dlg := dialog.NewInformation("Already Open", "New Messaging Window is already open", gui_MainWindow)
		dlg.Show()
		newSamplingWindow.RequestFocus()
	} else {

		newSamplingWindow = gui_Application.NewWindow("New Sampling Window")
		_setNewSamplingDefaults()
		newSamplingWindowIsOpen.Set(true)
		newSamplingWindow.Resize(fyne.NewSize(600, 450))
		newSamplingWindow.SetContent(newWindowContent())
		newSamplingWindow.CenterOnScreen()

		newSamplingWindow.Canvas().Focus(_nameEntry)
		newSamplingWindow.Canvas().Unfocus()
		topbar_new.SetDisable(true)
		newSamplingWindow.Show()
	}
	newSamplingWindow.RequestFocus()
	newSamplingWindow.SetOnClosed(func() {
		topbar_new.SetDisable(false)
		newSamplingWindowIsOpen.Set(false)
	})
}
