package main

import (
	"fmt"
	"time"

	"nipun.io/snake_ladder/cmd"
	"nipun.io/snake_ladder/reader"
)

func main() {
	reader := reader.NewConsoleReader()
	done := make(chan struct{})
	executor := cmd.NewGame(reader, done)
	executor.Init()
	go func() {
		executor.Execute()
	}()
	fmt.Println("started the execution")
	time.Sleep(time.Second * 10)
	go func() {
		i := 0
		for {
			if i == 10 {
				return
			}
			fmt.Println(executor.Status())
			time.Sleep(time.Second)
			i++
		}
	}()
	done <- struct{}{}
	close(done)
}
