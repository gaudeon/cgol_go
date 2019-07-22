package grid

import "math/rand"
import "time"

var MAX_STATE_VALUE uint32 = 2

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
        grid.Data[i] = rand.Uint32() % MAX_STATE_VALUE
    }

    return grid
}

func (g Grid) GetCellAtIndex(index uint32) (uint32){
    var value uint32 = 0

    if (0 <= index && index < g.Rows * g.Cols) {
        value = g.Data[index]
    }

    return value
}
