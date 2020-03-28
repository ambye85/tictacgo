package tictactoe_test

import (
	ttt "github.com/ambye85/tictacgo/internal/app/tictactoe"
	"reflect"
	"testing"
)

func TestNewGameHas9AvailableSpaces(t *testing.T) {
	game, _ := ttt.CreateGame()

	expectedSpaces := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if availableSpaces := game.AvailableSpaces(); !reflect.DeepEqual(availableSpaces, expectedSpaces) {
		t.Errorf("Expected available spaces to be %v, got %v", expectedSpaces, availableSpaces)
	}
}

func TestNewGameDefaultsToFirstPlayerGoesFirst(t *testing.T) {
	game, _ := ttt.CreateGame()

	if currentPlayer := game.CurrentPlayer(); currentPlayer.Name != "First Player" {
		t.Errorf("Expected current player name to be %s, got %s", "First Player", currentPlayer.Name)
	}
}

func TestNewGameConfigSecondPlayerGoesFirst(t *testing.T) {
	game, _ := ttt.CreateGame(ttt.InitialPlayer(ttt.PlayerTwo))

	if currentPlayer := game.CurrentPlayer(); currentPlayer.Name != "Second Player" {
		t.Errorf("Expected current player name to be %s, got %s", "Second Player", currentPlayer.Name)
	}
}

func TestNewGameConfigChangeFirstPlayersName(t *testing.T) {
	game, _ := ttt.CreateGame(ttt.ConfigurePlayer(ttt.PlayerOne, ttt.HumanPlayer, "Player 1"))

	if currentPlayer := game.CurrentPlayer(); currentPlayer.Name != "Player 1" {
		t.Errorf("Expected current player name to be %s, got %s", "Player 1", currentPlayer.Name)
	}
}

func TestNewGameConfigChangeSecondPlayersName(t *testing.T) {
	game, _ := ttt.CreateGame(
		ttt.InitialPlayer(ttt.PlayerTwo),
		ttt.ConfigurePlayer(ttt.PlayerTwo, ttt.HumanPlayer, "Player 2"),
	)

	if currentPlayer := game.CurrentPlayer(); currentPlayer.Name != "Player 2" {
		t.Errorf("Expected current player name to be %s, got %s", "Player 2", currentPlayer.Name)
	}
}
