package gui

import (
	"image"
	"image/color"
	"reflect"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type ToolbarItemType interface {
	widget.Label | *widget.Label | widget.Clickable | *widget.Clickable | widget.Editor | *widget.Editor | material.CheckBoxStyle
}

type ToolBarItem struct {
	Type     interface{}
	Text     string
	Disabled bool
}

func NewToolbarItem[T ToolbarItemType](w *T, txt string) *ToolBarItem {

	return &ToolBarItem{
		Type:     w,
		Text:     txt,
		Disabled: true,
	}
}

type ToolBar struct {
	Border widget.Border
	Items  []*ToolBarItem
}

// in a new toolbar we need to have
// toolbar items
func (t *ToolBar) Add(item *ToolBarItem) *ToolBar {
	t.Items = append(t.Items, item)
	return t
}
func NewToolbar() *ToolBar {
	return &ToolBar{
		Border: widget.Border{
			Color: color.NRGBA{
				R: 250,
				G: 0,
				B: 0,
				A: 250,
			},
			CornerRadius: unit.Dp(5),
			Width:        unit.Dp(1),
		},
	}
}

// we need to draw the newToolbar and get its Dimensions
func (tb *ToolBar) render(gtx layout.Context, th *material.Theme) layout.Dimensions {
	l := layout.Flex{}
	l.Axis = layout.Horizontal
	l.Alignment = layout.Middle

	st := layout.Stack{}.Layout(
		gtx,
		layout.Expanded(
			func(gtx layout.Context) layout.Dimensions {

				ops := gtx.Ops
				// log.Println(gtx.Constraints)

				rrect := clip.RRect{
					Rect: image.Rect(5, 1, gtx.Constraints.Max.X-5, gtx.Constraints.Min.Y-1),
					SE:   int(tb.Border.CornerRadius),
					SW:   int(tb.Border.CornerRadius),
					NW:   int(tb.Border.CornerRadius),
					NE:   int(tb.Border.CornerRadius),
				}
				defer rrect.Push(ops).Pop()
				bgColor := th.ContrastBg
				bgColor.A = 128
				cops := paint.ColorOp{
					Color: bgColor,
				}
				cops.Add(ops)
				pops := paint.PaintOp{}
				pops.Add(ops)
				paint.FillShape(
					ops,
					tb.Border.Color,
					clip.Stroke{
						Path:  rrect.Path(ops),
						Width: 1,
					}.Op(),
				)
				return layout.Dimensions{
					Size:     gtx.Constraints.Min,
					Baseline: 12,
				}

			},
		),
		layout.Stacked(
			func(gtx layout.Context) layout.Dimensions {
				return layout.Inset{
					Top:    unit.Dp(1),
					Bottom: unit.Dp(1),
					Left:   unit.Dp(1),
					Right:  unit.Dp(1),
				}.Layout(
					gtx,
					func(gtx layout.Context) layout.Dimensions {
						var flexitems []layout.FlexChild
						for i, tbi := range tb.Items {
							// log.Println("<", tbi.Type, ">=><", tbi.Text, ">")
							switch reflect.TypeOf(tbi.Type) {
							case reflect.TypeOf(&widget.Clickable{}):
								disabled := func(gtx layout.Context) layout.Context {
									if tb.Items[i].Disabled {
										return gtx.Disabled()
									} else {
										return gtx
									}
								}(gtx)
								v, _ := tbi.Type.(*widget.Clickable)
								// v.Layout(disabled,func(gtx layout.Context) layout.Dimensions {
								// 	return
								// })
								mat := material.Button(th, v, tb.Items[i].Text)

								//log.Println("<", tb.Items[i].Text, ">")
								flexitems = append(flexitems, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return layout.Inset{Top: 1, Bottom: 1, Left: 1, Right: 1}.Layout(disabled, mat.Layout)
								}))
							case reflect.TypeOf(&material.CheckBoxStyle{}):
								v, _ := tbi.Type.(*material.CheckBoxStyle)
								flexitems = append(flexitems, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return layout.Inset{Top: 1, Bottom: 1, Left: 1, Right: 1}.Layout(gtx, v.Layout)
								}))
							default:
								flexitems = append(flexitems, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return material.Label(th, unit.Sp(20), "Label Here").Layout(gtx)
								}))
							}
						}
						return l.Layout(
							gtx,
							flexitems...,
						)
					},
				)

			},
		),
	)
	return st
}
