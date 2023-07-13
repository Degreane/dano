/// Implementing New toobar custom Item widget

package guifyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type PsyButtonToolbar interface {
	widget.ToolbarItem
	SetHighPriority()
	SetNormalPriority()
	Settext(t string)
	SetDisable(b bool)
}
type customToolbarItem struct {
	*widget.Button
}

func (l *customToolbarItem) ToolbarObject() fyne.CanvasObject {
	return l.Button
}
func NewCustomToolbarButtonItem(text string, icon fyne.Resource, action func()) PsyButtonToolbar {
	l := widget.NewButtonWithIcon(text, icon, action)
	l.MinSize()
	l.Alignment = widget.ButtonAlignCenter
	l.IconPlacement = widget.ButtonIconLeadingText
	return &customToolbarItem{l}
}
func (l *customToolbarItem) SetDisable(t bool) {
	if t {
		l.Button.Disable()
	} else {
		l.Button.Enable()
	}
}
func (l *customToolbarItem) SetHighPriority() {
	l.Button.Importance = widget.DangerImportance
	//l.Importance = widget.HighImportance
	l.Refresh()
}
func (l *customToolbarItem) SetNormalPriority() {
	l.Button.Importance = widget.LowImportance
	l.Refresh()
}
func (l *customToolbarItem) Settext(t string) {
	l.SetText(t)
	l.Refresh()
}

type PsyLabelToolbar interface {
	widget.ToolbarItem
	SetHighPriority()
	SetNormalPriority()
	Settext(str string)
	SetBinding(bind binding.String)
}

type customLabelToolbarItem struct {
	*widget.Label
}

func NewCustomToolBarLabelItem(str string) PsyLabelToolbar {
	l := widget.NewLabel(str)
	l.Alignment = fyne.TextAlignCenter
	return &customLabelToolbarItem{l}
}
func (l *customLabelToolbarItem) ToolbarObject() fyne.CanvasObject {
	return l.Label
}
func (l *customLabelToolbarItem) SetHighPriority() {
	l.Label.TextStyle.Bold = true
	l.Label.TextStyle.Italic = true
}
func (l *customLabelToolbarItem) SetNormalPriority() {
	l.Label.TextStyle.Bold = false
	l.Label.TextStyle.Italic = false
}
func (l *customLabelToolbarItem) Settext(str string) {
	l.Label.SetText(str)
}
func (l *customLabelToolbarItem) SetBinding(b binding.String) {
	l.Bind(b)
}
