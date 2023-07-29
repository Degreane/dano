package gui

import (
	"image"
	"image/color"
	"log"
	"regexp"

	"fmt"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/degreane/dano/utils/binding"
)

const (
	Alpha          = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	AlphaNumeric   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	AlphaSmall     = "abcdefghijklmnopqrstuvwxyz"
	AlphaCapital   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numeric        = "0123456789"
	NumericDecimal = "0123456789."
)

var (
	newSamplingWindowDisplayed bool      = false
	newSamplingWindowDone      chan bool = make(chan bool, 1)

	newSamplinWindowBinding binding.Boolean = binding.NewBoolean()
	// newSamplingBatchNameBinding      binding.String  = binding.NewString()
	// newSamplingBatchInfoBinding      binding.String  = binding.NewString()
	// newSamplingBatchThresholdBinding binding.String  = binding.NewString()
	// newSamplingBatchPrecisionBinding binding.String  = binding.NewString()

	newSamplingBatchNameState      widget.Editor = widget.Editor{}
	newSamplingBatchInfoState      widget.Editor = widget.Editor{}
	newSamplingBatchPrecisionState widget.Editor = widget.Editor{}
	newSamplingBatchThresholdState widget.Editor = widget.Editor{}

	// newFlexSubContainer2  layout.Flex

)

