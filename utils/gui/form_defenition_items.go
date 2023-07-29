package gui

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math/rand"
	"reflect"
	"time"

	"gioui.org/font"
	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/degreane/dano/utils/names"
)

var (
	sampleClickable widget.Clickable = widget.Clickable{}
	RED             color.NRGBA      = color.NRGBA{
		R: 236,
		A: 236,
	}
	DisabledColor color.NRGBA = color.NRGBA{
		R: 192,
		G: 192,
		B: 192,
		A: 236,
	}
)

type FormItemDecoration struct {
	NormalBorder widget.Border
	ErrorBorder  widget.Border
}

type FormItemType interface {
	*widget.Editor | widget.Editor | *widget.Clickable | *widget.Bool | *widget.Label | *widget.Float | material.LabelStyle
}
type FormItem struct {
	Icon           widget.Icon
	Hint           string
	Validation     func() error
	Id             string
	Type           interface{}
	LabelText      string
	LabelTextWidth int
	disabled       bool
	singleLine     bool
	Bound          func()
	bound          interface{}
	boundValue     interface{}
	Decoration     *formDecoration
}

func NewFormItem[T FormItemType](w *T, txt string) *FormItem {
	Id := rand.Int63n(time.Now().UnixNano())
	IdStr := names.GetName(int(Id), "")
	decoration := new(formDecoration)
	decoration.Border = &widget.Border{
		Color:        color.NRGBA{R: 128, G: 128, B: 128, A: 236},
		CornerRadius: unit.Dp(2),
		Width:        unit.Dp(2),
	}
	decoration.Background = color.NRGBA{
		R: 32,
		G: 32,
		B: 32,
		A: 236,
	}
	decoration.Pad = &layout.Inset{
		Top:    unit.Dp(5),
		Bottom: unit.Dp(5),
		Left:   unit.Dp(5),
		Right:  unit.Dp(5),
	}
	// log.Println("Length of the label Text is <", len(fmt.Sprintf("[%[2]*[1]s]", txt, 16)), fmt.Sprintf("[%[2]*[1]s]", txt, 16), ">")
	return &FormItem{
		Id:             IdStr,
		Type:           w,
		LabelText:      fmt.Sprintf("%-[1]*[2]s", 10-len(txt), txt),
		LabelTextWidth: 10,
		disabled:       false,
		Decoration:     decoration,
		singleLine:     true,
	}
}

// Set Multiline for Editor Widgets
func (f *FormItem) SetMultiLine(multi bool) *FormItem {
	switch reflect.TypeOf(f.Type) {
	case reflect.TypeOf(&widget.Editor{}):
		f.singleLine = !multi
		f.Type.(*widget.Editor).SingleLine = multi
		// log.Println("Set Multiline <", multi, ">")
	}
	return f
}

func (f *FormItem) SetFilter(fltr string) *FormItem {
	switch reflect.TypeOf(f.Type) {
	case reflect.TypeOf(&widget.Editor{}):
		f.Type.(*widget.Editor).Filter = fltr
	}
	return f
}
func (f *FormItem) Disable() *FormItem {
	f.disabled = true
	switch reflect.TypeOf(f.Type) {
	case reflect.TypeOf(&widget.Editor{}):
		f.Type.(*widget.Editor).ReadOnly = true
	}
	return f
}
func (f *FormItem) Enabled() *FormItem {
	f.disabled = false
	switch reflect.TypeOf(f.Type) {
	case reflect.TypeOf(&widget.Editor{}):
		f.Type.(*widget.Editor).ReadOnly = false
	}
	return f
}
func (f *FormItem) ToggleState() *FormItem {
	f.disabled = !f.disabled
	switch reflect.TypeOf(f.Type) {
	case reflect.TypeOf(&widget.Editor{}):
		f.Type.(*widget.Editor).ReadOnly = f.disabled
	}
	return f

}

func (f *FormItem) SetText(txt string) *FormItem {
	f.LabelText = fmt.Sprintf("%[2]*[1]s", txt, f.LabelTextWidth-len(txt))
	return f
}

func (f *FormItem) SetHint(txt string) *FormItem {
	f.Hint = txt
	return f
}

