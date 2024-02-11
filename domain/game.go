package domain

type Game struct {
	Board   *Board
	Players []Player
}

func NewGame() *Game {
	return &Game{}
}
