package gui

import (
	"image/color"
	"strconv"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

var (
	newSamplingWindowDisplayed bool      = false
	newSamplingWindowDone      chan bool = make(chan bool, 1)
)

func formField(gtx layout.Context, label string, field material.EditorStyle, opts ...map[string]string) layout.Dimensions {
	var lRoF layout.FlexChild
	if len(opts) > 0 {
		if val, ok := opts[0]["flexed"]; ok {
			_val, err := strconv.ParseFloat(val, 32)
			if err != nil {
				_val = 0.8
			}
			lRoF = layout.Flexed(
				float32(_val),
				func(gtx layout.Context) layout.Dimensions {
					margin := layout.Inset{
						Top:    unit.Dp(1),
						Bottom: unit.Dp(1),
						Left:   unit.Dp(15),
						Right:  unit.Dp(15),
					}
					border := widget.Border{
						Color:        color.NRGBA{R: 192, G: 192, B: 192, A: 192},
						CornerRadius: unit.Dp(2),
						Width:        unit.Dp(1),
					}
					return margin.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return border.Layout(
							gtx,
							field.Layout,
						)
					})
				},
			)
		} else {
			lRoF = layout.Rigid(
				func(gtx layout.Context) layout.Dimensions {
					margin := layout.Inset{
						Top:    unit.Dp(1),
						Bottom: unit.Dp(1),
						Left:   unit.Dp(15),
						Right:  unit.Dp(15),
					}
					border := widget.Border{
						Color:        color.NRGBA{R: 192, G: 192, B: 192, A: 192},
						CornerRadius: unit.Dp(2),
						Width:        unit.Dp(1),
					}
					return margin.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return border.Layout(
							gtx,
							field.Layout,
						)
					})
				},
			)
		}
	}
	return layout.Flex{
		Axis:      layout.Horizontal,
		Spacing:   layout.SpaceEnd,
		Alignment: layout.Middle,
		WeightSum: 0,
	}.Layout(
		gtx,
		spaceW(40),
		layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {
				return material.Label(theme, unit.Sp(15), label).Layout(gtx)
			},
		),
		lRoF,
		spaceW(40),
	)
}

func newSamplingGUI() {
	if !newSamplingWindowDisplayed {
		newSamplingWindowDisplayed = !newSamplingWindowDisplayed
		var samplingWindowOps *op.Ops = new(op.Ops)
		samplingWindowOps.Reset()

		go func(newSamplingWindowDone chan bool) {
			newSamplingWindow := app.NewWindow(
				app.Title("New Sample"),
				app.Size(unit.Dp(600), unit.Dp(400)),
			)
			// define the fields we need:
			var (
				// // 1- Sampling Name
				// _newSamplingName *component.TextField = new(component.TextField)
				// // 2- Sampling Info
				// _newSamplingInfo *component.TextField = new(component.TextField)
				// // 3- sampling Threshold
				// _newSamplingThreshold *component.TextField = new(component.TextField)
				// // 4- sampling Precision
				// _newSamplingPrecision *component.TextField = new(component.TextField)
				// // 5- my own implementation
				// _field *Field = new(Field)
				// 6- my own form implementation
				_form *Form = NewForm()
			)
			_form.SetName("New Tempo Form ")
			_form.SetType(FormHorizontal)
			_form.AddField(&Field{
				Label:        "Field 1",
				Type:         InputText,
				DefaultValue: "F1",
				Name:         "Name Of Field 1",
				TextSize:     int(16),
			})
			_form.AddField(&Field{
				Label:        "Field 2",
				Type:         InputText,
				DefaultValue: "F2",
				Name:         "Name Of Field 2",
				TextSize:     int(16),
			})

			// _newSamplingName.Alignment = text.Middle
			// _newSamplingName.Editor.SingleLine = true
			// _newSamplingName.Helper = " Sampling Name "

			// _newSamplingInfo.Editor.SingleLine = false
			// _newSamplingInfo.Helper = " Info "

			// _newSamplingThreshold.Editor.SingleLine = true
			// _newSamplingThreshold.Alignment = text.Middle
			// _newSamplingThreshold.Filter = "0123456789."
			// _newSamplingThreshold.InputHint = key.HintNumeric

			// _newSamplingPrecision.Alignment = text.Middle
			// _newSamplingPrecision.Filter = "0123456789"
			// _newSamplingPrecision.InputHint = key.HintNumeric

			for windowEvent := range newSamplingWindow.Events() {

				switch windowEventType := windowEvent.(type) {
				case system.FrameEvent:
					samplingWindowGtx := layout.NewContext(samplingWindowOps, windowEventType)
					layout.Flex{
						Axis:    layout.Vertical,
						Spacing: layout.SpaceEnd,
					}.Layout(
						samplingWindowGtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return _form.Layout(gtx, theme)
						}),
					)
					windowEventType.Frame(samplingWindowGtx.Ops)
				case system.DestroyEvent:
					newSamplingWindowDone <- false

				}
			}
		}(newSamplingWindowDone)
		newSamplingWindowDisplayed = <-newSamplingWindowDone
	}
}

func newSamplingField(gtx layout.Context) layout.Dimensions {

	// return _field.Layout(gtx, theme)
	return layout.Dimensions{}
}
