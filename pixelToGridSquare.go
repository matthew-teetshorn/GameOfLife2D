package main

// pixelToGridSquare takes an x, y mouse coordinate and a screen width and height
// it returns row and column indices of the corresponding click corresponding
// isValid is used to determine if the user clicked outside the valid gameGrid
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
