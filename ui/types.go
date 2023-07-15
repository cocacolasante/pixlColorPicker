package ui

import (
	"fyne.io/fyne/v2"
	"pixl.io/apptype"
	"pixl.io/swatch"
)

type AppInit struct {
	PixlWindow fyne.Window
	State *apptype.State
	Swatches []*swatch.Swatch
}