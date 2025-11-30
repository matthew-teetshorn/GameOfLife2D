package main

import "fmt"

// pixelToGridSquare takes an x, y mouse coordinate and a screen width and height
// it returns row and column indices of the corresponding click corresponding
// isValid is used to determine if the user clicked outside the valid gameGrid
func pixelToGridSquare(x, y, sWidth, sHeight float32) (r, c int, isValid bool) {
	// Grid is square and contained within the lesser of sWidth and sHeight
	// dimension := min(sWidth, sHeight)
	sRatio := sWidth / sHeight
	gRatio := float32(gWidth / gHeight)

	fmt.Printf("Screen Ratio: %.02f, Grid Ratio: %.02f\n", sRatio, gRatio)

	isValid = true
	margin := float32(0.0)
	scale := float32(0.0)

	if sRatio < gRatio {
		actualGridWidth := sWidth
		actualGridHeight := actualGridWidth / gRatio

		margin = (sHeight - actualGridHeight) / 2

		// If we're outside the grid it's invalid
		if y < margin || y > sHeight-margin {
			return 0, 0, false
		}

		scale = sWidth / gWidth
		c = int(x / (ColWidth * scale))
		r = int((y - margin) / (ColHeight * scale))
	} else {
		actualGridHeight := sHeight
		actualGridWidth := actualGridHeight * gRatio

		margin = (sWidth - actualGridWidth) / 2

		// If we're outside the grid it's invalid
		if x < margin || x > sWidth-margin {
			return 0, 0, false
		}
		scale = sHeight / gHeight
		c = int((x - margin) / (ColWidth * scale))
		r = int(y / (ColHeight * scale))
	}

	// if dimension == sWidth {
	// 	// If we're outside the grid it's invalid
	// 	margin = (sHeight - dimension) / 2
	// 	if y < margin || y > sHeight-margin {
	// 		return 0, 0, false
	// 	}
	// 	scale = sWidth / gWidth
	// 	c = int(x / (ColWidth * scale))
	// 	r = int((y - margin) / (ColHeight * scale))
	// } else {
	// 	// If we're outside the grid it's invalid
	// 	margin = (sWidth - dimension) / 2
	// 	if x < margin || x > sWidth-margin {
	// 		return 0, 0, false
	// 	}
	// 	scale = sHeight / gHeight
	// 	c = int((x - margin) / (ColWidth * scale))
	// 	r = int(y / (ColHeight * scale))
	// }

	return
}
