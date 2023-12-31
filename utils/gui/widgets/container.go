// a container is a form of layout that embeds other formats inside it

package widgets

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"reflect"
	"regexp"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

)

type ContainerItemType interface {
	*widget.Editor | widget.Editor | *widget.Clickable | widget.Clickable | *widget.Bool | widget.Bool | *widget.Label | widget.Label | *widget.Float | material.LabelStyle | layout.FlexChild | *layout.FlexChild
}
type ContainerItemLayout uint

const (
	itemHorizontal ContainerItemLayout = iota
	itemVertical
)

// FloatEditor Actually enforces the containerItem Editor Widget to only accept floating numbers
func FloatEditor(c *ContainerItem) *ContainerItem {
	_filter := regexp.MustCompile(`^[0-9]+(\.)?([0-9]+)?$`)
	_dot := regexp.MustCompile(`\.`)
	_widget := c.Widget().(*widget.Editor)
	_widgetBytes := []byte(_widget.Text())
	if len(_widgetBytes) > 0 {
		if !_filter.Match(_widgetBytes) {
			loc := _dot.FindIndex(_widgetBytes)
			_widgetBytes = _dot.ReplaceAll(_widgetBytes, []byte{})
			_newStr := fmt.Sprintf("%s%s%s", _widgetBytes[:loc[0]], []byte{'.'}, _widgetBytes[loc[0]:])
			// log.Println("Regexp Index is <", loc, ">")
			_, _caretPos := _widget.CaretPos()
			_widget.SetText(_newStr)
			if len(_widget.Text()) > _caretPos && _caretPos > 0 {
				_widget.SetCaret(_caretPos-1, _caretPos)
			} else {
				_widget.SetCaret(len(_widget.Text()), len(_widget.Text()))
			}
		}
	}
	return c
}

// ContainerItem contains a widget that returns a layout.dimension
//
// # It has a Label
//
// # TODO Introduce Label Decoration, with getters and setters
//
// generalize widget decorations into widgetDecoration
type ContainerItem struct {
	widget          interface{}
	label           string
	layout          ContainerItemLayout
	widgetType      string
	widgetFontColor color.NRGBA
	widgetBgColor   color.NRGBA
	widgetHintColor color.NRGBA
	flexed          bool
	hint            string
	disabled        bool
	theme           material.Theme
	labelDecoration struct {
		backgroundColor color.NRGBA
		foregroundColor color.NRGBA
		pad             layout.Inset
		align           layout.Alignment
		margin          layout.Inset
		border          widget.Border
	}
}

func NewContainerItem[T ContainerItemType](w *T, lbl string, theme material.Theme) *ContainerItem {
	wt := reflect.TypeOf(w).String()
	// log.Println("Ths new Container Item type is <", wt, ">")
	cnt := &ContainerItem{
		theme:      theme,
		disabled:   false,
		widget:     w,
		label:      lbl,
		layout:     itemHorizontal,
		widgetType: wt,
		widgetFontColor: color.NRGBA{
			A: 255,
		},
		widgetBgColor: color.NRGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 255,
		},
		widgetHintColor: color.NRGBA{
			B: 255,
			A: 255,
		},
	}
	return cnt
}


func (ci *ContainerItem) Call(fn func(c *ContainerItem) *ContainerItem) *ContainerItem {
	return fn(ci)
}
func (ci *ContainerItem) SetWidgetHint(hint string) *ContainerItem {
	ci.hint = hint
	return ci
}
func (ci *ContainerItem) SetWidgetTextAlign(textAlign text.Alignment) *ContainerItem {
	if ci.widgetType == "*widget.Editor" {
		ci.widget.(*widget.Editor).Alignment = textAlign
	} else if ci.widgetType == "*widget.Label" {
		// log.Println("Setting Align for Label")
		ci.widget.(*widget.Label).Alignment = textAlign
	}
	return ci
}
func (ci *ContainerItem) SetWidgetMultiLine(multi bool) *ContainerItem {
	if ci.widgetType == "*widget.Editor" {
		ci.widget.(*widget.Editor).SingleLine = !multi
	}
	return ci
}
func (ci *ContainerItem) SetWidgetFilter(filter string) *ContainerItem {
	if ci.widgetType == "*widget.Editor" {
		ci.widget.(*widget.Editor).Filter = filter
	}
	return ci
}

