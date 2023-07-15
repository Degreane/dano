/// Implementing New toobar custom Item widget

package guifyne

import (
	"errors"
	"log"
	"regexp"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

// button PsyButtonToolbar
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

// End of PsyButtonToolbar

// Label PsyLabelToolbar
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
} // End of PsyLabelToolbar

// CheckBox PsyCheckboxToolbar
type PsyCheckboxToolbar interface {
	widget.ToolbarItem
	Settext(str string)
	SetBinding(bind binding.Bool)
}
type customCheckToolbarItem struct {
	*widget.Check
}

func NewCustomToolbarCheckItem(str string, fn func(bool)) PsyCheckboxToolbar {
	l := widget.NewCheck(str, fn)
	return &customCheckToolbarItem{l}
}
func (l *customCheckToolbarItem) SetBinding(bind binding.Bool) {
	l.Bind(bind)
}
func (l *customCheckToolbarItem) Settext(str string) {
	l.Text = str
}
func (l *customCheckToolbarItem) ToolbarObject() fyne.CanvasObject {
	return l.Check
}

// End of PsyCheckboxToolbar

// Entry PsyEntryToolbar
type PsyEntryToolbar interface {
	widget.ToolbarItem
	SetBinding(bind binding.String)
	OnChange(fn func(str string))
	Resize(size fyne.Size)
	Filter(re string)
	SetText(txt string)
	Get() *widget.Entry
}
type customEntryToolbarItem struct {
	*widget.Entry
}

func (l *customEntryToolbarItem) ToolbarObject() fyne.CanvasObject {
	return container.NewMax(l.Entry)

}
func NewCustomToolbarEntryItem() PsyEntryToolbar {
	e := widget.NewEntry()
	return &customEntryToolbarItem{
		e,
	}
}
func (l *customEntryToolbarItem) SetBinding(bind binding.String) {
	l.Entry.Bind(bind)
}
func (l *customEntryToolbarItem) OnChange(fn func(str string)) {
	l.Entry.OnChanged = fn
}
func (l *customEntryToolbarItem) Resize(size fyne.Size) {
	oSize := l.Entry.Size()
	log.Println("<Old Size ", oSize, ">")
	l.Entry.Resize(size)
	l.Entry.Refresh()
	log.Println("New Size <", l.Entry.Size(), ">")
	/*if size.Width < 60 {
		nSize := fyne.NewSize(60, oSize.Height)
		l.Entry.Resize(nSize)
	} else {
		nSize := fyne.NewSize(size.Width, oSize.Height)
		l.Entry.Resize(nSize)
		log.Println(l.Entry.Size())
	}*/
}
func (l *customEntryToolbarItem) Filter(re string) {
	l.Entry.Validator = fyne.StringValidator(
		func(s string) error {
			reg, err := regexp.Compile(re)
			if err != nil {
				log.Panicf("%+v", err)
			}
			l.Entry.SetValidationError(errors.New("Invalid Entry"))
			log.Printf("<String %s>", s)
			ok := reg.MatchString(s)
			if !ok {
				return errors.New("Invalid Entry")
			} else {
				return nil
			}

		},
	)
}
func (l *customEntryToolbarItem) SetText(txt string) {
	l.Entry.SetText(txt)
}
func (l *customEntryToolbarItem) Get() *widget.Entry {
	return l.Entry
}

// End of PsyEntryToolbar
