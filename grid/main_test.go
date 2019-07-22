package grid

import "testing"

var COLS uint32 = 10
var ROWS uint32 = 10
var SIZE int = int(COLS * ROWS)

func TestGenerateRandomizedGrid(t *testing.T) {
    our_grid := GenerateRandomizedGrid(COLS, ROWS)
    
    if our_grid.Cols != COLS {
        t.Error("Grid columns not properly defined")
    } 
    
    if our_grid.Rows != ROWS {
        t.Error("Grid rows not properly defined")
    }

    var sum uint32 = 0
    for i := 0; i < SIZE; i++ {
        sum += our_grid.Data[i]
    }

    if sum <= 0 {
        t.Error("Grid randomized to all zeros")
    }
}

func TestGetCellAtIndex(t *testing.T) {
    our_grid := GenerateRandomizedGrid(COLS, ROWS)

    all_values_correct := true

    for i := 0; i < SIZE; i++ {
        if our_grid.Data[i] != our_grid.GetCellAtIndex( uint32(i) ) {
            all_values_correct = false
        }
    }

    if ! all_values_correct {
        t.Error("grid.GetCellAtIndex values return did not always match what was in grid data")
    }
}
