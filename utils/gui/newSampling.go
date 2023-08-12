package gui

import (
	"image/color"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"

	"github.com/degreane/dano/utils/binding"
	"github.com/degreane/dano/utils/gui/widgets"
)

var (
	newSamplinWindowBinding binding.Boolean = binding.NewBoolean()
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
				app.Size(unit.Dp(500), unit.Dp(300)),
			)
			// newSamplingWindow.Option(func(m unit.Metric, c *app.Config) {
			// 	c.Size = image.Pt(400, 300)
			// })
			// define the fields we need:
			var (
				// _form *Form = NewForm("NewBatch")
				_cntr              *widgets.Container
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

			_cntr0 := widgets.NewContainer[layout.Flex](&layout.Flex{}).SetAxis(
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
			_cntr1 := widgets.NewContainer[layout.Flex](&layout.Flex{}).SetAxis(
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
			_cntr = widgets.NewContainer[layout.Flex](&layout.Flex{}).SetAxis(
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

			_cntrSubmit := widgets.NewContainer[layout.Flex](&layout.Flex{}).SetAxis(
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
			//C1 is the batch Name Container Item
			c1 := widgets.NewContainerItem[widget.Editor](
				_batchName,
				"Batch Name",
				*theme,
			).SetWidgetHint(
				"new batch name",
			).SetWidgetTextAlign(
				text.Middle,
			).SetWidgetMultiLine(
				false,
			).SetWidgetFilter(AlphaNumeric)
			// C2 is the batch Information Container Item
			c2 := widgets.NewContainerItem[widget.Editor](_batchInfo, "Information", *theme).SetWidgetHint(
				"new Batch Information",
			).SetWidgetMultiLine(
				true,
			).SetWidgetTextAlign(
				text.Start,
			).SetWidgetFilter("")

			// C3 is the Batch Precision Container Item
			c3 := widgets.NewContainerItem[widget.Editor](
				_batchPrecision,
				"Precision",
				*theme,
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

			// C4 is the Batch Threshold Container Item
			c4 := widgets.NewContainerItem[widget.Editor](
				_batchThreshold,
				"Threshold",
				*theme,
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
			// submitContainerBtnItem is the submit Button Container Item
			submitContainerBtnItem := widgets.NewContainerItem[widget.Clickable](_batchSubmitButton, "Submit", *theme)
			_cntrSubmit.Add(submitContainerBtnItem.SetFlexed(true).Disable())
			_cntr1.Add(c3).Add(c4).SetBackgroundColor(color.NRGBA{G: 128, A: 128})
			_cntr0.Add(c1).Add(c2).Add(_cntr1.SetPadding(&layout.Inset{}).SetMargin(&layout.Inset{})).Add(_cntrSubmit)
			_cntr.Add(_cntr0)
			for windowEvent := range newSamplingWindow.Events() {
				switch windowEventType := windowEvent.(type) {
				case system.FrameEvent:
					samplingWindowGtx := layout.NewContext(samplingWindowOps, windowEventType)
					c4.Call(widgets.FloatEditor)
					if _batchSubmitButton.Clicked() {
						tbItemNew.Enable()
						tbItemSave.Enable()
						tbItemLoad.Disable()
						tbItemInfo.SetwidgetText(c2.GetValue())
						tbItemName.SetwidgetText(c1.GetValue())
						tbItemPrecision.SetwidgetText(c3.GetValue())
						tbItemThreshold.SetwidgetText(c4.GetValue())
						// newSamplingWindow.Perform(system.Action(system.RTL))
						tbItemSampleCount.SetwidgetText("0")
						tbItemSampleNewNodeButton.Enable()
						newSamplingWindow.Perform(system.ActionClose)
					}
					if len(c1.GetValue()) == 0 {
						submitContainerBtnItem.Disable()
					} else {
						submitContainerBtnItem.Enable()
						tbItemName.SetLabel(c1.GetValue())
						tbItemPrecision.SetwidgetText(c3.GetValue())
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
