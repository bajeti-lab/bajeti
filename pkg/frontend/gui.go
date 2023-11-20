package frontend

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func MakeBottomBar() fyne.CanvasObject {
	bottomBar := container.NewAppTabs(
		container.NewTabItemWithIcon("Budget", theme.ContentAddIcon(), widget.NewLabel("Budget")),
		container.NewTabItemWithIcon("Expenses", theme.ContentAddIcon(), widget.NewLabel("Expenses")),
		container.NewTabItemWithIcon("Reports", theme.ContentAddIcon(), widget.NewLabel("Reports")),
		container.NewTabItemWithIcon("Settings", theme.SettingsIcon(), widget.NewLabel("Settings")),
	)

	return bottomBar
}

func MakeContent() fyne.CanvasObject {
	// Create a grid layout with 4 columns
	// contentGrid := layout.NewAdaptiveGridLayout(4)
	// Get the logo resource
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
	content := MakeContent()
	bottom := MakeBottomBar()

	objs := []fyne.CanvasObject{content, bottom}

	return container.New(newBajetiLayout(content, bottom), objs...)
}