func (ci *ContainerItem) SetLabel(lbl string) *ContainerItem {
	ci.label = lbl
	return ci
}
func (ci *ContainerItem) SetwidgetText(txt string) *ContainerItem {
	if ci.widgetType == "*widget.Editor" {
		ci.widget.(*widget.Editor).SetText(txt)
	} else if ci.widgetType == "*widget.Label" {
		ci.SetLabel(txt)
	}
	return ci
}
func (ci *ContainerItem) GetValue() string {
	if ci.widgetType == "*widget.Editor" {
		return ci.widget.(*widget.Editor).Text()
	} else if ci.widgetType == "*widget.Label" {
		return ci.label
	} else if ci.widgetType == "*widget.Bool" {
		return fmt.Sprintf("%t", ci.widget.(*widget.Bool).Value)
	}
	return ""
}
func (ci *ContainerItem) SetWidgetColor(clr color.NRGBA) *ContainerItem {
	ci.widgetFontColor = clr
	return ci
}
func (ci *ContainerItem) SetWidgetHintColor(clr color.NRGBA) *ContainerItem {
	ci.widgetHintColor = clr
	return ci
}
func (ci *ContainerItem) renderLabel(gtx layout.Context) layout.Dimensions {

	// log.Println("Label Item Constraints <", gtx.Constraints, ">")
	_macro := op.Record(gtx.Ops)
	mr := material.Label(&ci.theme, ci.theme.TextSize, ci.label)
	mr.Alignment = text.Middle
	r := mr.Layout(gtx)
	_callOps := _macro.Stop()
	_ops := new(op.Ops)

	_cops := paint.ColorOp{
		Color: color.NRGBA{},
	}
	_cops.Add(_ops)
	_pops := paint.PaintOp{}
	_pops.Add(_ops)
	_rect := clip.Rect{
		// Min: image.Pt(gtx.Constraints.Min.X, gtx.Constraints.Min.Y),
		Max: r.Size,
	}
	paint.FillShape(gtx.Ops, color.NRGBA{
		R: 255,
		G: 255,
		B: 255,
		A: 0,
	}, clip.Outline{
		Path: _rect.Path(),
	}.Op())
	_rect.Push(_ops).Pop()
	_callOps.Add(gtx.Ops)
	return r
}
func (ci *ContainerItem) Disable() *ContainerItem {
	ci.disabled = true
	return ci
}
func (ci *ContainerItem) Enable() *ContainerItem {
	ci.disabled = false
	return ci
}

