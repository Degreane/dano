package gui

import (
	"image"
	"image/color"
	"log"

	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

var (
	newBtnState  widget.Clickable
	newBtnStyle  material.ButtonStyle
	loadBtnState widget.Clickable
	loadBtnStyle material.ButtonStyle
	saveBtnState widget.Clickable
	saveBtnStyle material.ButtonStyle
	quitBtnState widget.Clickable
	quitBtnStyle material.ButtonStyle
	chkbox       widget.Bool
	// ops are the operations from the UI
	ops op.Ops
)

func renderTopBar(et system.FrameEvent, gtx layout.Context) layout.FlexChild {
	_color := theme.Palette.ContrastBg
	_color.A = 128
	r2 := layout.Rigid(
		func(gtx layout.Context) layout.Dimensions {

			dims := widget.Border{
				Width:        unit.Dp(1),
				Color:        _color,
				CornerRadius: unit.Dp(5),
			}.Layout(
				gtx,
				func(gtx layout.Context) layout.Dimensions {

					newBtnStyle = material.Button(theme, &newBtnState, "New")
					loadBtnStyle = material.Button(theme, &loadBtnState, "Load")
					saveBtnStyle = material.Button(theme, &saveBtnState, "Save")
					quitBtnStyle = material.Button(theme, &quitBtnState, "Quit")
					return layout.UniformInset(unit.Dp(10)).Layout(
						gtx,
						func(gtx layout.Context) layout.Dimensions {
							return layout.Flex{
								Axis:      layout.Horizontal,
								Spacing:   layout.SpaceEnd,
								Alignment: layout.Start,
							}.Layout(gtx,
								spaceW(unit.Dp(20)),
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return newBtnStyle.Layout(gtx)
								}),
								spaceW(unit.Dp(5)),
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return loadBtnStyle.Layout(gtx)
								}),
								spaceW(unit.Dp(5)),
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return saveBtnStyle.Layout(gtx)
								}),
								layout.Flexed(1.9, func(gtx layout.Context) layout.Dimensions {
									return material.Label(theme, unit.Sp(20), "").Layout(gtx)
								}),
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									var th *material.Theme = material.NewTheme(gofont.Regular())
									th.Bg = color.NRGBA{R: 200, G: 200}
									return quitBtnStyle.Layout(gtx.Disabled())
								}),
								spaceW(unit.Dp(20)),
							)
						},
					)

				},
			)
			_rect := clip.UniformRRect(image.Rect(0, 0, dims.Size.X, dims.Size.Y), 2).Op(gtx.Ops)

			paint.FillShape(gtx.Ops, _color, _rect)
			return dims
		},
	)
	if newBtnState.Clicked() {
		// nWindow := app.NewWindow(
		// 	app.Size(unit.Dp(600), unit.Dp(600)),
		// )
		// go func() {
		// 	for f := range nWindow.Events() {
		// 		switch ft := f.(type) {
		// 		case system.FrameEvent:
		// 			fmt.Println("FrameEvent SubWindow ")
		// 		case system.DestroyEvent:
		// 			fmt.Println(ft.Err)
		// 		}
		// 	}
		// }()
		//go drawMessage("new Sampling", "Adding New SamplingOf a very very \n Long Long file and text assumed inside \n Good Luck")
		go newSamplingGUI()
	}
	return r2
}
func renderNewToolbar(gtx layout.Context, th *material.Theme) layout.Dimensions {
	tb := NewToolbar()
	tbItem1 := NewToolbarItem[widget.Clickable](&newBtnState, "New")
	tbItem2 := NewToolbarItem[widget.Clickable](&loadBtnState, "Load")
	tbItem3 := NewToolbarItem[widget.Clickable](&saveBtnState, "Save")
	tbItem4 := NewToolbarItem[widget.Label](&widget.Label{}, "label Widget")

	chkboxS := material.CheckBox(th, &chkbox, "Auto")

	tbItem5 := NewToolbarItem[material.CheckBoxStyle](&chkboxS, "Auto")
	tb.Add(tbItem1).Add(tbItem2).Add(tbItem3).Add(tbItem4).Add(tbItem5)
	if saveBtnState.Clicked() {
		log.Println("Save Clicked")
	}
	if chkbox.Changed() {
		log.Println("Log Changed")
	}
	return tb.render(gtx, th)
}
