package tictactoe_test

import (
	"github.com/ambye85/tictacgo/internal/app/tictactoe"
	"reflect"
	"testing"
)

func TestEmptyBoardHas9FreeSpaces(t *testing.T) {
	board := tictactoe.NewBoard()
	if freeSpaces := board.FreeSpaces(); freeSpaces != 9 {
		t.Errorf("Expected 9 free spaces, got %d", freeSpaces)
	}
}

func TestNewGameHas9AvailableSpaces(t *testing.T) {
	config := tictactoe.Config{}
	game := tictactoe.CreateGame(config)

	expectedSpaces := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if availableSpaces := game.AvailableSpaces(); !reflect.DeepEqual(availableSpaces, expectedSpaces) {
		t.Errorf("Expected available spaces to be %v, got %v", expectedSpaces, availableSpaces)
	}
}

func TestNewGameDefaultsTo3x3Board(t *testing.T) {
	config := tictactoe.Config{}
	game := tictactoe.CreateGame(config)

	if totalSpaces := len(game.Board()); totalSpaces != 9 {
		t.Errorf("Expected total spaces to be %d, got %d", 9, totalSpaces)
	}
}

func TestNewGameDefaultsToFirstPlayerGoesFirst(t *testing.T) {
	config := tictactoe.Config{}
	game := tictactoe.CreateGame(config)

	if currentPlayer := game.CurrentPlayer(); currentPlayer.Name != "First Player" {
		t.Errorf("Expected current player name to be %s, got %s", "First Player", currentPlayer.Name)
	}
}

func TestNewGameConfigSecondPlayerGoesFirst(t *testing.T) {
	config := tictactoe.Config{FirstPlayer: 2}
	game := tictactoe.CreateGame(config)

	if currentPlayer := game.CurrentPlayer(); currentPlayer.Name != "Second Player" {
		t.Errorf("Expected current player name to be %s, got %s", "Second Player", currentPlayer.Name)
	}
}
