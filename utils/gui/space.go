package gui

import (
	"gioui.org/layout"
	"gioui.org/unit"
)

func spaceW(sp unit.Dp) layout.FlexChild {
	return layout.Rigid(
		func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{Width: sp, Height: sp}.Layout(gtx)
		},
	)
}
