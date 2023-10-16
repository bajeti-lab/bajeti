package frontend

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func MakeTitleBar() fyne.CanvasObject {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.MenuExpandIcon(), func() {}),
	)

	logo := canvas.NewImageFromResource(resourceLogoPng)
	logo.FillMode = canvas.ImageFillContain
	logo.SetMinSize(fyne.NewSize(70, 70))

	return container.NewStack(toolbar, logo)

}

func MakeGUI() fyne.CanvasObject {
	left := widget.NewLabel("Left")
	right := widget.NewLabel("Right")

	content := widget.NewLabel("Content")
	content.Wrapping = fyne.TextWrapWord
	content.Alignment = fyne.TextAlignCenter

	return container.NewBorder(MakeTitleBar(), nil, right, left, content)
}