func (f *FormItem) Render(gtx layout.Context) layout.FlexChild {

	// Refactoring the FormItem
	// instead of using stack we shall use Flex
	var ed layout.FlexChild
	switch reflect.TypeOf(f.Type) {
	case reflect.TypeOf(&widget.Editor{}):
		_lbl := material.Label(theme, unit.Sp(15), fmt.Sprintf("%s:", f.LabelText)) // label to be displayed
		_lbl.Font.Variant = font.Variant("Mono")                                    // font using Mono for fixed dimensions
		_editor := material.Editor(theme, f.Type.(*widget.Editor), f.Hint)          // hint for the Editor
		_editor.Editor.Alignment = text.Middle                                      // Text Alignment (ToDo should be constructed via method call )
		_editor.Editor.InputHint = key.HintAny                                      // Hint for Keyboard (ToDo should ve done via method call)
		_editor.HintColor = color.NRGBA{B: 236, A: 236}                             // placeholder color (Blue in our case)
		_editor.Font.Style = font.Italic                                            // font Italic
		_editor.Font.Variant = font.Variant("Smallcaps")                            // small caps for the font
		// _editor.Editor.SingleLine = true                                              // (Single/Multiple Line  Statically Set Should be done via method call)
		// _editor.Editor.SingleLine = false
		ed = layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {
				// border around everything in the widget
				return widget.Border{
					Color:        f.Decoration.Border.Color,
					CornerRadius: f.Decoration.Border.CornerRadius,
					Width:        f.Decoration.Border.Width,
				}.Layout(
					gtx,
					func(gtx layout.Context) layout.Dimensions {
						return f.Decoration.Pad.Layout(
							gtx,
							func(gtx layout.Context) layout.Dimensions {
								return layout.Flex{
									Axis: layout.Horizontal,
								}.Layout(
									gtx,
									layout.Rigid(
										_lbl.Layout,
									),
									layout.Rigid(
										func(gtx layout.Context) layout.Dimensions {
											_macro := op.Record(gtx.Ops)
											_l := layout.UniformInset(unit.Dp(5)).Layout(
												gtx,
												_editor.Layout,
											)
											_callOps := _macro.Stop()
											_ops := new(op.Ops)
											_pth := clip.RRect{
												Rect: image.Rect(
													0,
													0,
													_l.Size.X,
													_l.Size.Y,
												),
												SE: 5,
												SW: 5,
												NW: 5,
												NE: 5,
											}
											cops := paint.ColorOp{
												Color: f.Decoration.Background,
											}
											cops.Add(_ops)
											pops := paint.PaintOp{}
											pops.Add(_ops)
											paint.FillShape(
												_ops,
												f.Decoration.Border.Color,
												clip.Stroke{
													Path:  _pth.Path(_ops),
													Width: float32(f.Decoration.Border.Width),
												}.Op(),
											)
											_pth.Push(_ops).Pop()
											_callOps.Add(gtx.Ops)
											// log.Println("Layout Dims <", gtx.Constraints, theme.TextSize, _l, ">")
											return _l
										},
									),
								)
							},
						)

					},
				)
			},
		)

		// formFields = append(formFields, ed, layout.Rigid(spacer(unit.Dp(10), unit.Dp(10)).Layout))
	default:
		ed = layout.FlexChild{}
		log.Printf("Unknown Item <% +v>\n", reflect.TypeOf(f.Type))
	}
	return ed
}

// Form
type FormAxisType int

const (
	Vertical FormAxisType = iota
	Horizontal
	Inline
)

type formDecoration struct {
	Border     *widget.Border
	Pad        *layout.Inset
	Background color.NRGBA
}

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
	Decoration formDecoration
	Name       string      // name of the form
	Caption    string      // title displayed
	Id         string      // unique id of the form
	Fields     []*FormItem //form fields
	Axis       FormAxisType
	Disabled   bool
}

func NewForm(name string) *Form {
	idInt := rand.Int63n(time.Now().UnixNano())
	id := names.GetName(int(idInt), "")

	frm := &Form{
		Decoration: formDecoration{
			Border: &widget.Border{
				Color: color.NRGBA{
					R: 0,
					G: 0,
					B: 0,
					A: 128,
				},
				CornerRadius: unit.Dp(5),
				Width:        unit.Dp(1),
			},
			Pad: &layout.Inset{
				Top:    unit.Dp(15),
				Bottom: unit.Dp(15),
				Left:   unit.Dp(15),
				Right:  unit.Dp(15),
			},
			Background: color.NRGBA{
				R: 32,
				G: 32,
				B: 32,
				A: 32,
			},
		},

		Id:   id,
		Name: name,
		Axis: Horizontal,
	}
	return frm
}

