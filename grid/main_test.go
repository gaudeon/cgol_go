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

func TestGetCellAtXY(t *testing.T) {
    our_grid := GenerateRandomizedGrid(COLS, ROWS)

    all_values_correct := true

    for x := 0; x < int(COLS); x++ {
        for y := 0; y < int(ROWS); y++ {
            i := x + y * int(ROWS);

            if our_grid.Data[i] != our_grid.GetCellAtXY( uint32(x), uint32(y) ) {
                all_values_correct = false
            }
        }
    }

    if ! all_values_correct {
        t.Error("grid.GetCellAtXY values return did not always match what was in grid data")
    }
}

func TestSetCellAtIndex(t *testing.T) {
    our_grid := GenerateRandomizedGrid(COLS, ROWS)

    our_grid.SetCellAtIndex(10, 2)

    if our_grid.GetCellAtIndex(10) != 2 {
        t.Error("grid.SetCellAtIndex failed to set the correct value")
    }
}

func TestSetCellAtXY(t *testing.T) {
    our_grid := GenerateRandomizedGrid(COLS, ROWS)

    our_grid.SetCellAtXY(5, 5, 2)

    if our_grid.GetCellAtXY(5, 5) != 2 {
        t.Error("grid.SetCellAtXY failed to set the correct value")
    }
}
