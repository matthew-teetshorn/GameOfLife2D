package main

import (
	"image"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

const (
	Width           = 600
	Height          = 600
	ColWidth        = 30
	ColHeight       = 30
	GenLengthMillis = 500
)

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
	gridLineColor := color.RGBA{R: 0, G: 125, B: 125, A: 255}
	seeThrough := color.RGBA{R: 0, G: 0, B: 0, A: 0}

	a := app.New()
	w := a.NewWindow("Conway's Game of Life")

	// The grid ([][]Cell) slice that contains the actual game grid
	gameGrid := createGrid(Width/ColWidth, Height/ColHeight)

	// The boundary box for our grid image
	gridRect := image.Rect(0, 0, Width+1, Height+1)
	// The image that we draw the game cells onto as pixels
	gridImage := image.NewNRGBA(gridRect)
	updateImageGrid(&gameGrid, gridImage, aliveColor, deadColor, ColWidth, ColHeight)

	// Create an image overlay for the grid lines
	gridLinesImage := image.NewNRGBA(gridRect)
	createGridLines(gridLinesImage, gridLineColor, seeThrough, Width+1, Height+1, ColWidth, ColHeight)

	// Fyne cavnas element for containing grid lines
	canvasGridLines := canvas.NewImageFromImage(gridLinesImage)
	canvasGridLines.FillMode = canvas.ImageFillContain
	// Fyne canvas element for containing the image
	canvasImage := canvas.NewImageFromImage(gridImage)
	// Maintain aspect ratio of image on window resize
	canvasImage.FillMode = canvas.ImageFillContain
	// Fyne container for holding all items for our grid
	gridContainer := container.New(layout.NewStackLayout(), canvasImage, canvasGridLines)
	// Creating an overlay region to handle mouse click events
	// We are creating a closure over the anonymous function so it has gridContainer in scope
	overlayWidget := NewInvisibleButton(func(pe *fyne.PointEvent) {
		// Get the current size of the container
		cWidth := gridContainer.Size().Width
		cHeight := gridContainer.Size().Height
		mouseX, mouseY := pe.Position.X, pe.Position.Y
		r, c, ok := pixelToGridSquare(mouseX, mouseY, cWidth, cHeight)
		if ok {
			gameGrid[r][c].WasAlive = gameGrid[r][c].IsAlive
			gameGrid[r][c].IsAlive = !gameGrid[r][c].IsAlive
			updateImageGrid(&gameGrid, gridImage, aliveColor, deadColor, ColWidth, ColHeight)
			canvasImage.Refresh()
		}
	})

	gridContainer.Add(overlayWidget)

	// Spawn go routine that handles the game update tasks on a time tick
	go func() {
		ticker := time.NewTicker(GenLengthMillis * time.Millisecond)
		defer ticker.Stop()

		for range ticker.C {
			nextGeneration(&gameGrid)
			fyne.Do(func() {
				updateImageGrid(&gameGrid, gridImage, aliveColor, deadColor, ColWidth, ColHeight)
				canvasImage.Refresh()
			})
		}
	}()

	w.SetContent(gridContainer)
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

func createGridLines(img ChangeableImage, lColor, cColor color.Color, sWidth, sHeight, cWidth, cHeight int) {
	for x := range sWidth {
		for y := range sHeight {
			if x%cWidth == 0 || y%cHeight == 0 {
				img.Set(x, y, lColor)
			} else {
				img.Set(x, y, cColor)
			}
		}
	}
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
			(*gameGrid)[r][c].WasAlive = (*gameGrid)[r][c].WasAlive
			(*gameGrid)[r][c].IsAlive = (*gameGrid)[r][c].IsAlive
			// (*gameGrid)[r][c].WasAlive = (*gameGrid)[r][c].IsAlive
			// (*gameGrid)[r][c].IsAlive = !(*gameGrid)[r][c].IsAlive
		}
	}
}

func pixelToGridSquare(x, y, sWidth, sHeight float32) (r, c int, isValid bool) {
	// Grid is square and contained within the lesser of sWidth and sHeight
	dimension := min(sWidth, sHeight)
	isValid = true
	margin := float32(0.0)
	scale := float32(0.0)

	if dimension == sWidth {
		// If we're outside the grid it's invalid
		margin = (sHeight - dimension) / 2
		if y < margin || y > sHeight-margin {
			return 0, 0, false
		}
		scale = sWidth / Width
		c = int(x / (ColWidth * scale))
		r = int((y - margin) / (ColHeight * scale))
	} else {
		// If we're outside the grid it's invalid
		margin = (sWidth - dimension) / 2
		if x < margin || x > sWidth-margin {
			return 0, 0, false
		}
		scale = sHeight / Height
		c = int((x - margin) / (ColWidth * scale))
		r = int(y / (ColHeight * scale))
	}

	return
}
