package main

import (
    "testing"
    "fmt"
)

func TestGenerateGrid(t *testing.T) {
    our_grid := GenerateGrid()

    if fmt.Sprintf("%T", our_grid) != "grid.Grid" {
        t.Error("Generated grid is not typeof Grid")
    }
}
