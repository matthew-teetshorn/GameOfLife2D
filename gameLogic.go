package main

// nextGeneration calculates the next generation of cells based on specific rules
func nextGeneration(currGrid, nextGrid *[][]Cell) {
	for r, row := range *currGrid {
		for c := range row {
			count := countNeighbors(currGrid, r, c)
			currCell := (*currGrid)[r][c]
			(*nextGrid)[r][c].WasAlive = currCell.IsAlive
			if currCell.IsAlive {
				if count < 2 || count > 3 {
					(*nextGrid)[r][c].IsAlive = false
				} else {
					(*nextGrid)[r][c].IsAlive = true
				}
			} else if count == 3 {
				(*nextGrid)[r][c].IsAlive = true
			} else {
				(*nextGrid)[r][c].IsAlive = false
			}
		}
	}
}

func countNeighbors(gameGrid *[][]Cell, currRow, currCol int) int {
	count := 0

	for dR := -1; dR <= 1; dR++ {
		for dC := -1; dC <= 1; dC++ {
			if dR == 0 && dC == 0 {
				continue
			}
			nextRow := currRow + dR
			nextCol := currCol + dC
			if nextRow >= 0 && nextRow < Height/ColHeight && nextCol >= 0 && nextCol < Width/ColWidth {
				if (*gameGrid)[nextRow][nextCol].IsAlive {
					count++
				}
			}
		}
	}

	return count
}
