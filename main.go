package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/bajeti-lab/bajeti/pkg/frontend"
)

func main() {
	// create the database
	// db, err := utils.CreateDatabase()

	// if err != nil {
	// 	panic("failed to connect database")
	// }

	// create GUI for the user to interact with using fyne
	a := app.New()
	a.Settings().SetTheme(frontend.NewBajetiTheme())
	w := a.NewWindow("Bajeti")
	w.Resize(fyne.NewSize(frontend.WindowWidth, frontend.WindowHeight))

	w.SetContent(frontend.MakeGUI())
	w.ShowAndRun()

}
