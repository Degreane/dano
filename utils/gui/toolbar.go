package gui

import (
	"image"
	"image/color"
	"log"
	"math/rand"
	"reflect"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/degreane/dano/utils/binding"
	"github.com/degreane/dano/utils/names"
)

type ToolbarItemType interface {
	widget.Label | *widget.Label | widget.Clickable | *widget.Clickable | widget.Editor | *widget.Editor | material.CheckBoxStyle
}

type ToolBarItem struct {
	Id       string
	Type     interface{}
	Text     string
	Disabled bool
	Bound    binding.Boundable
}

func NewToolbarItem[T ToolbarItemType](w *T, txt string) *ToolBarItem {
	Id := rand.Intn(100000000)
	IdStr := names.GetName(Id, "")
	return &ToolBarItem{
		Id:       IdStr,
		Type:     w,
		Text:     txt,
		Disabled: false,
	}
}

func (t *ToolBarItem) Disable() *ToolBarItem {
	t.Disabled = true
	return t
}
func (t *ToolBarItem) Enable() *ToolBarItem {
	t.Disabled = false
	return t
}
func (t *ToolBarItem) ToggleState() *ToolBarItem {
	t.Disabled = !t.Disabled
	return t
}
func (t *ToolBarItem) Bind(b binding.String) {
	t.Bound = binding.NewBound(t, func() {
		log.Printf("A New Boundable of Widget %+v bound to value %s\n", t.Text, b.TryGet())
	})
	b.Register(t)
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
							case reflect.TypeOf(&widget.Label{}):
								// v, _ := tbi.Type.(*widget.Label)
								mat := material.Label(th, unit.Sp(12), tb.Items[i].Text)
								flexitems = append(flexitems, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return layout.Inset{Top: 1, Bottom: 1, Left: 1, Right: 1}.Layout(gtx, mat.Layout)
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
