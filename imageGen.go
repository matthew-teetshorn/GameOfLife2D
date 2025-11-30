package main

import "image/color"

// Creates an overlay image with visible gridlines and see through squares
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
