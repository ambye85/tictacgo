package integration_test

import (
	ttt "github.com/ambye85/tictacgo/internal/app/tictactoe"
	"reflect"
	"testing"
)

func TestCreateNewGame(t *testing.T) {
	game, err := ttt.CreateGame(
		ttt.InitialPlayer(ttt.PlayerTwo),
		ttt.ConfigurePlayer(ttt.PlayerOne, ttt.HumanPlayer, "Player 1"),
		ttt.ConfigurePlayer(ttt.PlayerTwo, ttt.HumanPlayer, "Player 2"),
	)
	if err != nil {
		t.Fatalf("Error creating game:\n%+v", err)
	}

	expectedSpaces := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if availableSpaces := game.AvailableSpaces(); !reflect.DeepEqual(availableSpaces, expectedSpaces) {
		t.Errorf("Expected available spaces to be %v, got %v.", expectedSpaces, availableSpaces)
	}

	expectedPlayer := ttt.Player{
		Kind: ttt.HumanPlayer,
		Name: "Player 2",
	}
	if currentPlayer := game.CurrentPlayer(); currentPlayer != expectedPlayer {
		t.Errorf("Expected current player name to be %s, got %s", expectedPlayer.Name, currentPlayer.Name)
	}
}
