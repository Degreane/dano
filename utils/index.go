package utils

import (
	"github.com/degreane/dano/utils/gui"
	"github.com/degreane/dano/utils/guifyne"
)

var (
	GetName func(in int, prefix string) string = getName
	InitGui func()                             = gui.InitGUI
	RunGUI  func()                             = gui.RunGUI
	RunFyne func()                             = guifyne.RunGUI
)
