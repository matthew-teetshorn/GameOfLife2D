package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
	w := a.NewWindow("Conway's Game of Life")

	grid := createGrid()
	w.SetContent(grid)
	w.Resize(fyne.NewSize(600, 600))
	w.ShowAndRun()
}

func createGrid() *fyne.Container {
	grid := container.NewGridWithColumns(30)

	for y := range 30 {
		for x := range 30 {
			bg := canvas.NewRectangle(color.Gray{0x30})
			if x%2 == y%2 {
				bg.FillColor = color.Gray{0xE0}
			}

			grid.Add(bg)
		}
	}

	return grid
}
