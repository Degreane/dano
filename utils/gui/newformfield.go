package gui

import (
	"fmt"
	"image"
	"image/color"

	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type FieldType int

const (
	InputText FieldType = iota
	InputTextArea
	InputEmail
	InputPhone
	InputInteger
	InputDecimal
	InputSelect
	InputCheckBox
	InputRadioButton
)

type FormType int

const (
	FormNone FormType = iota
	FormHorizontal
	FormInline
	FormVertical
)

type FieldDecoration int

const (
	DecorationLabelFirst FieldDecoration = iota
	DecorationLabelLast
	DecorationLabelInline
	DecorationLabelEmbed
)

type Form struct {
	Fields []*Field
	Name   string
	Type   FormType
}

func NewForm() *Form {
	return new(Form)
}
func (frm *Form) SetName(name string) {
	frm.Name = name
}
func (frm *Form) SetType(tp FormType) {
	frm.Type = tp
}
func (frm *Form) AddField(fld *Field) {
	frm.Fields = append(frm.Fields, fld)
}
func (frm *Form) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	var fFields []layout.FlexChild
	fFields = append(
		fFields,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			border := widget.Border{
				Color:        color.NRGBA{B: 255, A: 255},
				CornerRadius: unit.Dp(10),
				Width:        unit.Dp(1),
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				defer op.Offset(image.Point{
					X: gtx.Constraints.Min.X,
					Y: gtx.Constraints.Min.Y,
				}).Push(gtx.Ops).Pop()
				lbl := material.Label(th, unit.Sp(20), frm.Name).Layout(gtx)

				return lbl
			})
			return border
		}),
	)
	for _, field := range frm.Fields {
		fFields = append(fFields, field.Layout(gtx, th))
	}
	fields := layout.Flex{
		Axis:      layout.Vertical,
		Spacing:   layout.SpaceEnd,
		Alignment: layout.Middle,
	}.Layout(
		gtx,
		fFields...,
	)
	form := layout.Flex{
		Axis:      layout.Vertical,
		Spacing:   layout.SpaceEnd,
		Alignment: layout.Middle,
	}.Layout(
		gtx,
		layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {
				return layout.Stack{}.Layout(
					gtx,
					layout.Expanded(
						func(gtx layout.Context) layout.Dimensions {
							border := widget.Border{
								Color:        color.NRGBA{B: 250, A: 250},
								CornerRadius: unit.Dp(2),
								Width:        1,
							}.Layout(
								gtx,
								func(gtx layout.Context) layout.Dimensions {
									return fields
								},
							)
							return border
						},
					),
				)
			},
		),
	)
	return form

}

type Field struct {
	Label        string
	Type         FieldType
	DefaultValue string
	Name         string
	Validate     bool
	Hint         string
	Icon         string
	TextSize     int
	Decoration   FieldDecoration
	Value        interface{}
	Editor       widget.Editor
}

func gap(size int, gtx layout.Context) layout.Dimensions {
	return layout.Spacer{Width: unit.Dp(size), Height: unit.Dp(size)}.Layout(gtx)
}
func (f *Field) Layout(gtx layout.Context, theme *material.Theme) layout.FlexChild {
	// In The System we have mostly a Label and the Gadget itself
	// create record_border which borders the whole set
	record_border := widget.Border{
		Color: color.NRGBA{
			R: 192,
			G: 250,
			B: 192,
			A: 250,
		},
		CornerRadius: unit.Dp(5),
		Width:        unit.Dp(1),
	}
	// create record_margin to be placed inside the record_border, wrapping the whole set
	record_margin := layout.Inset{
		Top:    unit.Dp(5),
		Bottom: unit.Dp(5),
		Left:   unit.Dp(10),
		Right:  unit.Dp(5),
	}
	record_flex := layout.Flex{
		Axis:      layout.Vertical,
		Spacing:   layout.SpaceEnd,
		Alignment: layout.Middle,
		WeightSum: 100,
	}
	// create the label
	label := material.Label(theme, unit.Sp(f.TextSize), fmt.Sprintf("%s :", f.Label)) // Label definition
	// Widget Itself
	// 1- Check Widget Type
	if f.Type >= InputText && f.Type <= InputDecimal {

		if f.Type == InputText {
			f.Editor.SingleLine = true
		} else if f.Type == InputTextArea {
			f.Editor.SingleLine = false
		} else if f.Type == InputInteger {
			f.Editor.InputHint = key.HintNumeric
			f.Editor.Filter = "0123456789"
		} else if f.Type == InputDecimal {
			f.Editor.InputHint = key.HintNumeric
			f.Editor.Filter = "0123456789."
		} else if f.Type == InputEmail {
			f.Editor.InputHint = key.HintEmail
			f.Editor.Filter = "0123456789abcdefghijklmnopqrstuvwxyz._-@"
		} else {
			f.Editor.InputHint = key.HintTelephone
		}
		// inputWidgetStyle := material.Editor(theme, &inputWidget, f.DefaultValue)
		return layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {
				return record_flex.Layout(
					gtx,
					layout.Rigid(
						func(gtx layout.Context) layout.Dimensions {
							return gap(10, gtx)
						},
					),
					layout.Rigid(
						func(gtx layout.Context) layout.Dimensions {
							return record_border.Layout(
								gtx,
								func(gtx layout.Context) layout.Dimensions {
									return record_margin.Layout(
										gtx,
										func(gtx layout.Context) layout.Dimensions {
											return layout.Flex{
												Axis:      layout.Horizontal,
												Spacing:   layout.SpaceEnd,
												Alignment: layout.Middle,
											}.Layout(
												gtx,
												layout.Rigid(label.Layout),
												layout.Rigid(
													func(gtx layout.Context) layout.Dimensions {
														return gap(20, gtx)
													},
												),
												layout.Rigid(material.Editor(theme, &f.Editor, f.DefaultValue).Layout),
											)
										},
									)
								},
							)
						},
					),
				)
			},
		)

	}
	return layout.Rigid(
		func(gtx layout.Context) layout.Dimensions {
			return layout.Dimensions{
				Size: gtx.Constraints.Min,
			}
		},
	)

}
