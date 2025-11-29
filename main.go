package main

import (
	"image"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

const (
	Width     = 600
	Height    = 600
	ColWidth  = 30
	ColHeight = 30
)

// ex: grab a timer for updating in a concurrent go routine
// go func() {
// 	count := 0
// 	for range time.Tick(time.Second) {
// 		count++
// 	}
// }()

// ChangeableImage defines an interface for image types that can be modified using Set().
type ChangeableImage interface {
	image.Image // Embed image.Image to include its methods (Bounds, ColorModel, At)
	Set(x, y int, c color.Color)
}

// Cell keeps track of wheter cell has been born/died between generations
type Cell struct {
	WasAlive bool
	IsAlive  bool
}

func main() {
	aliveColor := color.RGBA{R: 200, G: 200, B: 200, A: 255}
	deadColor := color.RGBA{R: 20, G: 20, B: 20, A: 255}

	a := app.New()
	w := a.NewWindow("Conway's Game of Life")

	gameGrid := createGrid(Width/ColWidth, Height/ColHeight)
	gridRect := image.Rect(0, 0, Width, Height)
	gridImage := image.NewNRGBA(gridRect)
	gridRaster := canvas.NewRasterFromImage(gridImage)
	updateImageGrid(&gameGrid, gridImage, aliveColor, deadColor, ColWidth, ColHeight)

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			nextGeneration(&gameGrid)
			fyne.Do(func() {
				updateImageGrid(&gameGrid, gridImage, aliveColor, deadColor, ColWidth, ColHeight)
				gridRaster.Refresh()
			})
		}
	}()

	w.SetContent(gridRaster)
	w.Resize(fyne.NewSize(Width, Height))
	w.Show()
	a.Run()
}

// createGrid creates the starting grid structure
func createGrid(rows, cols int) [][]Cell {
	grid := make([][]Cell, rows)

	for i := range grid {
		grid[i] = make([]Cell, cols)
	}

	for x := range rows {
		for y := range cols {
			grid[x][y] = Cell{WasAlive: true, IsAlive: false}
		}
	}

	return grid
}

// updateImageGrid takes an image grid and updates the cell's pixel colors for cells which have been
// born/died since the previous generation
func updateImageGrid(gameGrid *[][]Cell, img ChangeableImage, alive, dead color.Color, cWidth, cHeight int) {
	for r, row := range *gameGrid {
		for c, cell := range row {
			if cell.IsAlive != cell.WasAlive {
				colStart := cWidth * c
				rowStart := cHeight * r
				for x := colStart; x < colStart+cWidth; x++ {
					for y := rowStart; y < rowStart+cHeight; y++ {
						if cell.IsAlive {
							img.Set(x, y, alive)
						} else {
							img.Set(x, y, dead)
						}
					}
				}
			}
		}
	}
}

// nextGeneration calculates the next generation of cells based on specific rules
func nextGeneration(gameGrid *[][]Cell) {
	for r, row := range *gameGrid {
		for c := range row {
			(*gameGrid)[r][c].WasAlive = (*gameGrid)[r][c].IsAlive
			(*gameGrid)[r][c].IsAlive = !(*gameGrid)[r][c].IsAlive
		}
	}
}
