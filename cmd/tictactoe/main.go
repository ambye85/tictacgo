package main

import (
	"fmt"
	"github.com/ambye85/tictacgo/internal/app/tictactoe"
	"os"
)

func main() {
	game, err := tictactoe.CreateGame()
	if err != nil {
		fmt.Println("Error creating game, exiting...")
		os.Exit(1)
	}
	fmt.Printf("%+v\n", game)
}
