// cgol_go - unit testing for main package
package main

import (
	"fmt"
	"testing"
)

func TestGenerateGrid(t *testing.T) {
	our_grid := GenerateGrid()

	if fmt.Sprintf("%T", our_grid) != "grid.Grid" {
		t.Error("Generated grid is not typeof Grid")
	}
}
