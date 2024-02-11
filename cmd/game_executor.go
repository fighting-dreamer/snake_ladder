package cmd

import "nipun.io/snake_ladder/domain"

type IGameExecutor interface {
	Init() error
	Execute() error
	Status() (domain.Game, error)
}
