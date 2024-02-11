package cmd

import (
	"fmt"

	"nipun.io/snake_ladder/domain"
	"nipun.io/snake_ladder/reader"
)

type LocalGameExecutor struct {
	game   *domain.Game
	reader reader.IReader
	Done   chan struct{}
}

func NewGame(reader reader.IReader, done chan struct{}) *LocalGameExecutor {
	return &LocalGameExecutor{
		reader: reader,
		Done:   done,
	}
}

func (lg *LocalGameExecutor) Init() error {
	game := domain.NewGame()
	boardSize := lg.reader.ReadInt()
	board := domain.NewBoard(boardSize)
	game.Board = board

	playersCnt := lg.reader.ReadInt()
	snakesCnt, ladderCnt := lg.reader.ReadInt(), lg.reader.ReadInt()
	for i := 0; i < playersCnt; i++ {
		playerName := lg.reader.ReadString()
		game.Players = append(game.Players, domain.Player{
			Name:       playerName,
			CurrentPos: 0,
		})
	}

	for i := 0; i < snakesCnt; i++ {
		start, end := lg.reader.ReadInt(), lg.reader.ReadInt()
		board.Snakes[start] = end
	}

	for i := 0; i < ladderCnt; i++ {
		start, end := lg.reader.ReadInt(), lg.reader.ReadInt()
		board.Ladders[start] = end
	}

	lg.game = game
	return nil
}

func (lg *LocalGameExecutor) Execute() error {
	// default player 1 takes first turn
	turnIndex := 0
	nextIndex := func(index int, playerCnt int) int {
		return (index + 1) % playerCnt
	}

	for {
		select {
		case <-lg.Done:
			break
		default:
			dice := lg.reader.ReadInt()
			newPos := lg.game.Players[turnIndex].CurrentPos + dice
			if !lg.game.Board.IsValidPos(newPos) {
				turnIndex = nextIndex(turnIndex, len(lg.game.Players))
				continue
			}

			if lg.game.Board.IsSnakeStart(newPos) {
				newPos, _ = lg.game.Board.SnakeEndPos(newPos)
			}
			if lg.game.Board.IsLadderStart(newPos) {
				newPos, _ = lg.game.Board.LadderEndPos(newPos)
			}

			lg.game.Players[turnIndex].CurrentPos = newPos
			if lg.game.Board.HaveReachedEnd(newPos) {
				break
			}
			turnIndex = nextIndex(turnIndex, len(lg.game.Players))
		}
		fmt.Println("broke from the select statement")
		break
	}

	// turnIndex points to the player that have reached end.
	return nil
}

func (lg *LocalGameExecutor) Status() (domain.Game, error) {
	return *lg.game, nil
}
