package integration_test

import (
	"github.com/ambye85/tictacgo/internal/app/tictactoe"
	"reflect"
	"testing"
)

func TestCreateNewGame(t *testing.T) {
	config := tictactoe.Config{
		FirstPlayer:   2,
		BoardSize:     3,
		PlayerOneType: tictactoe.HumanPlayer,
		PlayerOneName: "Player 1",
		PlayerTwoType: tictactoe.HumanPlayer,
		PlayerTwoName: "Player 2",
	}

	game := tictactoe.CreateGame(config)

	expectedSpaces := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if availableSpaces := game.AvailableSpaces(); !reflect.DeepEqual(availableSpaces, expectedSpaces) {
		t.Errorf("Expected available spaces to be %v, got %v.", expectedSpaces, availableSpaces)
	}

	expectedPlayer := tictactoe.Player{
		Kind: tictactoe.HumanPlayer,
		Name: "Player 2",
	}
	if currentPlayer := game.CurrentPlayer(); currentPlayer != expectedPlayer {
		t.Errorf("Expected current player name to be %s, got %s", expectedPlayer.Name, currentPlayer.Name)
	}
}
