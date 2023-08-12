/*
*
Steps:

1- Declare a boolean to binding to determine if window is shown or not

2- Declare a new layout context

3- Declare new Operations (ops)

With New Node Window We shall be:

I- Adding A window that:

	a- holds a new container which encapsulates the creation of new nodes
*/
package gui

import (
	"log"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/degreane/dano/utils/binding"
	"github.com/degreane/dano/utils/gui/widgets"
)

var (
	newNodesWindowGUIBinding binding.Boolean = binding.NewBoolean()
)

func newNodesWindowGUI(th *material.Theme) {
	if !newNodesWindowGUIBinding.TryGet() {
		newNodesWindowGUIBinding.Set(true)
		_nOps := new(op.Ops)
		_nOps.Reset()
		var (
			// main container to hold the rest of container items
			_mainContainer *widgets.Container

			_controlsContainer *widgets.Container

			// the number of nodes to be added
			_noOfNodesToAdd *widgets.ContainerItem = widgets.NewContainerItem[widget.Editor](&widget.Editor{}, "No Of Nodes", *th).SetWidgetFilter(Numeric).SetwidgetText("1").SetWidgetTextAlign(text.Middle)

			// if adding in automode then _noOfNodesToAdd is set to whatever needs to be added and boundaries are created and displayed
			// if automode is set to false then _noOfNodesToAdd is set to 1
			_noOfNodesAuto *widgets.ContainerItem = widgets.NewContainerItem[widget.Bool](&widget.Bool{}, "Auto Create", *th)

			// _autoModeContainer  is a subContainer holding the automode fields
			_autoModeContainer *widgets.Container
			_minValue          *widgets.ContainerItem = widgets.NewContainerItem[widget.Editor](&widget.Editor{}, "Min.", *th).SetWidgetFilter(NumericDecimal).SetWidgetTextAlign(text.Middle).SetwidgetText("19.0")
			_maxValue          *widgets.ContainerItem = widgets.NewContainerItem[widget.Editor](&widget.Editor{}, "Max.", *th).SetWidgetFilter(NumericDecimal).SetWidgetTextAlign(text.Middle).SetwidgetText("29.0")

			// manualModeContainer holds the widgets to manually add the node (1 node at a time)
			// _manualModeContainer  *widgets.Container
			_noOfNodesToAddButton *widgets.ContainerItem = widgets.NewContainerItem[widget.Clickable](&widget.Clickable{}, "Add", *th).SetFlexed(true)
		)

		go func() {
			_window := app.NewWindow(app.Title("New Nodes "), app.Size(unit.Dp(600), unit.Dp(200)))
			for windowEvent := range _window.Events() {
				switch eventType := windowEvent.(type) {
				case system.FrameEvent:
					_windowCTX := layout.NewContext(_nOps, eventType) // define the layout.Context
					// Defining The Widgets in the New Window.
					// main container
					_mainContainer = widgets.NewContainer[layout.Flex](&layout.Flex{}).SetAlign(layout.Middle).SetAxis(layout.Vertical)

					// automode container
					_controlsContainer = widgets.NewContainer[layout.Flex](&layout.Flex{}).SetBorder(&widget.Border{
						Width: unit.Dp(0),
					}).SetPadding(&layout.Inset{
						Top: unit.Dp(1), Bottom: unit.Dp(1), Left: unit.Dp(10), Right: unit.Dp(10),
					}).SetAlign(layout.Middle).Add(_noOfNodesToAdd).Add(_noOfNodesAuto)

					//
					_mainContainer.Add(_controlsContainer)

					if _noOfNodesToAddButton.Widget().(*widget.Clickable).Clicked() {
						log.Printf("I am Clicked with Value of <%s>\n", _noOfNodesToAdd.GetValue())
					}
					if _noOfNodesAuto.GetValue() == "true" {
						_autoModeContainer = widgets.NewContainer[layout.Flex](&layout.Flex{})
						_autoModeContainer.Add(_minValue).Add(_maxValue)
						_mainContainer.Add(_autoModeContainer)
						_minValue.Call(widgets.FloatEditor)
						_maxValue.Call(widgets.FloatEditor)
					}
					if _noOfNodesAuto.Widget().(*widget.Bool).Changed() {
						if _noOfNodesAuto.GetValue() == "true" {
							log.Println("AutoMode True")
						} else {
							log.Println("AutoMode False")
						}
					}
					_mainContainer.Add(_noOfNodesToAddButton)
					_mainContainer.Render(_windowCTX)
					eventType.Frame(_windowCTX.Ops)
				case system.DestroyEvent:
					newNodesWindowGUIBinding.Set(false)
				}
			}
		}()
	}
}
