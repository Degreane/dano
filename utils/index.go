package utils

import (
	"github.com/degreane/dano/utils/gui"
	"github.com/degreane/dano/utils/names"
)

var (
	GetName func(in int, prefix string) string = names.GetName
	InitGui func()                             = gui.InitGUI
	RunGUI  func()                             = gui.RunGUI
	// RunFyne func()                             = guifyne.RunGUI
)