func (ci *ContainerItem) renderWidget(gtx layout.Context) layout.Dimensions {
	// log.Println("Label Item Constraints <", gtx.Constraints, ci.widgetType, ">")
	_macro := op.Record(gtx.Ops)
	// r := material.Label(theme, unit.Sp(15), ci.label).Layout(gtx)
	var r layout.Dimensions
	var _max image.Point
	if ci.widgetType == "*widget.Editor" {
		w := ci.widget.(*widget.Editor)
		m := material.Editor(&ci.theme, w, ci.hint)
		m.Color = ci.widgetFontColor
		m.HintColor = ci.widgetHintColor
		r = m.Layout(gtx)

		_max = r.Size
	} else if ci.widgetType == "*widget.Clickable" {
		w := ci.widget.(*widget.Clickable)
		ci.theme.TextSize = unit.Sp(ci.theme.TextSize)
		r = material.Button(&ci.theme, w, ci.label).Layout(gtx)
		_max = r.Size
		// log.Println("We Have a Max of <", _max, ">")
	} else {
		log.Println("Getting  <", ci.widgetType, ">")
	}
	_callOps := _macro.Stop()
	_ops := new(op.Ops)

	_cops := paint.ColorOp{
		Color: color.NRGBA{},
	}
	_cops.Add(_ops)
	_pops := paint.PaintOp{}
	_pops.Add(_ops)
	_rect := clip.Rect{
		// Min: image.Pt(gtx.Constraints.Min.X, gtx.Constraints.Min.Y),
		Max: _max,
	}
	paint.FillShape(gtx.Ops, color.NRGBA{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	}, clip.Outline{
		Path: _rect.Path(),
	}.Op())
	_rect.Push(_ops).Pop()
	_callOps.Add(gtx.Ops)
	return r
}
func (ci *ContainerItem) SetFlexed(flexed bool) *ContainerItem {
	ci.flexed = flexed
	return ci
}
func (ci *ContainerItem) Widget() interface{} {
	return ci.widget
}
func (ci *ContainerItem) Render(oGtx layout.Context) layout.Dimensions {
	// log.Println("Container Item render GTX <", gtx.Constraints, ">")
	// if the layout is itemInline then both the label and widget are laid side by side
	var gtx layout.Context
	if ci.disabled {
		gtx = oGtx.Disabled()
	} else {
		gtx = oGtx
	}
	_macro := op.Record(gtx.Ops) // start macro recording
	var l layout.Flex = layout.Flex{
		WeightSum: 100,
		Spacing:   layout.SpaceEnd,
	}
	if ci.layout == itemHorizontal {
		l.Axis = layout.Horizontal
	} else {
		l.Axis = layout.Vertical
	}

	var dims layout.Dimensions
	if ci.widgetType == "*widget.Editor" {
		dims = l.Layout(
			gtx,
			layout.Flexed(
				25,
				func(gtx layout.Context) layout.Dimensions {
					return layout.UniformInset(unit.Dp(5)).Layout(
						gtx,
						ci.renderLabel,
					)
				},
			),
			layout.Flexed(
				75,
				func(gtx layout.Context) layout.Dimensions {
					return layout.UniformInset(unit.Dp(5)).Layout(
						gtx,
						ci.renderWidget,
					)
				},
			),
		)
	} else if ci.widgetType == "*widget.Clickable" {
		if ci.flexed {
			dims = l.Layout(
				gtx,
				layout.Flexed(
					100,
					func(gtx layout.Context) layout.Dimensions {
						return layout.UniformInset(unit.Dp(5)).Layout(
							gtx,
							ci.renderWidget,
						)
					},
				),
			)
		} else {
			dims = l.Layout(
				gtx,
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return layout.UniformInset(unit.Dp(5)).Layout(
							gtx,
							ci.renderWidget,
						)
					},
				),
			)
		}

	} else if ci.widgetType == "*widget.Label" {
		dims = l.Layout(
			gtx,
			layout.Flexed(
				100,
				func(gtx layout.Context) layout.Dimensions {
					return widget.Border{
						Color: color.NRGBA{
							R: 255,
							A: 128,
						},
						CornerRadius: unit.Dp(1),
						Width:        unit.Dp(1),
					}.Layout(
						gtx,
						func(gtx layout.Context) layout.Dimensions {
							return layout.UniformInset(unit.Dp(5)).Layout(
								gtx,
								ci.renderLabel,
							)
						},
					)
				},
			),
		)
	} else if ci.widgetType == "*layout.FlexChild" {
		if ci.flexed {
			dims = l.Layout(
				gtx,
				layout.Flexed(
					100,
					func(gtx layout.Context) layout.Dimensions {
						return layout.Spacer{}.Layout(gtx)
					},
				),
			)
		} else {
			dims = l.Layout(
				gtx,
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						return layout.Spacer{}.Layout(gtx)
					},
				),
			)
		}
	} else if ci.widgetType == "*widget.Bool" {
		dims = l.Layout(
			gtx,
			layout.Rigid(
				func(gtx layout.Context) layout.Dimensions {
					return material.CheckBox(&ci.theme, ci.widget.(*widget.Bool), ci.label).Layout(gtx)
				},
			),
		)
	}

	// log.Println("Container Item Dims  <", dims, ">")
	_callOps := _macro.Stop() // stop macro recording
	_ops := new(op.Ops)       // start New Ops in order to insert the new Container Item inside it
	_ops.Reset()
	_rect := clip.RRect{
		Rect: image.Rect(0, 0, dims.Size.X, dims.Size.Y),
		SE:   5,
		SW:   5,
		NW:   5,
		NE:   5,
	}
	_cops := paint.ColorOp{
		Color: color.NRGBA{
			G: 128,
			A: 128,
		},
	} // color Operation
	_cops.Add(_ops)
	_pops := paint.PaintOp{}
	_pops.Add(_ops)
	paint.FillShape(gtx.Ops, color.NRGBA{
		B: 128,
		A: 128,
	}, clip.Outline{
		Path: _rect.Path(gtx.Ops),
	}.Op())
	paint.FillShape(gtx.Ops, color.NRGBA{
		B: 128,
		A: 255,
	}, clip.Stroke{
		Path:  _rect.Path(gtx.Ops),
		Width: 1,
	}.Op())
	_rect.Push(_ops).Pop()
	_callOps.Add(gtx.Ops)
	return dims
}

