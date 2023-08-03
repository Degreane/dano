package gui

import (
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/degreane/dano/utils/binding"
	"github.com/degreane/dano/utils/gui/widgets"
)

var (
	// defining Operations
	ops op.Ops
	// defining First NavBar Widgets
	newBtnState      widget.Clickable
	loadBtnState     widget.Clickable
	saveBtnState     widget.Clickable
	quitBtnState     widget.Clickable
	spacerFlexState  layout.FlexChild
	tbItemSave       *widgets.ToolBarItem
	tbItemNew        *widgets.ToolBarItem
	tbItemLoad       *widgets.ToolBarItem
	tbItemFlexSpacer *widgets.ToolBarItem
	tbItemQuit       *widgets.ToolBarItem
	// defining Second NavBar Widgets
	sampleName      widget.Label
	sampleInfo      widget.Editor
	samplePrecision widget.Editor
	sampleThreshold widget.Editor
	tbItemName      *widgets.ContainerItem
	tbItemInfo      *widgets.ContainerItem
	tbItemPrecision *widgets.ContainerItem
	tbItemThreshold *widgets.ContainerItem

	// defining bindings
	batchLoaded binding.Boolean = binding.NewBoolean()
	batchSaved  binding.Boolean = binding.NewBoolean()
)

func initFirstToolBarItems(th *material.Theme) {
	spacerFlexState = layout.Flexed(0.2, func(gtx layout.Context) layout.Dimensions {
		return material.Label(th, unit.Sp(10), "").Layout(gtx)
	})
	tbItemNew = widgets.NewToolbarItem[widget.Clickable](&newBtnState, "New")
	tbItemLoad = widgets.NewToolbarItem[widget.Clickable](&loadBtnState, "Load")

	tbItemSave = widgets.NewToolbarItem[widget.Clickable](&saveBtnState, "Save")
	tbItemFlexSpacer = widgets.NewToolbarItem[layout.FlexChild](&spacerFlexState, "")
	tbItemQuit = widgets.NewToolbarItem[widget.Clickable](&quitBtnState, "Quit")
	if batchLoaded.TryGet() {
		tbItemSave.Enable()
	} else {
		tbItemSave.Disable()
	}
	batchLoaded.AddListener(func() {
		// log.Printf("Running Listener")
		if batchLoaded.TryGet() {
			tbItemSave.Enable()
		} else {
			tbItemSave.Disable()
		}
	})
}
func initSecondToolBar(th *material.Theme) {

	// sampleName.Alignment = text.Middle
	// log.Println("Item Is of type <", reflect.TypeOf(sampleName), ">")
	tbItemName = widgets.NewContainerItem[widget.Label](&sampleName, " ", *th)
	tbItemPrecision = widgets.NewContainerItem[widget.Editor](&samplePrecision, "Precision", *th).Disable()
	tbItemThreshold = widgets.NewContainerItem[widget.Editor](&sampleThreshold, "Threshold", *th).Disable()
	tbItemInfo = widgets.NewContainerItem[widget.Editor](&sampleInfo, "Info", *th).Disable().SetWidgetMultiLine(true).SetLabel("Info")
}
func renderNewToolbar(gtx layout.Context, th *material.Theme) layout.Dimensions {
	// tb is the main toolbar.
	tb := widgets.NewToolbar()

	tb.Add(tbItemNew).Add(tbItemLoad).Add(tbItemSave).Add(tbItemFlexSpacer).Add(tbItemQuit)
	// Should Check to see if there is a batch loaded or not
	// if no batch loaded then we should disable tbItemSave
	// if there is a Batch Loaded then we should enable tbItemSave

	if newBtnState.Clicked() {
		tbItemNew.Disable()
		tbItemLoad.Disable()
		newSamplingGUI()
	}
	// tbItemSave.Disable()
	if quitBtnState.Clicked() {
		w.Perform(system.ActionClose)
		// os.Exit(0)
	}
	return tb.Render(gtx, th)
}
func renderSecondToolbar(gtx layout.Context, th *material.Theme) layout.Dimensions {
	tb := widgets.NewContainer[layout.Flex](&layout.Flex{})
	clr := theme.Bg
	clr.A = 128
	tb.Add(tbItemName.SetWidgetTextAlign(text.Middle)).Add(tbItemPrecision.SetWidgetTextAlign(text.Middle)).Add(tbItemThreshold).Add(tbItemInfo).SetMargin(&layout.Inset{
		Top:    0,
		Bottom: 1,
		Left:   1,
		Right:  1,
	}).SetBackgroundColor(clr)
	return tb.Render(gtx)
}
