package tictactoe

const (
	emptySpace               = -1
	PlayerOne   PlayerNumber = 0
	PlayerTwo   PlayerNumber = 1
	HumanPlayer PlayerType   = iota + 1
)

type PlayerType int

type PlayerNumber int

type Player struct {
	Kind PlayerType
	Name string
}

type GameOption func(*Game) error

func InitialPlayer(n PlayerNumber) GameOption {
	return func(g *Game) error {
		return g.initialPlayer(n)
	}
}

func ConfigurePlayer(n PlayerNumber, t PlayerType, name string) GameOption {
	return func(g *Game) error {
		return g.configurePlayer(n, t, name)
	}
}

type Game struct {
	board         []int
	players       []Player
	currentPlayer PlayerNumber
}

func CreateGame(options ...GameOption) (*Game, error) {
	playerOne := Player{Kind: HumanPlayer, Name: "First Player"}
	playerTwo := Player{Kind: HumanPlayer, Name: "Second Player"}
	board := make([]int, 9, 9)
	for i := range board {
		board[i] = emptySpace
	}

	game := &Game{
		board:         board,
		players:       []Player{playerOne, playerTwo},
		currentPlayer: PlayerOne,
	}

	for _, option := range options {
		err := option(game)
		if err != nil {
			return nil, err
		}
	}

	return game, nil
}

func (g *Game) configurePlayer(n PlayerNumber, t PlayerType, name string) error {
	p := &g.players[n]
	p.Kind = t
	p.Name = name
	return nil
}

func (g *Game) initialPlayer(n PlayerNumber) error {
	g.currentPlayer = n
	return nil
}

func (g *Game) AvailableSpaces() []int {
	var available []int
	for i := 0; i < len(g.board); i++ {
		available = append(available, i+1)
	}
	return available
}

func (g *Game) CurrentPlayer() Player {
	return g.players[g.currentPlayer]
}