type ContainerType interface {
	layout.Flex | *layout.Flex | layout.Stack | *layout.Stack
}

type ContainerDecoration struct {
	Border     *widget.Border
	Pad        *layout.Inset
	Margin     *layout.Inset
	Background color.NRGBA
}
type Container struct {
	Layout     interface{}
	Children   []interface{}
	Ctx        layout.Context
	Theme      *material.Theme
	Decoration ContainerDecoration
	Dims       layout.Constraints
}

func NewContainer[T ContainerType](lout *T) *Container {
	decorationBorder := &widget.Border{
		Color:        color.NRGBA{A: 255},
		CornerRadius: unit.Dp(5),
		Width:        unit.Dp(1),
	}
	decorationMargin := &layout.Inset{
		Top:    unit.Dp(5),
		Bottom: unit.Dp(5),
		Left:   unit.Dp(5),
		Right:  unit.Dp(5),
	}
	decorationPad := &layout.Inset{
		Top:    unit.Dp(5),
		Bottom: unit.Dp(5),
		Left:   unit.Dp(5),
		Right:  unit.Dp(5),
	}
	decoration := ContainerDecoration{
		Border:     decorationBorder,
		Margin:     decorationMargin,
		Pad:        decorationPad,
		Background: color.NRGBA{B: 128, A: 128},
	}
	cntr := &Container{
		Layout:     lout,
		Ctx:        layout.Context{},
		Theme:      new(material.Theme),
		Decoration: decoration,
	}
	return cntr
}

func (c *Container) Add(child interface{}) *Container {
	// switch reflect.TypeOf(child) {
	// case reflect.TypeOf(&Container{}):
	// 	child.(*Container).Ctx = c.Ctx

	// }
	// log.Println("COntainer Add Child <", reflect.TypeOf(child).String(), ">")
	c.Children = append(c.Children, child)
	return c
}
func (c *Container) SetBackgroundColor(clr color.NRGBA) *Container {
	c.Decoration.Background = clr
	return c
}
func (c *Container) SetAxis(axis layout.Axis) *Container {
	switch reflect.TypeOf(c.Layout) {
	case reflect.TypeOf(&layout.Flex{}):
		c.Layout.(*layout.Flex).Axis = axis
	}
	return c
}
func (c *Container) SetAlign(align layout.Alignment) *Container {
	switch reflect.TypeOf(c.Layout) {
	case reflect.TypeOf(&layout.Flex{}):
		c.Layout.(*layout.Flex).Alignment = align
	}
	return c
}
func (c *Container) SetBorder(brdr *widget.Border) *Container {
	c.Decoration.Border = brdr
	return c
}
func (c *Container) SetMargin(mrgn *layout.Inset) *Container {
	c.Decoration.Margin = mrgn
	return c
}
func (c *Container) SetPadding(pad *layout.Inset) *Container {
	c.Decoration.Pad = pad
	return c
}
func (c *Container) SetSpacing(spc layout.Spacing) *Container {
	switch reflect.TypeOf(c.Layout) {
	case reflect.TypeOf(&layout.Flex{}):
		c.Layout.(*layout.Flex).Spacing = spc
	}
	return c
}

