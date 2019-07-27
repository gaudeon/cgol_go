// cgol_go - John Conway's Game of Life in Go
package main

import (
	"fmt"
	"github.com/gaudeon/cgol_go/grid"
)

func main() {
	our_grid := GenerateGrid()

	println(fmt.Sprintf("%d x %d Grid Generated...", our_grid.Cols, our_grid.Rows))
}

func GenerateGrid() grid.Grid {
	return grid.GenerateRandomizedGrid(10, 10)
}
