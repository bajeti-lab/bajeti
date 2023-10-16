package frontend

import (
	"fyne.io/fyne/v2"
	"fyne.io/x/fyne/widget"
)

func MakeSearch() fyne.CanvasObject {
	xentry := widget.NewCompletionEntry([]string{"Search"})

	xentry.OnChanged = func(s string) {
		// do something with the new value of s
	}
	return xentry
}
