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
