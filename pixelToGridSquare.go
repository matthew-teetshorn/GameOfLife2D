package main

// pixelToGridSquare takes an x, y mouse coordinate and a screen width and height
// it returns row and column indices of the corresponding click corresponding
// isValid is used to determine if the user clicked outside the valid gameGrid
//
// Due to the way Fyne handles containers and resizing there is a likelihood that the container
// in which the game grid renders will have extra "margin" space on one side or the other:
//
// ++++++++++++++++++++++++++++
// +        |        |        +
// +        |        |        +
// + margin |  Grid  | margin +
// +        |        |        +
// +        |        |        +
// ++++++++++++++++++++++++++++
//
// ++++++++++
// +        |
// + margin |
// +________|
// +        |
// +        |
// +  Grid  |
// +        |
// +________|
// +        |
// + margin |
// +        |
// ++++++++++
//
// We are determining the margin area mathematically as a function of which screen dimension
// is constraining the size of the game grid and using the ratio of the grid dimensions to scale
func pixelToGridSquare(x, y, sWidth, sHeight float32) (r, c int, isValid bool) {
	sRatio := sWidth / sHeight
	gRatio := float32(gWidth / gHeight)

	isValid = true
	var margin float32
	var scale float32
	var actualGridWidth float32
	var actualGridHeight float32

	if sRatio < gRatio {
		actualGridWidth = sWidth
		actualGridHeight = actualGridWidth / gRatio

		margin = (sHeight - actualGridHeight) / 2

		// If we're outside the grid it's invalid
		if y < margin || y > sHeight-margin {
			return 0, 0, false
		}

		scale = sWidth / gWidth
		c = int(x / (ColWidth * scale))
		r = int((y - margin) / (ColHeight * scale))
	} else {
		actualGridHeight = sHeight
		actualGridWidth = actualGridHeight * gRatio

		margin = (sWidth - actualGridWidth) / 2

		// If we're outside the grid it's invalid
		if x < margin || x > sWidth-margin {
			return 0, 0, false
		}
		scale = sHeight / gHeight
		c = int((x - margin) / (ColWidth * scale))
		r = int(y / (ColHeight * scale))
	}

	return
}
