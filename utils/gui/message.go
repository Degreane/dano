package gui

import (
	"image/color"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

var (
	messageDisplayed chan bool = make(chan bool, 1)
	displayed        bool      = false
)

func drawMessage(title string, message string) {
	if !displayed {
		displayed = !displayed
		// var wg sync.WaitGroup
		var messageOps op.Ops
		messageOps.Reset()
		// wg.Add(1)

		go func(messageDisplayed chan bool, message string) {
			messageWindow := app.NewWindow(
				app.Size(unit.Dp(600), unit.Dp(150)),
				app.Title(title),
			)
			for messageEvent := range messageWindow.Events() {
				switch messageType := messageEvent.(type) {
				case system.FrameEvent:

					messageGtx := layout.NewContext(&messageOps, messageType)
					// var messageWidgetState widget.Label
					messageWidgetStyle := material.Label(theme, unit.Sp(20), message)
					// messageButton := material.Button(theme, &widget.Clickable{}, "ClickMe")
					layout.Flex{
						Axis:      layout.Vertical,
						WeightSum: 0,
						Alignment: layout.Middle,
					}.Layout(
						messageGtx,
						// spaceW(unit.Dp(30)),
						layout.Rigid(
							func(gtx layout.Context) layout.Dimensions {
								p := clip.Path{}
								p.Begin(gtx.Ops)
								p.MoveTo(
									f32.Pt(
										10,
										float32(gtx.Constraints.Min.Y),
									),
								)
								p.LineTo(
									f32.Pt(
										float32(gtx.Constraints.Max.X-10),
										float32(gtx.Constraints.Min.Y),
									),
								)
								p.Close()
								paint.FillShape(gtx.Ops, color.NRGBA{B: 250, A: 250},
									clip.Stroke{
										Path:  p.End(),
										Width: 1,
									}.Op(),
								)
								return layout.Dimensions{
									Size: gtx.Constraints.Min,
								}
							}),
						spaceW(20),
						layout.Rigid(
							func(gtx layout.Context) layout.Dimensions {
								return layout.Flex{
									Axis:      layout.Horizontal,
									Alignment: layout.Start,
								}.Layout(
									gtx,
									spaceW(unit.Dp(20)),

									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return messageWidgetStyle.Layout(gtx)
									}),
								)
							},
						),
						spaceW(10),
						layout.Rigid(
							func(gtx layout.Context) layout.Dimensions {
								p := clip.Path{}
								p.Begin(gtx.Ops)
								p.MoveTo(
									f32.Pt(
										10,
										float32(gtx.Constraints.Min.Y),
									),
								)
								p.LineTo(
									f32.Pt(
										float32(gtx.Constraints.Max.X-10),
										float32(gtx.Constraints.Min.Y),
									),
								)
								p.Close()
								paint.FillShape(gtx.Ops, color.NRGBA{B: 250, A: 250},
									clip.Stroke{
										Path:  p.End(),
										Width: 1,
									}.Op(),
								)
								return layout.Dimensions{
									Size: gtx.Constraints.Max,
								}
							}),
					)
					messageType.Frame(messageGtx.Ops)

				case system.DestroyEvent:
					// wg.Done()
					messageDisplayed <- false
				}
			}
		}(messageDisplayed, message)
		// wg.Wait()
		displayed = <-messageDisplayed
	}

}