func (frm *Form) Add(fld *FormItem) *Form {
	frm.Fields = append(frm.Fields, fld)

	return frm
}
func spacer(width unit.Dp, height unit.Dp) layout.Spacer {
	return layout.Spacer{
		Width:  width,
		Height: height,
	}
}

// to render the form
//
// 1- outer layout is a Flex
//
// 2- inner Layout is a Stack
func (f *Form) render(gtx layout.Context) layout.Dimensions {
	// defining form fields to populate in the form desired
	var formFields []layout.FlexChild = make([]layout.FlexChild, 0)
	for _, fi := range f.Fields {
		formFields = append(formFields, fi.Render(gtx))
		// switch reflect.TypeOf(fi.Type) {
		// case reflect.TypeOf(&widget.Editor{}):
		// 	_lbl := material.Label(theme, unit.Sp(15), fmt.Sprintf("%s : ", fi.LabelText))
		// 	_lbl.Font.Variant = font.Variant("Mono")
		// 	_editor := material.Editor(theme, fi.Type.(*widget.Editor), fi.Hint)
		// 	_editor.Editor.Alignment = text.Middle
		// 	_editor.Editor.InputHint = key.HintAny
		// 	_editor.HintColor = color.NRGBA{B: 236, A: 236}
		// 	_editor.Font.Style = font.Italic
		// 	_editor.Font.Variant = font.Variant(font.Bold.String())
		// 	_editor.Editor.SingleLine = true

		// 	ed := layout.Rigid(
		// 		func(gtx layout.Context) layout.Dimensions {
		// 			// border around everything in the widget
		// 			return widget.Border{
		// 				Color:        fi.Decoration.Border.Color,
		// 				CornerRadius: fi.Decoration.Border.CornerRadius,
		// 				Width:        fi.Decoration.Border.Width,
		// 			}.Layout(
		// 				gtx,
		// 				func(gtx layout.Context) layout.Dimensions {
		// 					return fi.Decoration.Pad.Layout(
		// 						gtx,
		// 						func(gtx layout.Context) layout.Dimensions {
		// 							return layout.Flex{
		// 								Axis: layout.Horizontal,
		// 							}.Layout(
		// 								gtx,
		// 								layout.Rigid(
		// 									_lbl.Layout,
		// 								),
		// 								layout.Flexed(
		// 									0.1,
		// 									func(gtx layout.Context) layout.Dimensions {

		// 										_l := layout.UniformInset(unit.Dp(2)).Layout(
		// 											gtx,
		// 											_editor.Layout,
		// 										)
		// 										_pth := clip.RRect{
		// 											Rect: image.Rect(
		// 												0,
		// 												0,
		// 												_l.Size.X,
		// 												_l.Size.Y,
		// 											),
		// 											SE: 5,
		// 											SW: 5,
		// 											NW: 5,
		// 											NE: 5,
		// 										}
		// 										paint.FillShape(
		// 											gtx.Ops,
		// 											fi.Decoration.Border.Color,
		// 											clip.Stroke{
		// 												Path:  _pth.Path(gtx.Ops),
		// 												Width: float32(fi.Decoration.Border.Width),
		// 											}.Op(),
		// 										)
		// 										log.Println("Layout Dims <", gtx.Constraints, theme.TextSize, _l, ">")
		// 										return _l
		// 									},
		// 								),
		// 							)
		// 						},
		// 					)

		// 				},
		// 			)
		// 		},
		// 	)
		// 	formFields = append(formFields, ed, layout.Rigid(spacer(unit.Dp(10), unit.Dp(10)).Layout))
		// default:
		// 	log.Printf("Unknown Item <% +v>\n", fi)
		// }
	}
	outerLayout := layout.Flex{
		Axis:      layout.Vertical,
		Spacing:   layout.SpaceEnd,
		Alignment: layout.Middle,
	}.Layout(
		gtx,
		layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {
				// log.Println("Flex 1<", gtx.Constraints, ">")
				return widget.Border{
					Color:        f.Decoration.Border.Color,
					CornerRadius: f.Decoration.Border.CornerRadius,
					Width:        f.Decoration.Border.Width,
				}.Layout(
					gtx,
					func(gtx layout.Context) layout.Dimensions {
						return f.Decoration.Pad.Layout(
							gtx,
							func(gtx layout.Context) layout.Dimensions {
								return layout.Flex{
									Axis: layout.Vertical,
								}.Layout(
									gtx,
									formFields...,
								)
							},
						)
					},
				)

			},
		),
		layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {
				log.Println("Flex 2<", gtx.Constraints, ">")
				return material.Label(theme, unit.Sp(12), "I Am Here").Layout(gtx)
			},
		),

		// layout.Expanded(
		// 	func(gtx layout.Context) layout.Dimensions {
		// 		//log.Printf("[Outer Layout Expanded] constraints %+v", gtx.Constraints)
		// 		return layout.Inset{
		// 			Top:    5,
		// 			Bottom: 5,
		// 			Left:   5,
		// 			Right:  5,
		// 		}.Layout(
		// 			gtx,
		// 			func(gtx layout.Context) layout.Dimensions {
		// 				return widget.Border{
		// 					Color:        f.Border.Color,
		// 					CornerRadius: f.Border.CornerRadius,
		// 					Width:        unit.Dp(5), //ToDo replace with f.border.width
		// 				}.Layout(
		// 					gtx,
		// 					func(gtx layout.Context) layout.Dimensions {
		// 						dims := gtx.Constraints
		// 						//log.Printf("[Outer Layout Border ] %+v\n", dims)
		// 						return layout.Dimensions{
		// 							Size: dims.Max,
		// 						}
		// 					},
		// 				)
		// 			},
		// 		)
		// 	},
		// ),
		// layout.Stacked(
		// 	func(gtx layout.Context) layout.Dimensions {
		// 		// log.Printf("[Stacked ]%+v\n", gtx.Constraints)
		// 		if sampleClickable.Clicked() {
		// 			log.Println("Sample Clicked ")
		// 		}
		// 		return material.Button(theme, &sampleClickable, "clickMe").Layout(gtx)
		// 	},
		// ),
	)
	_frmPath := clip.RRect{
		Rect: image.Rect(0, 0, outerLayout.Size.X, outerLayout.Size.Y),
		SE:   int(f.Decoration.Border.CornerRadius),
		SW:   int(f.Decoration.Border.CornerRadius),
		NW:   int(f.Decoration.Border.CornerRadius),
		NE:   int(f.Decoration.Border.CornerRadius),
	}
	paint.FillShape(
		gtx.Ops,
		f.Decoration.Background,
		clip.Outline{
			Path: _frmPath.Path(gtx.Ops),
		}.Op(),
	)
	return outerLayout
	// frmLayout := layout.Flex{
	// 	Axis: func() layout.Axis {
	// 		if f.Axis == Inline {
	// 			return layout.Horizontal
	// 		} else {
	// 			return layout.Vertical
	// 		}
	// 	}(),
	// }
	// return frmLayout.Layout(
	// 	gtx,
	// 	layout.Rigid(
	// 		func(gtx layout.Context) layout.Dimensions {
	// 			return layout.Stack{
	// 				Alignment: layout.Center,
	// 			}.Layout(
	// 				gtx,
	// 				layout.Expanded(func(gtx layout.Context) layout.Dimensions {
	// 					ops := gtx.Ops
	// 					rrect := clip.RRect{
	// 						Rect: image.Rect(5, 1, gtx.Constraints.Max.X-5, gtx.Constraints.Min.Y-1),
	// 						SE:   int(f.Border.CornerRadius),
	// 						SW:   int(f.Border.CornerRadius),
	// 						NW:   int(f.Border.CornerRadius),
	// 						NE:   int(f.Border.CornerRadius),
	// 					}
	// 					defer rrect.Push(ops).Pop()
	// 					bgColor := theme.ContrastBg
	// 					bgColor.A = 128
	// 					cops := paint.ColorOp{
	// 						Color: bgColor,
	// 					}
	// 					cops.Add(ops)
	// 					pops := paint.PaintOp{}
	// 					pops.Add(ops)
	// 					paint.FillShape(
	// 						ops,
	// 						f.Border.Color,
	// 						clip.Stroke{
	// 							Path:  rrect.Path(ops),
	// 							Width: 1,
	// 						}.Op(),
	// 					)
	// 					return layout.Dimensions{
	// 						Size:     gtx.Constraints.Min,
	// 						Baseline: 12,
	// 					}

	// 				}),
	// 				layout.Stacked(
	// 					func(gtx layout.Context) layout.Dimensions {
	// 						return material.Label(theme, unit.Sp(12), "Label In Form").Layout(gtx)
	// 					},
	// 				),
	// 			)
	// 		},
	// 	),
	// )
}
