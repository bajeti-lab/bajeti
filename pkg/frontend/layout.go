package frontend

import "fyne.io/fyne/v2"

const WindowHeight = 768
const WindowWidth = 768
const bottomHeight = 360

type bajetiLayout struct {
	content fyne.CanvasObject
	bottom  fyne.CanvasObject
}

func newBajetiLayout(content fyne.CanvasObject, bottom fyne.CanvasObject) fyne.Layout {
	return &bajetiLayout{content: content, bottom: bottom}

}

func (l *bajetiLayout) Layout(_ []fyne.CanvasObject, size fyne.Size) {
	l.content.Resize(fyne.NewSize(size.Width, size.Height-bottomHeight))
	l.content.Move(fyne.NewPos(0, 0))

	l.bottom.Resize(fyne.NewSize(size.Width, bottomHeight))
	l.bottom.Move(fyne.NewPos(0, size.Height-bottomHeight))

}

func (l *bajetiLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(WindowWidth, WindowHeight)

}
