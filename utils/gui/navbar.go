package gui

import (
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/degreane/dano/utils/binding"
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
	tbItemSave       *ToolBarItem
	tbItemNew        *ToolBarItem
	tbItemLoad       *ToolBarItem
	tbItemFlexSpacer *ToolBarItem
	tbItemQuit       *ToolBarItem
	// defining Second NavBar Widgets

	// defining bindings
	batchLoaded binding.Boolean = binding.NewBoolean()
	batchSaved  binding.Boolean = binding.NewBoolean()
)

func initFirstToolBarItems(th *material.Theme) {
	spacerFlexState = layout.Flexed(0.2, func(gtx layout.Context) layout.Dimensions {
		return material.Label(th, unit.Sp(10), "").Layout(gtx)
	})
	tbItemNew = NewToolbarItem[widget.Clickable](&newBtnState, "New")
	tbItemLoad = NewToolbarItem[widget.Clickable](&loadBtnState, "Load")

	tbItemSave = NewToolbarItem[widget.Clickable](&saveBtnState, "Save")
	tbItemFlexSpacer = NewToolbarItem[layout.FlexChild](&spacerFlexState, "")
	tbItemQuit = NewToolbarItem[widget.Clickable](&quitBtnState, "Quit")
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
func renderNewToolbar(gtx layout.Context, th *material.Theme) layout.Dimensions {
	// tb is the main toolbar.
	tb := NewToolbar()

	tb.Add(tbItemNew).Add(tbItemLoad).Add(tbItemSave).Add(tbItemFlexSpacer).Add(tbItemQuit)
	// Should Check to see if there is a batch loaded or not
	// if no batch loaded then we should disable tbItemSave
	// if there is a Batch Loaded then we should enable tbItemSave

	if newBtnState.Clicked() {
		tbItemNew.Disable()
		newSamplingGUI()
	}
	// tbItemSave.Disable()
	if quitBtnState.Clicked() {
		w.Perform(system.ActionClose)
		// os.Exit(0)
	}
	return tb.render(gtx, th)
}
func renderSecondToolbar(gtx layout.Context, th *material.Theme) layout.Dimensions {
	tb := NewToolbar()

	return tb.render(gtx, th)
}
