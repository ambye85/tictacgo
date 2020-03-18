package tictactoe

type Board struct {
	freeSpaces int
}

func NewBoard() Board {
	return Board{
		freeSpaces: 9,
	}
}

func (b *Board) FreeSpaces() int {
	return b.freeSpaces
}

type PlayerType int

type Player struct {
	Kind PlayerType
	Name string
}

const emptySpace = -1

const (
	HumanPlayer = iota + 1
)

type Config struct {
	FirstPlayer   int
	BoardSize     int
	PlayerOneType PlayerType
	PlayerOneName string
	PlayerTwoType PlayerType
	PlayerTwoName string
}

type Game struct {
	board []int
	currentPlayer Player
}

func CreateGame(config Config) Game {
	if config.BoardSize == 0 {
		config.BoardSize = 3
	}

	playerOne := Player{}
	playerOne.Kind = config.PlayerOneType
	if config.PlayerOneName == "" {
		playerOne.Name = "First Player"
	} else {
		playerOne.Name = config.PlayerOneName
	}

	playerTwo := Player{}
	playerTwo.Kind = config.PlayerTwoType
	if config.PlayerTwoName == "" {
		playerTwo.Name = "Second Player"
	} else {
		playerTwo.Name = config.PlayerTwoName
	}

	initialPlayer := playerOne
	if config.FirstPlayer == 2 {
		initialPlayer = playerTwo
	}

	var board []int
	for i := 0; i < config.BoardSize * config.BoardSize; i++ {
		board = append(board, emptySpace)
	}

	return Game{
		board: board,
		currentPlayer: initialPlayer,
	}
}

func (g *Game) AvailableSpaces() []int {
	var available []int
	for i := 0; i < len(g.board); i++ {
		available = append(available, i+1)
	}
	return available
}

func (g *Game) Board() []int {
	return g.board
}

func (g *Game) CurrentPlayer() Player {
	return g.currentPlayer
}
