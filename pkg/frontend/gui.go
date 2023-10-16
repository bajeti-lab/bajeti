package frontend

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func MakeSidebar() fyne.CanvasObject {
	sidebar := container.NewVBox(
		widget.NewToolbar(
			widget.NewToolbarSpacer(),
			widget.NewToolbarAction(theme.MenuIcon(), func() {}),
		),
		widget.NewButton("Home", func() {}),
		widget.NewButton("Budgets", func() {}),
		widget.NewButton("Transactions", func() {}),
		widget.NewButton("Reports", func() {}),
		widget.NewButton("Settings", func() {}),
	)

	return sidebar

}

func MakeContent() fyne.CanvasObject {
	logo := canvas.NewImageFromResource(resourceLogoPng)
	logo.FillMode = canvas.ImageFillContain
	logo.SetMinSize(fyne.NewSize(45, 45))

	logoGrid := container.New(
		layout.NewGridLayout(4),
		logo,
	)

	salutationSection := container.NewVBox(
		widget.NewLabel("Hi,"),
	)

	searchSection := container.NewHBox(
		MakeSearch(),
	)
	searchSection.Position()

	content := container.NewVBox(
		// content := canvas.NewRectangle(color.Gray{Y: 0xee})
		logoGrid,
		salutationSection,
		searchSection,
	)

	return content

}

func MakeGUI() fyne.CanvasObject {
	left := MakeSidebar()
	content := MakeContent()

	objs := []fyne.CanvasObject{left, content}

	return container.New(newBajetiLayout(left, content), objs...)
}
