package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

// A custom invisible widget that can be tapped.
type invisibleButton struct {
	widget.BaseWidget
	onTapped func(pe *fyne.PointEvent)
}

// Ensure the type implements fyne.Tappable
// var _ fyne.Tappable = (*invisibleButton)(nil)

func NewInvisibleButton(onTapped func(*fyne.PointEvent)) *invisibleButton {
	b := &invisibleButton{
		onTapped: onTapped,
	}
	b.ExtendBaseWidget(b)
	return b
}

func (b *invisibleButton) CreateRenderer() fyne.WidgetRenderer {
	// The renderer uses a completely transparent rectangle for the visual
	rect := canvas.NewRectangle(color.Transparent)

	return widget.NewSimpleRenderer(rect)
}

func (b *invisibleButton) Tapped(pe *fyne.PointEvent) {
	if b.onTapped != nil {
		b.onTapped(pe)
	}
}
