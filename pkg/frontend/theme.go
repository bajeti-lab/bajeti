//go:generate fyne bundle -o bundled.go -package frontend assets

package frontend

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type BajetiTheme struct {
	fyne.Theme
}

func NewBajetiTheme() fyne.Theme {
	return &BajetiTheme{Theme: theme.DefaultTheme()}
}

func (t *BajetiTheme) Color(name fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	return t.Theme.Color(name, theme.VariantLight)
}

func (t *BajetiTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNameText {
		return 12
	}

	return t.Theme.Size(name)
}

func (t *BajetiTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return t.Theme.Icon(name)
}
