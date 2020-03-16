package integration_test

import (
	"github.com/ambye85/tictacgo/internal/app/tictactoe"
	"testing"
)

func TestCreatesBoardWithAllEmptySpaces(t *testing.T) {
	board := tictactoe.NewBoard()
	if freeSpaces := board.FreeSpaces(); freeSpaces != 9 {
		t.Errorf("Expected 9 free spaces, got %d", freeSpaces)
	}
}