func (c *Container) Render(gtx layout.Context) layout.Dimensions {
	var dims layout.Dimensions
	c.Ctx = gtx
	switch reflect.TypeOf(c.Layout) {
	case reflect.TypeOf(&layout.Flex{}):
		// log.Println("Rendering Flex Container")
		var children []layout.FlexChild = make([]layout.FlexChild, 0)
		var nOfChildren int = len(c.Children)
		c.Layout.(*layout.Flex).WeightSum = 100
		for _, child := range c.Children {
			switch reflect.TypeOf(child) {
			case reflect.TypeOf(&Container{}):
				v, _ := child.(*Container)
				// children = append(children, layout.Flexed(
				// 	float32(100/nOfChildren),
				// 	func(gtx layout.Context) layout.Dimensions {
				// 		return v.Render(gtx)
				// 	},
				// ))
				if c.Layout.(*layout.Flex).Axis == layout.Horizontal {
					children = append(children, layout.Flexed(
						float32(100/nOfChildren),
						func(gtx layout.Context) layout.Dimensions {
							return v.Render(gtx)
						},
					))
				} else {
					children = append(children, layout.Rigid(
						// float32(100/nOfChildren),
						func(gtx layout.Context) layout.Dimensions {

							return v.Render(gtx)

						},
					))
				}
			case reflect.TypeOf(&FormItem{}):
				v, _ := child.(*FormItem)
				children = append(children, v.Render(c.Ctx))
			case reflect.TypeOf(&ContainerItem{}):
				v, _ := child.(*ContainerItem)
				if c.Layout.(*layout.Flex).Axis == layout.Horizontal {
					children = append(children, layout.Flexed(
						float32(100/nOfChildren),
						func(gtx layout.Context) layout.Dimensions {
							return v.Render(gtx)
						},
					))
				} else {
					children = append(children, layout.Rigid(
						// float32(100/nOfChildren),
						func(gtx layout.Context) layout.Dimensions {

							return v.Render(gtx)

						},
					))
				}

			}

		}
		_macro := op.Record(c.Ctx.Ops)

		dims = c.Decoration.Margin.Layout( // margin
			c.Ctx,
			func(gtx layout.Context) layout.Dimensions {
				// log.Println("1st Margin <", gtx.Constraints, ">")
				return c.Decoration.Border.Layout( // border
					gtx,
					func(gtx layout.Context) layout.Dimensions {
						// log.Println("2nd Border <", gtx.Constraints, ">")
						return c.Decoration.Pad.Layout(
							gtx,

							func(gtx layout.Context) layout.Dimensions {
								if c.Layout.(*layout.Flex).Axis == layout.Horizontal {
									return layout.Flex{
										Axis:    layout.Vertical,
										Spacing: layout.SpaceEnd,
									}.Layout(
										gtx,
										layout.Rigid(
											func(gtx layout.Context) layout.Dimensions {
												return c.Layout.(*layout.Flex).Layout(
													gtx,
													children...,
												)
											},
										),
									)
								} else {
									// children = append(children, layout.Flexed(
									// 	20,
									// 	func(gtx layout.Context) layout.Dimensions {
									// 		return layout.Spacer{
									// 			Height: 100,
									// 		}.Layout(gtx)
									// 	},
									// ))
									return c.Layout.(*layout.Flex).Layout(
										gtx,
										children...,
									)
								}

							},
						)

					},
				)
			},
		)
		_callOp := _macro.Stop()
		_ops := new(op.Ops)
		_rect := image.Rect(
			int(c.Decoration.Margin.Left)+int(c.Decoration.Border.Width),
			int(c.Decoration.Margin.Top)+int(c.Decoration.Border.Width),
			dims.Size.X-(int(c.Decoration.Margin.Right)+int(c.Decoration.Border.Width)),
			dims.Size.Y-(int(c.Decoration.Margin.Bottom)+int(c.Decoration.Border.Width)),
		)
		// log.Println("Rectangle <", _rect, ">")
		_pth := clip.RRect{
			Rect: _rect,
			SE:   int(c.Decoration.Border.CornerRadius),
			SW:   int(c.Decoration.Border.CornerRadius),
			NW:   int(c.Decoration.Border.CornerRadius),
			NE:   int(c.Decoration.Border.CornerRadius),
		}

		cops := paint.ColorOp{
			Color: c.Decoration.Background,
		}
		cops.Add(_ops)
		pops := paint.PaintOp{}
		pops.Add(_ops)

		// log.Println("Pth <", _pth.Rect.Size(), ">")
		c.Dims = c.Ctx.Constraints
		paint.FillShape(c.Ctx.Ops, c.Decoration.Background, clip.Outline{
			Path: _pth.Path(c.Ctx.Ops),
		}.Op())
		_pth.Push(_ops).Pop()
		_callOp.Add(c.Ctx.Ops)
	default:
		log.Println("Default Layout <", reflect.TypeOf(c.Layout), ">")
	}
	// log.Println("DIms Size <", dims, ">")

	return dims
}
