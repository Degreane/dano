package gui

import (
	"image"
	"image/color"
	"math/rand"
	"time"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/degreane/dano/utils/names"
)

type FormItemDecoration struct {
	NormalBorder widget.Border
	ErrorBorder  widget.Border
}

type FormItemType interface {
	*widget.Editor | *widget.Clickable | *widget.Bool | *widget.Label | *widget.Float | material.LabelStyle
}
type FormItem struct {
	Icon       widget.Icon
	Hint       string
	Validation func() error
	Id         string
	Type       interface{}
	Text       string
	Disabled   bool
	Bound      func()
	bound      interface{}
	boundValue interface{}
}

func NewFormItem[T FormItemType](w *T, txt string) *FormItem {
	Id := rand.Int63n(time.Now().UnixNano())
	IdStr := names.GetName(int(Id), "")
	return &FormItem{
		Id:       IdStr,
		Type:     w,
		Text:     txt,
		Disabled: false,
	}
}

func (f *FormItem) Disable() *FormItem {
	f.Disabled = true
	return f
}
func (f *FormItem) Enabled() *FormItem {
	f.Disabled = false
	return f
}
func (f *FormItem) ToggleState() *FormItem {
	f.Disabled = !f.Disabled
	return f
}

// Form
type FormAxisType int

const (
	Vertical FormAxisType = iota
	Horizontal
	Inline
)

// The Form :
//
// should hold a border
//
// should hold a Name
//
// should hold a caption
//
// should hold an Id
//
// an array of Fields
//
// Axis Type (Horizontal,Vertical)
type Form struct {
	Border   *widget.Border // border around the form
	Name     string         // name of the form
	Caption  string         // title displayed
	Id       string         // unique id of the form
	Fields   []*FormItem    //form fields
	Axis     FormAxisType
	Disabled bool
}

func NewForm(name string) *Form {
	idInt := rand.Int63n(time.Now().UnixNano())
	id := names.GetName(int(idInt), "")
	frm := &Form{
		Border: &widget.Border{
			Color: color.NRGBA{
				R: 236,
				G: 236,
				B: 236,
				A: 128,
			},
			CornerRadius: unit.Dp(5),
			Width:        unit.Dp(1),
		},
		Id:   id,
		Name: name,
		Axis: Horizontal,
	}
	return frm
}
func (f *Form) render(gtx layout.Context) layout.Dimensions {
	frmLayout := layout.Flex{
		Axis: func() layout.Axis {
			if f.Axis == Inline {
				return layout.Horizontal
			} else {
				return layout.Vertical
			}
		}(),
	}
	return frmLayout.Layout(
		gtx,
		layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {
				return layout.Stack{
					Alignment: layout.Center,
				}.Layout(
					gtx,
					layout.Expanded(func(gtx layout.Context) layout.Dimensions {
						ops := gtx.Ops
						rrect := clip.RRect{
							Rect: image.Rect(5, 1, gtx.Constraints.Max.X-5, gtx.Constraints.Min.Y-1),
							SE:   int(f.Border.CornerRadius),
							SW:   int(f.Border.CornerRadius),
							NW:   int(f.Border.CornerRadius),
							NE:   int(f.Border.CornerRadius),
						}
						defer rrect.Push(ops).Pop()
						bgColor := theme.ContrastBg
						bgColor.A = 128
						cops := paint.ColorOp{
							Color: bgColor,
						}
						cops.Add(ops)
						pops := paint.PaintOp{}
						pops.Add(ops)
						paint.FillShape(
							ops,
							f.Border.Color,
							clip.Stroke{
								Path:  rrect.Path(ops),
								Width: 1,
							}.Op(),
						)
						return layout.Dimensions{
							Size:     gtx.Constraints.Min,
							Baseline: 12,
						}

					}),
					layout.Stacked(
						func(gtx layout.Context) layout.Dimensions {
							return material.Label(theme, unit.Sp(12), "Label In Form").Layout(gtx)
						},
					),
				)
			},
		),
	)
}
