package frontend

import "fyne.io/fyne/v2"

const sidewidth = 220
const WindowHeight = 1024
const WindowWidth = 768
const topHeight = 0

type bajetiLayout struct {
	left, content fyne.CanvasObject
}

func newBajetiLayout(left, content fyne.CanvasObject) fyne.Layout {
	return &bajetiLayout{left: left, content: content}

}

func (l *bajetiLayout) Layout(_ []fyne.CanvasObject, size fyne.Size) {

	l.left.Move(fyne.NewPos(0, topHeight))
	l.left.Resize(fyne.NewSize(sidewidth, size.Height-topHeight))

	l.content.Move(fyne.NewPos(sidewidth, topHeight))
	l.content.Resize(fyne.NewSize(size.Width-sidewidth, size.Height-topHeight))
}

func (l *bajetiLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	leftWidth := l.left.MinSize().Width
	contentHeight := l.content.MinSize().Height

	return fyne.NewSize(leftWidth+contentHeight, topHeight+contentHeight)
}