func newSamplingGUI() {
	if !newSamplinWindowBinding.TryGet() {
		newSamplinWindowBinding.Set(true)
		var samplingWindowOps *op.Ops = new(op.Ops)
		samplingWindowOps.Reset()
		newSamplinWindowBinding.AddListener(func() {
			if !newSamplinWindowBinding.TryGet() {
				tbItemNew.Enable()
				tbItemLoad.Enable()
				w.Invalidate()
			} else {
				tbItemNew.Disable()
				tbItemLoad.Disable()
			}
		})

		go func() {
			newSamplingWindow := app.NewWindow(
				app.Title("New Batch Sampling"),
				app.Size(unit.Dp(800), unit.Dp(600)),
			)
			newSamplingWindow.Option(func(m unit.Metric, c *app.Config) {
				c.Size = image.Pt(800, 600)
			})
			// define the fields we need:
			var (
				// _form *Form = NewForm("NewBatch")
				_cntr              *Container
				_batchName         *widget.Editor    = new(widget.Editor)
				_batchInfo         *widget.Editor    = new(widget.Editor)
				_batchPrecision    *widget.Editor    = new(widget.Editor)
				_batchThreshold    *widget.Editor    = new(widget.Editor)
				_batchSubmitButton *widget.Clickable = new(widget.Clickable)

				_ColorWhite = color.NRGBA{
					R: 255,
					G: 255,
					B: 255,
					A: 0,
				}
			)

			_cntr0 := NewContainer[layout.Flex](&layout.Flex{}).SetAxis(
				layout.Vertical,
			).SetBorder(
				&widget.Border{
					Color: color.NRGBA{
						R: 255,
						A: 128,
					},
					CornerRadius: unit.Dp(2),
					Width:        unit.Dp(0),
				},
			).SetBackgroundColor(_ColorWhite)
			_cntr1 := NewContainer[layout.Flex](&layout.Flex{}).SetAxis(
				layout.Horizontal,
			).SetBorder(
				&widget.Border{
					Color: color.NRGBA{
						R: 255,
						A: 128,
					},
					CornerRadius: unit.Dp(2),
					Width:        unit.Dp(0),
				},
			).SetBackgroundColor(_ColorWhite)
			_cntr = NewContainer[layout.Flex](&layout.Flex{}).SetAxis(
				layout.Vertical,
			).SetBorder(
				&widget.Border{
					Color: color.NRGBA{
						R: 255,
						A: 128,
					},
					CornerRadius: unit.Dp(2),
					Width:        unit.Dp(1),
				},
			).SetAlign(
				layout.Start,
			).SetMargin(
				&layout.Inset{
					Top:    unit.Dp(10),
					Bottom: unit.Dp(10),
					Left:   unit.Dp(10),
					Right:  unit.Dp(10),
				},
			).SetPadding(
				&layout.Inset{
					Top:    unit.Dp(10),
					Bottom: unit.Dp(10),
					Left:   unit.Dp(10),
					Right:  unit.Dp(10),
				},
			).SetBackgroundColor(color.NRGBA{
				R: 25,
				G: 192,
				B: 192,
				A: 25,
			}).SetSpacing(layout.SpaceEnd)

			_cntrSubmit := NewContainer[layout.Flex](&layout.Flex{}).SetAxis(
				layout.Horizontal,
			).SetBorder(
				&widget.Border{
					Color: color.NRGBA{
						R: 255,
						A: 128,
					},
					CornerRadius: unit.Dp(2),
					Width:        unit.Dp(0),
				},
			)
			c1 := NewContainerItem[widget.Editor](
				_batchName,
				"Batch Name",
			).SetWidgetHint(
				"new batch name",
			).SetWidgetTextAlign(
				text.Middle,
			).SetWidgetMultiLine(
				false,
			).SetWidgetFilter(AlphaNumeric)
			c2 := NewContainerItem[widget.Editor](_batchInfo, "Information").SetWidgetHint(
				"new Batch Information",
			).SetWidgetMultiLine(
				true,
			).SetWidgetTextAlign(
				text.Start,
			).SetWidgetFilter("")
			c3 := NewContainerItem[widget.Editor](
				_batchPrecision,
				"Precision",
			).SetwidgetText(
				"2",
			).SetWidgetFilter(
				Numeric,
			).SetWidgetMultiLine(
				false,
			).SetWidgetTextAlign(
				text.Middle,
			).SetWidgetHint(
				"Precision (Default 2.0)",
			)
			c4 := NewContainerItem[widget.Editor](
				_batchThreshold,
				"Threshold",
			).SetWidgetFilter(
				NumericDecimal,
			).SetwidgetText(
				"3.0",
			).SetWidgetMultiLine(
				false,
			).SetWidgetTextAlign(
				text.Middle,
			).SetWidgetHint(
				"Threshold (default 3.0)",
			)

			submitContainerBtnItem := NewContainerItem[widget.Clickable](_batchSubmitButton, "Submit")
			_cntrSubmit.Add(submitContainerBtnItem.Disable())
			_cntr1.Add(c3).Add(c4).SetBackgroundColor(color.NRGBA{G: 128, A: 128})
			_cntr0.Add(c1).Add(c2).Add(_cntr1.SetPadding(&layout.Inset{}).SetMargin(&layout.Inset{})).Add(_cntrSubmit)

			_cntr.Add(_cntr0)

			_newSamplingBatchInfo := NewFormItem[widget.Editor](&newSamplingBatchInfoState, "Info").SetMultiLine(true).SetFilter("")
			_newSamplingBatchInfo.SetHint("Batch Information")
			for windowEvent := range newSamplingWindow.Events() {

				switch windowEventType := windowEvent.(type) {
				case system.FrameEvent:
					samplingWindowGtx := layout.NewContext(samplingWindowOps, windowEventType)
					c4.Call(func(c *ContainerItem) *ContainerItem {

						_filter := regexp.MustCompile(`^[0-9]+(\.)?([0-9]+)?$`)
						_dot := regexp.MustCompile(`\.`)
						_widget := c.widget.(*widget.Editor)
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
						} // should implement decoration of the containerItem

						return c
					})
					if _batchSubmitButton.Clicked() {
						log.Println("Value of <", c1.GetValue(), ">")
					}
					if len(c1.GetValue()) == 0 {
						submitContainerBtnItem.Disable()
					} else {
						submitContainerBtnItem.Enable()
					}
					_cntr.Render(samplingWindowGtx)
					// log.Println("Dimensions are <", _cntr.Dims, ">")
					windowEventType.Frame(samplingWindowGtx.Ops)
				case system.DestroyEvent:

					newSamplinWindowBinding.Set(false)

				}
			}
		}()
	}

}
