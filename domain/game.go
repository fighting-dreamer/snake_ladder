package domain

import "math/rand"

type Game struct {
	board   Board
	players []Player
}

func makePlayers(names []string, count int) []Player {
	players := []Player{}
	for i := 0; i < count; i++ {
		players = append(players, Player{
			Name:       names[i],
			CurrentPos: 0,
		})
	}

	return players
}

func NewGame(size int, playerCount int, playerNames []string, snakes [][2]int, ladders [][2]int) Game {
	players := makePlayers(playerNames, playerCount)
	return Game{
		board:   NewBoard(size, snakes, ladders),
		players: players,
	}
}

func (g *Game) isValidPos(pos int) bool {
	if pos >= 0 && pos <= g.board.Size {
		return true
	}

	return false
}

func (g *Game) move(playerID int) (retryflag bool, nextPos int) {
	diceRoll := 1 + rand.Intn(6)
	if diceRoll == 6 {
		retryflag = true
	}
	currPos := g.players[playerID].CurrentPos
	nextPos = currPos + diceRoll

	if !g.isValidPos(nextPos) {
		return false, currPos
	}

	if tail, ok := g.board.Snakes[nextPos]; ok {
		nextPos = tail
	}
	if upPos, ok := g.board.Ladders[nextPos]; ok {
		nextPos = upPos
	}

	return
}
