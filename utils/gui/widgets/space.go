package widgets

import (
	"gioui.org/layout"
	"gioui.org/unit"

)

func SpaceW(sp unit.Dp) layout.FlexChild {
	return layout.Rigid(
		func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{Width: sp, Height: sp}.Layout(gtx)
		},
	)
}
