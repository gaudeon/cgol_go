// grid - responsible for grid data and grid helper functions
package grid

import "math/rand"
import "time"

const MAX_STATE_VALUE uint32 = 1
const (
	Dead uint32 = iota
	Alive
)

type Grid struct {
	Cols uint32
	Rows uint32
	Data []uint32
}

func GenerateRandomizedGrid(cols uint32, rows uint32) Grid {
	grid_size := int(cols * rows)

	grid := Grid{Cols: cols, Rows: rows, Data: make([]uint32, grid_size)}

	rand.Seed(time.Now().Unix())

	for i := 0; i < grid_size; i++ {
		grid.Data[i] = rand.Uint32() % (MAX_STATE_VALUE + 1)
	}

	return grid
}

func (g Grid) GetCellAtIndex(index uint32) uint32 {
	var value uint32 = Dead

	if 0 <= index && index < g.Rows*g.Cols {
		value = g.Data[index]
	}

	return value
}

func (g Grid) GetCellAtXY(x uint32, y uint32) uint32 {
	var value uint32 = Dead

	if 0 <= x && x < g.Cols && 0 <= y && y < g.Rows {
		value = g.Data[x+y*g.Rows]
	}

	return value
}

func (g Grid) SetCellAtIndex(index uint32, value uint32) {
	if 0 <= index && index < g.Rows*g.Cols {
		g.Data[index] = value
	}
}

func (g Grid) SetCellAtXY(x uint32, y uint32, value uint32) {
	if 0 <= x && x < g.Cols && 0 <= y && y < g.Rows {
		g.Data[x+y*g.Rows] = value
	}
}
