package gui

import (
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func renderMainWindow() error {

	// startButton is a clickable widget
	var startButton widget.Clickable
	// th defines the material design style
	th := theme
	var btn material.ButtonStyle = material.Button(th, &startButton, "ClickMe")

	for e := range w.Events() {
		switch et := e.(type) {
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, et)
			layout.Flex{
				Axis:      layout.Vertical,
				Spacing:   layout.SpaceEnd,
				Alignment: layout.Start,
			}.Layout(gtx,
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return renderNewToolbar(gtx, th)
					},
				),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return renderSecondToolbar(gtx, th)
					},
				),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return layout.Dimensions{
						Size: gtx.Constraints.Min,
					}
				}),
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						margins := layout.Inset{
							Top:    unit.Dp(1),
							Bottom: unit.Dp(1),
							Left:   unit.Dp(20),
							Right:  unit.Dp(20),
						}
						return margins.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return btn.Layout(gtx)
						})
					},
				),
				spaceW(10),
			)

			et.Frame(gtx.Ops)
		case system.DestroyEvent:
			return et.Err

		}

	}
	return nil
}
