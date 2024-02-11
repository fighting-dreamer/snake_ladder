package domain

import (
	"errors"
	"fmt"
)

type Board struct {
	Size    int // one side of the board 's value example for 10X10 the size is 10
	Snakes  map[int]int
	Ladders map[int]int
}

func NewBoard(size int) *Board {
	return &Board{
		Size:    size,
		Snakes:  map[int]int{},
		Ladders: map[int]int{},
	}
}

func (b *Board) IsValidPos(pos int) bool {
	return pos > 0 && pos <= b.Size*b.Size
}

func (b *Board) IsSnakeStart(pos int) bool {
	_, ok := b.Snakes[pos]
	return ok
}

func (b *Board) IsLadderStart(pos int) bool {
	_, ok := b.Ladders[pos]
	return ok
}

func (b *Board) SnakeEndPos(pos int) (int, error) {
	if b.IsSnakeStart(pos) {
		return b.Snakes[pos], nil
	}
	return -1, errors.New(fmt.Sprintf("Snake does not start at %d", pos))
}

func (b *Board) LadderEndPos(pos int) (int, error) {
	if b.IsLadderStart(pos) {
		return b.Ladders[pos], nil
	}
	return -1, errors.New(fmt.Sprintf("Ladder does not start at %d", pos))
}

func (b *Board) HaveReachedEnd(pos int) bool {
	return b.Size*b.Size == pos
}
