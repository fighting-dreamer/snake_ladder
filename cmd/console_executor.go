package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ConsoleExecutor struct {
}

func NewConsoleExecutor() *ConsoleExecutor {
	return &ConsoleExecutor{}
}

func (ce *ConsoleExecutor) Execute() error {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		words := strings.Split(text, " ")
		fmt.Println(words)
	}
}
