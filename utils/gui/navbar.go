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

// func renderTopBar(et system.FrameEvent, gtx layout.Context) layout.FlexChild {
// 	_color := theme.Palette.ContrastBg
// 	_color.A = 128
// 	r2 := layout.Rigid(
// 		func(gtx layout.Context) layout.Dimensions {

// 			dims := widget.Border{
// 				Width:        unit.Dp(1),
// 				Color:        _color,
// 				CornerRadius: unit.Dp(5),
// 			}.Layout(
// 				gtx,
// 				func(gtx layout.Context) layout.Dimensions {

// 					newBtnStyle = material.Button(theme, &newBtnState, "New")
// 					loadBtnStyle = material.Button(theme, &loadBtnState, "Load")
// 					saveBtnStyle = material.Button(theme, &saveBtnState, "Save")
// 					quitBtnStyle = material.Button(theme, &quitBtnState, "Quit")
// 					return layout.UniformInset(unit.Dp(10)).Layout(
// 						gtx,
// 						func(gtx layout.Context) layout.Dimensions {
// 							return layout.Flex{
// 								Axis:      layout.Horizontal,
// 								Spacing:   layout.SpaceEnd,
// 								Alignment: layout.Start,
// 							}.Layout(gtx,
// 								spaceW(unit.Dp(20)),
// 								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
// 									return newBtnStyle.Layout(gtx)
// 								}),
// 								spaceW(unit.Dp(5)),
// 								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
// 									return loadBtnStyle.Layout(gtx)
// 								}),
// 								spaceW(unit.Dp(5)),
// 								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
// 									return saveBtnStyle.Layout(gtx)
// 								}),
// 								layout.Flexed(1.9, func(gtx layout.Context) layout.Dimensions {
// 									return material.Label(theme, unit.Sp(20), "").Layout(gtx)
// 								}),
// 								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
// 									var th *material.Theme = material.NewTheme(gofont.Regular())
// 									th.Bg = color.NRGBA{R: 200, G: 200}
// 									return quitBtnStyle.Layout(gtx.Disabled())
// 								}),
// 								spaceW(unit.Dp(20)),
// 							)
// 						},
// 					)

// 				},
// 			)
// 			_rect := clip.UniformRRect(image.Rect(0, 0, dims.Size.X, dims.Size.Y), 2).Op(gtx.Ops)

//				paint.FillShape(gtx.Ops, _color, _rect)
//				return dims
//			},
//		)
//		if newBtnState.Clicked() {
//			// nWindow := app.NewWindow(
//			// 	app.Size(unit.Dp(600), unit.Dp(600)),
//			// )
//			// go func() {
//			// 	for f := range nWindow.Events() {
//			// 		switch ft := f.(type) {
//			// 		case system.FrameEvent:
//			// 			fmt.Println("FrameEvent SubWindow ")
//			// 		case system.DestroyEvent:
//			// 			fmt.Println(ft.Err)
//			// 		}
//			// 	}
//			// }()
//			//go drawMessage("new Sampling", "Adding New SamplingOf a very very \n Long Long file and text assumed inside \n Good Luck")
//			go newSamplingGUI()
//		}
//		return r2
//	}
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
