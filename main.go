package main

import (
	"nipun.io/snake_ladder/cmd"
)

func main() {
	// executor := cmd.NewFileExecutor()
	executor := cmd.NewConsoleExecutor()
	executor.Execute()
}
